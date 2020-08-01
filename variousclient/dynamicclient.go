package main

import (
	"fmt"
	"github.com/youshintop/kubernetes-samples/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

//client-go dynamicclient demo
func main() {
	dynamicClient, err := dynamic.NewForConfig(common.LoadConfig())
	if err != nil {
		panic(err)
	}

	gvr := schema.GroupVersionResource{
		Version: "v1",
		Resource: "pods",
	}

	unstructObj, err := dynamicClient.Resource(gvr).Namespace(corev1.NamespaceDefault).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), podList)
	if err != nil {
		panic(err)
	}

	for _, pod := range podList.Items {
		fmt.Printf("Namespace:%v \t Name:%v \t Status:%v \n", pod.Namespace, pod.Name, pod.Status.Phase)
	}
}
