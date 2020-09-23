package state

import "github.com/roboll/helmfile/pkg/kubeconfig"

type EnvironmentSpec struct {
	Values  []interface{} `yaml:"values,omitempty"`
	Secrets []string      `yaml:"secrets,omitempty"`

	// MissingFileHandler instructs helmfile to fail when unable to find a environment values file listed
	// under `environments.NAME.values`.
	//
	// Possible values are  "Error", "Warn", "Info", "Debug". The default is "Error".
	//
	// Use "Warn", "Info", or "Debug" if you want helmfile to not fail when a values file is missing, while just leaving
	// a message about the missing file at the log-level.
	MissingFileHandler *string `yaml:"missingFileHandler,omitempty"`

	// Kubeconfig instructs helmfile to use an authentication provider when connecting to Kubernetes for this environment.
	Kubeconfig *KubeconfigSpec `yaml:"kubeconfig,omitempty"`
}

type KubeconfigSpec struct {
	AWS *EKSConfigSpec `yaml:"aws,omitempty"`
}

type EKSConfigSpec struct {
	ClusterName string `yaml:"clusterName,omitempty"`
}

func (e *EnvironmentSpec) getKubeconfigProvider() kubeconfig.Provider {
	if e.Kubeconfig == nil {
		return nil
	}
	if e.Kubeconfig.AWS != nil {
		return &kubeconfig.EKSProvider{ClusterName: e.Kubeconfig.AWS.ClusterName}
	}
	return nil
}
