package kubeconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Provider interface {
	GetKubeconfig() (string, error)
}

type EKSProvider struct{ ClusterName string }

func (e *EKSProvider) GetKubeconfig() (kubeconfig string, err error) {
	var tempDir string

	tempDir, err = ioutil.TempDir("", "")
	defer func() {
		if err != nil {
			os.RemoveAll(tempDir)
		}
	}()

	kubeconfig = filepath.Join(tempDir, "kubeconfig.yaml")

	if out, err := exec.Command(
		"aws", "eks", "update-kubeconfig",
		fmt.Sprintf("--name=%s", e.ClusterName),
		fmt.Sprintf("--kubeconfig=%s", kubeconfig),
	).CombinedOutput(); err != nil {
		err = fmt.Errorf("%s: %s", err.Error(), string(out))
		return "", err
	}

	return
}
