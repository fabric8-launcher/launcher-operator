package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LauncherSpec defines the desired state of Launcher
type LauncherSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	GitHub         GitHubConfig `json:"github"`
	CreatorEnabled bool         `json:"creator"`
}

// GitHubConfig defines the GitHub configuration
type GitHubConfig struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// LauncherStatus defines the observed state of Launcher
type LauncherStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Launcher is the Schema for the launchers API
// +k8s:openapi-gen=true
type Launcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LauncherSpec   `json:"spec,omitempty"`
	Status LauncherStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LauncherList contains a list of Launcher
type LauncherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Launcher `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Launcher{}, &LauncherList{})
}
