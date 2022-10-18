package kubectl

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// setKubeClient set up kubernetes clientset from the given kubeconfig and return with that
func setKubeClient(kubeconfig *string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}

// CreateNamespace create namespace to provided kubeconfig kubecontext
func CreateNamespace(namespace string, kubeconfig *string) {
	var clientset = setKubeClient(kubeconfig)

	//TODO: Check if setKubeClient failed

	nsName := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}

	clientset.CoreV1().Namespaces().Create(context.Background(), nsName, metav1.CreateOptions{})
}

// IsNamespaceExists check the given namespace is exists already or not
func IsNamespaceExists(namespace string, kubeconfig *string) bool {
	var clientset = setKubeClient(kubeconfig)

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, namespaceItem := range namespaces.Items {
		if namespaceItem.Name == namespace {
			return true
		}
	}

	return false
}
