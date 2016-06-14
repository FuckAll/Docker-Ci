/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/26 14:13
 */

package infrastructure

import (
	"github.com/FuckAll/Docker-Ci/api"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
)

var Etcd CreateEtcdOpts

var EtcdContainerId string

type CreateEtcdOpts struct {
	Image string
	Name  string
	Init  string
}

func init() {
	etcdmap := (conf.Config.Infrastructure["etcd"]).(map[string]interface{})
	Etcd.Name = (etcdmap["name"]).(string)
	Etcd.Image = (etcdmap["image"]).(string)
	Etcd.Init = (etcdmap["init"]).(string)
}

func StartEtcd() error {
	err := CreateEtcdContainer()
	if err != nil {
		return err
	}

	err = StartEtcdContainer()
	if err != nil {
		return err
	}
	return nil
}

func StopEtcd() error {
	err := StopEtcdContainer()
	if err != nil {
		return err
	}
	err = RemoveEtcdContainer()
	if err != nil {
		return nil
	}
	return nil
}

func CreateEtcdContainer() error {
	name := conf.Tracer + "-" + Etcd.Name
	etcdComtainerId, err := api.CreateContainer(name, Etcd.Image, []string{"app:/test"})
	if err != nil {
		log.Terror(conf.Tracer, "CreateEtcdContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "CreateEtcdContainer Complate!")
	EtcdContainerId = etcdComtainerId
	return nil
}

func StartEtcdContainer() error {
	err := api.StartContainer(EtcdContainerId, conf.Config.Bridge)
	if err != nil {
		log.Terror(conf.Tracer, "StartEtcdContainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StartEtcdContainer Complate!")
	return nil
}

func StopEtcdContainer() error {
	err := api.StopContainer(EtcdContainerId, 20)
	if err != nil {
		log.Terror(conf.Tracer, "StopEtcdContainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StopEtcdContainer Complate!")
	return nil

}

func RemoveEtcdContainer() error {
	err := api.RemoveContainer(EtcdContainerId, false)
	if err != nil {
		log.Terror("RemoveEtcdContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "RemoveEtcdContainer Complate!")
	return nil
}
