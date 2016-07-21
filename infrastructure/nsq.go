/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/26 14:13
 */

package infrastructure

import (
	"time"

	"github.com/FuckAll/Docker-Ci/api"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
)

// Nsq used
var Nsqd CreateNsqdOpts

// NsqdContainerID used
var NsqdContainerID string

//CreateNsqdOpts used
type CreateNsqdOpts struct {
	Image string
	Name  string
	Init  string
}

func init() {
	nsqdmap := (conf.Config.Infrastructure["nsqd"]).(map[string]interface{})
	Nsqd.Name = (nsqdmap["name"]).(string)
	Nsqd.Image = (nsqdmap["image"]).(string)
	Nsqd.Init = (nsqdmap["init"]).(string)
}

// StartEtcd used to start etcd
func StartNsqd() error {
	err := createNsqdContainer()
	if err != nil {
		return err
	}

	err = startNsqdContainer()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	return nil
}

// StopEtcd used to stop etcd
func StopNsqd() error {
	err := stopNsqdContainer()
	if err != nil {
		return err
	}
	err = removeNsqdContainer()
	if err != nil {
		return nil
	}
	return nil
}

// CreateNsqdContainer userd to create etcd Container
func createNsqdContainer() error {

	name := conf.Tracer + "-" + Nsqd.Name
	id, err := api.CreateContainer(name, Nsqd.Image, []string{"app:/test"})
	if err != nil {
		log.Terror(conf.Tracer, "CreateNsqdContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "CreateNsqdContainer Complate!")
	NsqdContainerID = id
	return nil
}

// StartNsqdContainer used to Start nsqd container
func startNsqdContainer() error {
	err := api.StartContainer(NsqdContainerID, conf.Config.Bridge)
	if err != nil {
		log.Terror(conf.Tracer, "StartNsqdtainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StartNsqdContainer Complate!")
	return nil
}

// StopNsqdContainer used to Stop Nsqd container
func stopNsqdContainer() error {
	err := api.StopContainer(NsqdContainerID, 20)
	if err != nil {
		log.Terror(conf.Tracer, "StopNsqdContainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StopNsqdContainer Complate!")
	return nil

}

// RemoveNsqdContainer used to Remobve Nsqd container
func removeNsqdContainer() error {
	err := api.RemoveContainer(NsqdContainerID, false)
	if err != nil {
		log.Terror("RemoveNsqdContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "RemoveNsqdContainer Complate!")
	return nil
}
