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

func ReadYAMLResourceFile(path string) banzaicloud.IstioControlPlane {
	var data = fileRead(path)
	var controlPlane banzaicloud.IstioControlPlane
	err := yaml.Unmarshal(data, controlPlane)
	if err != nil {
		panic("Aww, this resource file cannot convert to IstioControlPlane resource")
	}

	return controlPlane
}

func ReadYAMLChartsFile(path string) {
	// TODO: Read chars from file in install command
}