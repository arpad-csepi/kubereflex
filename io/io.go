package io

import (
	"os"
	"gopkg.in/yaml.v2"
	banzaicloud "github.com/banzaicloud/istio-operator/api/v2/v1alpha1"
)

func fileRead(path string) []byte {
	data, err := os.ReadFile("/tmp/dat")
    if err != nil {
		panic("Nope, file not found")
	}

	return data
}

func ReadYAMLResourceFile(path string) {
	var data = fileRead(path)
	var ControlPlane banzaicloud.IstioControlPlane
	yaml.Unmarshal(data, ControlPlane)
}

func ReadYAMLChartsFile(path string) {
	// TODO: Read chars from file in install command
}