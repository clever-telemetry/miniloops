package local

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	loops "github.com/clever-telemetry/miniloops/apis/loops/v1"
	"github.com/clever-telemetry/miniloops/pkg/client"
	"github.com/clever-telemetry/miniloops/pkg/warp10"
	"github.com/sirupsen/logrus"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
)

type Script struct {
	sync.RWMutex
	object     *loops.Loop
	endpoint   string
	warpscript string
	every      time.Duration
	version    int64
	stop       chan struct{}
	recorder   record.EventRecorder
	ticker     *time.Ticker
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
}

func (script *Script) SetEndpoint(endpoint string) {
	script.Lock()
	defer script.Unlock()

	script.endpoint = endpoint
}

func (script *Script) SetWarpScript(ws string) {
	script.Lock()
	defer script.Unlock()

	script.warpscript = ws
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

func (script *Script) SetVersion(v int64) {
	script.Lock()
	defer script.Unlock()

	script.version = v
}

func (script *Script) Start() {
	script.Lock()
	defer script.Unlock()

	if script.ticker != nil {
		return
	}

	script.ticker = time.NewTicker(script.every)
	go func() {
		for range script.ticker.C {

			res, err := script.Exec()
			if err != nil {
				script.recorder.AnnotatedEventf(script.object, map[string]string{
					"version": fmt.Sprintf("%d", script.version),
				}, "Warning", "ExecError", err.Error())
				continue
			}

			script.recorder.AnnotatedEventf(
				script.object,
				map[string]string{
					"version":            fmt.Sprintf("%d", script.version),
					"fetched":            fmt.Sprintf("%d", res.Fetched()),
					"ops":                fmt.Sprintf("%d", res.Ops()),
					"serverSideDuration": res.Elapsed().String(),
				},
				"Normal",
				"ExecSuccess",
				"Success (fetched=%d ops=%d elapsed=%s)",
				res.Fetched(), res.Ops(), res.Elapsed(), res.StackRawString(),
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
				ws.WriteString(fmt.Sprintf("'%v' '%s' STORE\n", string(v), string(k)))
			}
		}
	}
	ws.WriteString("LINEON\n")
	ws.WriteString(script.warpscript)

	logrus.Info("WarpScript\n", ws.String())

	res := warp10.NewRequest(script.endpoint, script.warpscript).Exec()
	if res.IsErrored() {
		return nil, res.Error()
	}

	return res, nil
}
