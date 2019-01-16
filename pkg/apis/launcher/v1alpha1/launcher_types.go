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
	CreatorEnabled bool         `json:"creator"`
	Environment string `json:"environment"`
	Catalog CatalogConfig `json:"catalog"`
	OpenShift OpenShiftConfig `json:"openshift"`
	Git         GitConfig `json:"git"`
	Jms         JmsConfig `json:"jms"`
	Keycloak    KeycloakConfig `json:"keycloak"`
	Segment SegmentConfig `json:"segment"`
	Sentry SentryConfig `json:"sentry"`
}

type CatalogConfig struct {
	Url string `json:"url"`
	Ref string `json:"ref"`
	Filter string `json:"filter"`
	ReindexToken    string `json:"reindexToken"`
}

type OpenShiftConfig struct {
	ApiUrl string `json:"apiUrl"`
	ConsoleUrl string `json:"consoleUrl"`
	SubscriptionToken string `json:"subscriptionToken"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Impersonate bool `json:"impersonate"`
	Clusters []OpenShiftClusterConfig `json:"clusters"`
}

type OpenShiftClusterConfig struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ApiUrl string `json:"apiUrl"`
	ConsoleUrl string `json:"consoleUrl"`
	Type string `json:"type"`
}

// GitConfig defines the Git configuration
type GitConfig struct {
	Provider    string `json:"provider"`
	Url 	string `json:"url"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type JmsConfig struct {
	Url 	string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type KeycloakConfig struct {
	Url 	string `json:"url"`
	Realm 	string `json:"realm"`
	ClientId	string `json:"clientId"`
}

type SegmentConfig struct {
	Token string `json:"token"`
}

type SentryConfig struct {
	CreatorToken string `json:"creatorToken"`
	BackendToken string `json:"backendToken"`
	FrontendToken string `json:"frontendToken"`
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
