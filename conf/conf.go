/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 17:11
 */

package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pborman/uuid"
	"github.com/wothing/log"
)

var (
	Tracer string
)

var dockerCi DockerCi

type DockerCi struct {
	ProjectPath    string                 // This is a absolute PATH
	DockerApi      string                 //Docker Api Default "tcp://127.0.0.1:2375"
	Bridge         string                 //Docker Bridge Default bridge
	ServicesImage  string                 // Service Base image for example: alpine:latest
	Infrastructure map[string]interface{} // Base Infrastructure for example: pgsql redis consul
	Services       []Service
}

type Service struct {
	Name           string
	DockerFilePath string
	BuildCommand   string
	Env            map[string]interface{}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Llevel)

	Tracer = uuid.New()[:8]
	data, err := ioutil.ReadFile("dockerci.json")
	if err != nil {
		log.Tfatalf(Tracer, "read dockerapi.json error: %v", err)
	}

	cm := make(map[string]interface{})
	err = json.Unmarshal(data, &cm)
	if err != nil {
		log.Tfatalf(Tracer, "dockerci.json unmarshal error: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Tfatalf(Tracer, "dockerci.json file illegal--> %v", r)
		}
	}()

	dockerCi.ProjectPath = cm["ProjectPath"].(string) // This is a absolute PATH
	dockerCi.DockerApi = cm["DockerApi"].(string)     // Docker api
	dockerCi.Bridge = cm["Bridge"].(string)
	dockerCi.ServicesImage = cm["ServicesImage"].(string)
	dockerCi.Infrastructure = cm["Infrastructure"].(map[string]interface{})

	services := cm["Services"].([]interface{})
	for _, v := range services {
		v1 := v.(map[string]interface{})
		s := Service{
			Name:           v1["Name"].(string),
			DockerFilePath: v1["DockerFilePath"].(string),
			BuildCommand:   v1["BuildCommand"].(string),
			Env:            v1["Env"].(map[string]interface{}),
		}
		dockerCi.Services = append(dockerCi.Services, s)

	}
	log.Tinfo(Tracer, "load dockerci.json succeed")
}
