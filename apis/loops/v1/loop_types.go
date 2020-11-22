package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	// Loop is the Schema for the Loops API
	// +kubebuilder:object:root=true
	// +kubebuilder:storageversion
	// +kubebuilder:printcolumn:name="Announced",type=string,JSONPath=`.status.lastExecution`,description="Last Execution",priority=1
	// +kubebuilder:resource:shortName=if
	// +kubebuilder:subresource:status
	// +genclient
	Loop struct {
		meta.TypeMeta   `json:",inline"`
		meta.ObjectMeta `json:"metadata,omitempty"`

		Spec   LoopSpec   `json:"spec,omitempty"`
		Status LoopStatus `json:"status,omitempty"`
	}

	// LoopList contains a list of Loop
	// +kubebuilder:object:root=true
	LoopList struct {
		meta.TypeMeta `json:",inline"`
		meta.ListMeta `json:"metadata,omitempty"`
		Items         []Loop `json:"items"`
	}

	LoopSpec struct {
		Endpoint string        `json:"endpoint,omitempty"`
		Script   string        `json:"script,omitempty"`
		Every    meta.Duration `json:"every,omitempty"`
		// +optional
		Imports Imports `json:"imports,omitempty"`
	}

	// Define Loops status
	LoopStatus struct {
		Deployed             bool       `json:"deployed,omitempty"`
		LastExecution        *meta.Time `json:"lastExecution,omitempty"`
		LastExecutionSuccess *meta.Time `json:"lastExecutionSuccess,omitempty"`
	}

	Imports    []LoopImport
	LoopImport struct {
		Secret ImportSecret `json:"secret,omitempty"`
	}

	ImportSecret struct {
		Name string `json:"name"`
	}
)
