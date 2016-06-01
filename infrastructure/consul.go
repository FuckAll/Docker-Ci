/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/26 14:13
 */

package infrastructure

import (
	"bytes"
	"errors"
	"github.com/FuckAll/Docker-Ci/api"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
	"net/http"
	//	"strconv"
	"time"
)

var Consul CreateConsulOpts

var client = &http.Client{}

var ConsulContainerId string

type CreateConsulOpts struct {
	Image string
	Name  string
	Init  string
}

func init() {
	consulmap := (conf.Config.Infrastructure["consul"]).(map[string]interface{})
	Consul.Name = (consulmap["name"]).(string)
	Consul.Image = (consulmap["image"]).(string)
	Consul.Init = (consulmap["init"]).(string)
}

func StartConsul() error {
	err := CreateConsulContainer()
	if err != nil {
		return err
	}

	err = StartConsulContainer()
	if err != nil {
		return err
	}
	for i := 0; ; i++ {
		if i > 30 {
			log.Infof("After for a long time we can't connection to consul")
		}

		if ConsulCheck() {
			log.Tinfof(conf.Tracer, "connection to consul success")
			break
		} else {
			log.Tinfof(conf.Tracer, "Try connection to consul %d time(s)", i+1)
			time.Sleep(time.Second)
		}
	}
	err = ConsulRegister()
	if err != nil {
		return err
	}

	return nil
}

func CreateConsulContainer() error {
	name := conf.Tracer + "-" + Consul.Name
	consulContainerId, err := api.CreateContainerWithCmd(name, Consul.Image, []string{"app:/test"}, []string{"consul", "agent", "-dev", "-bind=0.0.0.0", "-client=0.0.0.0"})
	if err != nil {
		log.Terror(conf.Tracer, "CreateConsulContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "CreateConsulContainer Complate!")
	ConsulContainerId = consulContainerId
	return nil
}

func StartConsulContainer() error {
	err := api.StartContainer(ConsulContainerId, conf.Config.Bridge)
	if err != nil {
		log.Terror(conf.Tracer, "StartConsulContainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "CreateConsulContainer Complate!")
	return nil
}

func StopConsulContainer() error {
	err := api.StopContainer(ConsulContainerId, 20)
	if err != nil {
		log.Terror(conf.Tracer, "StopConsulContainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StopConsulContainer Complate!")
	return nil

}

func RemoveConsulContainer() error {
	err := api.RemoveContainer(ConsulContainerId, false)
	if err != nil {
		log.Terror("RemoveConsulContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "RemoveConsulContainer Complate!")
	return nil
}

func ConsulCheck() bool {
	url := "http://" + conf.Tracer + "-" + Consul.Name + "." + conf.Config.Bridge + ":8500/v1/agent/services"
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	if resp.Status != "200 OK" {
		return false
	} else {
		return true
	}
}

func ConsulRegister() error {
	url := "http://" + conf.Tracer + "-" + Consul.Name + ":8500/v1/agent/service/register"
	for _, service := range conf.Config.Services {
		port := service.Env["P"].(string)
		//p, _ := strconv.Atoi(port)
		jsonStr := []byte(`{"Name":"` + service.Name + `", "Port":` + port + `, "Address":"` + conf.Tracer + "-" + service.Name + `"}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.Status != "200 OK" {
			log.Tinfo(conf.Tracer, "REG service error ")
			return errors.New("REG service error")
		}
	}

	return nil
}
