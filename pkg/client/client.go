package client

import (
	"github.com/clever-telemetry/miniloops/client/loops"
	loopv1 "github.com/clever-telemetry/miniloops/client/loops/typed/loops/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
)

var (
	config           *rest.Config
	EventRecorderFor func(name string) record.EventRecorder
)

func Set(_config *rest.Config) {
	config = _config
}

func Native() *kubernetes.Clientset {
	return kubernetes.NewForConfigOrDie(config)
}

func Loops() loops.Interface {
	return loops.NewForConfigOrDie(config)
}

func LoopsFor(namespace string) loopv1.LoopInterface {
	return Loops().LoopsV1().Loops(namespace)
}
