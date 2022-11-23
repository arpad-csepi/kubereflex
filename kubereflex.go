package kubereflex

import (
	"time"

	helm "github.com/arpad-csepi/kubereflex/helm"
	"github.com/arpad-csepi/kubereflex/io"
	kubectl "github.com/arpad-csepi/kubereflex/kubectl"
)

// TODO: Optional parameters like args, namespace (maybe in KLI pass nil parameter to here)
// TODO: Default value to namespace (maybe define defaults in KLI call this function)
func InstallHelmChart(chartUrl string, repositoryName string, chartName string, releaseName string, namespace string, args map[string]string, kubeconfig *string) {
	if !kubectl.IsNamespaceExists(namespace, kubeconfig) {
		kubectl.CreateNamespace(namespace, kubeconfig)
	}

	if !helm.IsRepositoryExists(repositoryName) {
		helm.RepositoryAdd(repositoryName, chartUrl)
	}

	helm.Install(repositoryName, chartName, releaseName, namespace, args)
}

func UninstallHelmChart(releaseName string, namespace string) {
	// TODO: Some check before run Uninstall

	helm.Uninstall(releaseName, namespace)
}

func Verify(deploymentName string, namespace string, kubeconfig *string, timeout time.Duration) {
	kubectl.Verify(deploymentName, namespace, kubeconfig, timeout)
}

func ReadYAMLResourceFile(path string) {
	io.ReadYAMLResourceFile(path)
}