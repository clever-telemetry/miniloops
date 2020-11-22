package local

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	loops "github.com/clever-telemetry/miniloops/apis/loops/v1"
	"github.com/clever-telemetry/miniloops/pkg/client"
	"github.com/clever-telemetry/miniloops/pkg/runner/metrics"
	"github.com/clever-telemetry/miniloops/pkg/warp10"
	"github.com/iancoleman/strcase"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
)

type Script struct {
	sync.RWMutex
	object       *loops.Loop
	every        time.Duration
	stop         chan struct{}
	recorder     record.EventRecorder
	ticker       *time.Ticker
	execCount    prometheus.Counter
	errorCount   prometheus.Counter
	execDuration prometheus.Counter
	fetched      prometheus.Counter
	operations   prometheus.Counter
}

func NewScript() *Script {
	return &Script{
		RWMutex:  sync.RWMutex{},
		stop:     make(chan struct{}),
		recorder: client.EventRecorderFor("runner"),
	}
}

func (script *Script) SetObject(o *loops.Loop) {
	script.Lock()
	defer script.Unlock()

	script.object = o
	script.execCount = metrics.ExecCount.WithLabelValues(o.GetNamespace(), o.GetName())
	script.errorCount = metrics.ExecErrorCount.WithLabelValues(o.GetNamespace(), o.GetName())
	script.execDuration = metrics.ExecDuration.WithLabelValues(o.GetNamespace(), o.GetName())
	script.fetched = metrics.FetchedCount.WithLabelValues(o.GetNamespace(), o.GetName())
	script.operations = metrics.OperationsCount.WithLabelValues(o.GetNamespace(), o.GetName())
}

func (script *Script) SetEvery(d time.Duration) {
	script.Lock()
	defer script.Unlock()

	if script.every == d {
		return
	}

	script.every = d

	if script.ticker != nil {
		script.ticker.Reset(d)
	}
}

func (script *Script) Start() {
	script.Lock()
	defer script.Unlock()

	if script.ticker != nil || script.object == nil {
		return
	}

	script.ticker = time.NewTicker(script.every)
	go func() {
		for range script.ticker.C {
			start := time.Now()

			res, err := script.Exec()
			script.execCount.Inc()
			script.execDuration.Add(float64(time.Since(start).Milliseconds()))

			serr := client.
				LoopsFor(script.object.GetNamespace()).
				SetLastExecution(context.Background(), script.object.GetName(), time.Now(), err == nil)
			if serr != nil {
				logrus.WithError(serr).Error("cannot set loop last execution")
			}

			if err != nil {
				script.errorCount.Inc()
				script.recorder.AnnotatedEventf(script.object, map[string]string{
					"version": fmt.Sprintf("%d", script.object.GetGeneration()),
				}, "Warning", "ExecError", err.Error())
				continue
			}

			script.recorder.AnnotatedEventf(
				script.object,
				map[string]string{
					"version":            fmt.Sprintf("%d", script.object.GetGeneration()),
					"fetched":            fmt.Sprintf("%d", res.Fetched()),
					"ops":                fmt.Sprintf("%d", res.Ops()),
					"serverSideDuration": res.Elapsed().String(),
					"stack":              res.StackRawString(),
				},
				"Normal",
				"ExecSuccess",
				"Success (fetched=%d ops=%d elapsed=%s): %+v",
				res.Fetched(),
				res.Ops(),
				res.Elapsed(),
				serializeInterfaceSlice(res.StackSlice()),
			)

		}
	}()
}

func (script *Script) Stop() {
	script.Lock()
	defer script.Unlock()

	if script.ticker == nil {
		return
	}

	script.ticker.Stop()
}

func (script *Script) Exec() (*warp10.Response, error) {
	script.Lock()
	defer script.Unlock()

	ws := bytes.NewBuffer([]byte{})

	ws.WriteString(fmt.Sprintf("'%s' 'LOOP_NAME' STORE\n", script.object.GetName()))
	ws.WriteString(fmt.Sprintf("'%s' 'LOOP_NAMESPACE' STORE\n", script.object.GetNamespace()))

	//for k, v := range script.object.GetAnnotations() {
	//	ws.WriteString(fmt.Sprintf("`%s` `annotation.%s` STORE\n", v, k))
	//}

	for k, v := range script.object.GetLabels() {
		ws.WriteString(fmt.Sprintf("`%s` `label.%s` STORE\n", v, k))
	}

	for _, loopImport := range script.object.Spec.Imports {
		if loopImport.Secret.Name != "" {
			secret, err := client.
				Native().
				CoreV1().
				Secrets(script.object.GetNamespace()).
				Get(context.Background(), loopImport.Secret.Name, meta.GetOptions{})
			if err != nil {
				return nil, err
			}

			for k, v := range secret.Data {
				ws.WriteString(fmt.Sprintf(
					"'%v' '%s' STORE\n",
					string(v),
					fmt.Sprintf("%s.%s", strcase.ToLowerCamel(secret.GetName()), strcase.ToLowerCamel(k)),
				))
			}
		}
	}
	ws.WriteString("LINEON\n")
	ws.WriteString(script.object.Spec.Script)

	logrus.Debug("WarpScript\n", ws.String())

	res := warp10.NewRequest(script.object.Spec.Endpoint, ws.String()).Exec()
	if res.IsErrored() {
		return nil, res.Error()
	}

	return res, nil
}

func serializeInterfaceSlice(is []interface{}) string {
	s := make([]string, len(is))
	for i, ic := range is {
		s[i] = fmt.Sprintf("%v", ic)
	}
	return strings.Join(s, " ")
}
