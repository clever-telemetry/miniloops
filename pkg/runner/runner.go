package runner

import (
	v1 "github.com/clever-telemetry/miniloops/apis/loops/v1"
	"github.com/clever-telemetry/miniloops/pkg/runner/local"
	"k8s.io/apimachinery/pkg/types"
)

type Runner interface {
	UpsertLoop(*v1.Loop) error
	DeleteLoop(types.NamespacedName) error
	Has(types.NamespacedName) (bool, error)
	NamespaceCount() int
	LoopsCount() int
}

func Local() Runner {
	return local.NewLocalRunner()
}
