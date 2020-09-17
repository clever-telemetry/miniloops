package pkg

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apimachinery/pkg/runtime"
	client "k8s.io/client-go/kubernetes/scheme"

	loops "github.com/clever-telemetry/miniloops/apis/loops/v1"
)

var Scheme = runtime.NewScheme()

func init() {
	apiextensions.AddToScheme(Scheme)
	client.AddToScheme(Scheme)
	loops.AddToScheme(Scheme)
}
