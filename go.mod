module github.com/clever-telemetry/miniloops

go 1.15

replace github.com/go-logr/logr => github.com/go-logr/logr v0.2.1

require (
	github.com/go-logr/logr v0.2.1
	github.com/iancoleman/strcase v0.1.2
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	k8s.io/apiextensions-apiserver v0.19.0
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.0
	sigs.k8s.io/controller-runtime v0.6.3
	sigs.k8s.io/kustomize/kustomize/v3 v3.8.7 // indirect
)
