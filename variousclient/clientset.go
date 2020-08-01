package main

import (
	"fmt"
	"github.com/youshintop/kubernetes-samples/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// client-go clientset demo
func main()  {
	clientset, err := kubernetes.NewForConfig(common.LoadConfig())
	if err != nil {
		panic(err)
	}

	podList, err := clientset.CoreV1().Pods(corev1.NamespaceDefault).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, pod := range podList.Items {
		fmt.Printf("Namespace:%v \t Name:%v \t Status:%v \n", pod.Namespace, pod.Name, pod.Status.Phase)
	}
}
