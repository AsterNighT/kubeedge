//go:generate apiserver-runtime-gen
package main

import (
	devices "github.com/kubeedge/kubeedge/cloud/pkg/apis/devices/v1alpha2"

	"k8s.io/klog/v2"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"
)

func main() {
	err := builder.APIServer.
		WithResource(&devices.Device{}).
		WithResource(&devices.DeviceModel{}).
		Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
