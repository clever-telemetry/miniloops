package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (loop *Loop) NamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: loop.GetNamespace(),
		Name:      loop.GetName(),
	}
}

func (loop *Loop) Equal(other *Loop) bool {
	return loop.EqualMeta(other.ObjectMeta) &&
		loop.Spec.Equal(other.Spec) &&
		loop.Status.Equal(other.Status)
}

func (loop *Loop) EqualMeta(other meta.ObjectMeta) bool {
	if loop.GetNamespace() != other.GetNamespace() ||
		loop.GetName() != other.GetName() ||
		loop.GetGeneration() != other.GetGeneration() ||
		!mapStringStringEqual(loop.GetLabels(), other.GetLabels()) ||
		!mapStringStringEqual(loop.GetAnnotations(), other.GetAnnotations()) {
		return false
	}

	return true
}

func (spec *LoopSpec) Equal(other LoopSpec) bool {
	if spec.Endpoint != other.Endpoint ||
		spec.Every.Duration != other.Every.Duration ||
		spec.Script != other.Script {
		return false
	}

	if !spec.Imports.Equal(other.Imports) {
		return false
	}

	return true
}

func (status *LoopStatus) Equal(other LoopStatus) bool {
	if status == nil {
		return false
	}

	if (status.LastExecution == nil) != (other.LastExecution == nil) {
		return false
	}
	if (status.LastExecution != nil && other.LastExecution != nil) && status.LastExecution.Time != other.LastExecution.Time {
		return false
	}

	if status.Deployed != other.Deployed {
		return false
	}

	return true
}

func (imports Imports) Equal(other Imports) bool {
	if len(imports) != len(other) {
		return false
	}

	for i, loopImport := range imports {
		if !other[i].Equal(loopImport) {
			return false
		}
	}

	return true
}

func (loopImport LoopImport) Equal(other LoopImport) bool {
	if loopImport.Secret.Name != other.Secret.Name {
		return false
	}

	return true
}
