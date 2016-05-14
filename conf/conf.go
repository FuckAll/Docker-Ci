/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 17:11
 */

package conf

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/pborman/uuid"
	"github.com/wothing/log"
)

var (
	Tracer string

	Concurrent int

	REPO        string
	ProjectPath string // This is a absolute PATH
	SQLDir      string

	DockerRegistryPosition string
	REV                    string

	PGImage string

	RedisImage string

	ConsulImage string

	Services    []Service
	ServicesRun []Service //TODO use?
)

var (
	Push          bool   //= flag.Bool("push", false, "show build version")
	BuildList     string //= flag.String("b", "all", "building list such as : appway,interway,user split by ,")
	DependanceTag string //= flag.String("depend", "latest", "Detected modes not build docker image tag")
	TestOnly      bool   //= flag.Bool("t", false, "after build and run collect garbage container")
)

type Service struct {
	Name string
	Path string
	Para string
}

func init() {
	log.SetFlags(log.LstdFlags | log.Llevel)

	flag.Parse()
	Tracer = uuid.New()[:8]
	logfile := "/log/" + Tracer + ".log"
	_, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 066)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile("woci.json")
	if err != nil {
		log.Tfatalf(Tracer, "read woci.json error: %v", err)
	}

	cm := make(map[string]interface{})
	err = json.Unmarshal(data, &cm)
	if err != nil {
		log.Tfatalf(Tracer, "woci.json unmarshal error: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Tfatalf(Tracer, "woci.json file illegal--> %v", r)
		}
	}()

	Concurrent = int(cm["Concurrent"].(float64))
	REPO = cm["REPO"].(string)
	ProjectPath = cm["ProjectPath"].(string) // This is a absolute PATH
	SQLDir = cm["SQLDir"].(string)

	DockerRegistryPosition = cm["DockerRegistryPosition"].(string)

	//REV = cm["REV"].(string)

	PGImage = cm["PGImage"].(string)

	RedisImage = cm["RedisImage"].(string)

	ConsulImage = cm["ConsulImage"].(string)

	services := cm["Services"].([]interface{})
	for _, v := range services {
		s := Service{
			Name: v.(map[string]interface{})["Name"].(string),
			Path: v.(map[string]interface{})["Path"].(string),
			Para: v.(map[string]interface{})["Para"].(string),
		}
		Services = append(Services, s)
	}

	log.Tinfo(Tracer, "load woci.json succeed")
}
