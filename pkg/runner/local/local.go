package local

import (
	"sync"

	v1 "github.com/clever-telemetry/miniloops/apis/loops/v1"
	"k8s.io/apimachinery/pkg/types"
)

type (
	LocalRunner struct {
		data     map[string]*Script
		dataLock sync.RWMutex
	}
)

func NewLocalRunner() *LocalRunner {
	return &LocalRunner{
		data:     map[string]*Script{},
		dataLock: sync.RWMutex{},
	}
}

func (r *LocalRunner) UpsertLoop(loop *v1.Loop) error {
	if loop == nil {
		return nil
	}

	r.dataLock.Lock()
	script, ok := r.data[loop.NamespacedName().String()]
	if !ok {
		script = NewScript()
		r.data[loop.NamespacedName().String()] = script
	}
	r.dataLock.Unlock()

	script.SetObject(loop)
	script.SetEvery(loop.Spec.Every.Duration)
	script.Start()

	return nil
}

func (r *LocalRunner) DeleteLoop(namespacedName types.NamespacedName) error {
	has, err := r.Has(namespacedName)
	if err != nil {
		return err
	}
	if !has {
		return nil
	}

	r.dataLock.Lock()
	script := r.data[namespacedName.String()]
	delete(r.data, namespacedName.String())
	r.dataLock.Unlock()

	script.Stop()
	return nil
}

func (r *LocalRunner) Has(namespacedName types.NamespacedName) (bool, error) {
	r.dataLock.RLock()
	defer r.dataLock.RUnlock()

	_, ok := r.data[namespacedName.String()]
	return ok, nil
}

func (r *LocalRunner) NamespaceCount() int {
	r.dataLock.RLock()
	defer r.dataLock.RUnlock()

	nns := map[string]struct{}{}

	for _, loop := range r.data {
		loop.RLock()

		nns[loop.object.GetNamespace()] = struct{}{}

		loop.RUnlock()
	}

	return len(nns)
}

func (r *LocalRunner) LoopsCount() int {
	r.dataLock.RLock()
	defer r.dataLock.RUnlock()

	return len(r.data)
}
