package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type RequestStruct struct {
	Request []*RequestParamStruct `yaml:"request"`
}

type RequestParamStruct struct {
	Address string `yaml:"address"`
	Param   string `yaml:"param"`
}

func readReqeustYaml() (*RequestStruct, error) {
	RequestConfig := new(RequestStruct)

	request, err := ioutil.ReadFile("etc/request_url.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(request, RequestConfig)
	if err != nil {
		return nil, err
	}

	return RequestConfig, nil
}
