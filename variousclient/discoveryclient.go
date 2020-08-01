package main

import (
	"fmt"
	"github.com/youshintop/kubernetes-samples/common"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

//client-go discovery client demo
func main() {
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(common.LoadConfig())
	if err != nil {
		panic(err)
	}

	_, APIResourecList, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}

	for _, list := range APIResourecList {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err)
		}

		for _, resource := range list.APIResources {
			fmt.Printf("name:%v, \t group:%v, \t, version:%v\n", resource.Name, gv.Group, gv.Version)
		}
	}
}
