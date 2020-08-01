package main

import (
	"errors"
	"fmt"
	"github.com/youshintop/kubernetes-samples/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// client-go resetclient demo.
func main() {
	config := common.LoadConfig()
	if config == nil {
		panic(errors.New("Failed load kubeconfig."))
	}

	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		fmt.Println("Failed create restClient.")
	}

	var podList = &corev1.PodList{}
	restClient.Get().
			Namespace(corev1.NamespaceDefault).
			Resource("pods").
			VersionedParams(
				&metav1.ListOptions{},
				scheme.ParameterCodec).
		Do().Into(podList)

	if err!= nil {
		fmt.Println("Failed get pod list.")
	}

	for _, pod := range podList.Items {
		fmt.Printf("Namespace:%v \t Name:%v \t Status:%v \n", pod.Namespace, pod.Name, pod.Status.Phase)
	}
}
