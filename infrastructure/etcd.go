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
	"github.com/coreos/go-etcd/etcd"
	"github.com/wothing/log"
)

// Etcd used
var Etcd CreateEtcdOpts

// EtcdContainerID used
var EtcdContainerID string

//CreateEtcdOpts used
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

// StartEtcd used to start etcd
func StartEtcd() error {
	err := createEtcdContainer()
	if err != nil {
		return err
	}

	err = startEtcdContainer()
	if err != nil {
		return err
	}
	err = etcdInit()
	if err != nil {
		return err
	}
	return nil
}

// StopEtcd used to stop etcd
func StopEtcd() error {
	err := stopEtcdContainer()
	if err != nil {
		return err
	}
	err = removeEtcdContainer()
	if err != nil {
		return nil
	}
	return nil
}

// CreateEtcdContainer userd to create etcd Container
func createEtcdContainer() error {

	name := conf.Tracer + "-" + Etcd.Name
	id, err := api.CreateContainer(name, Etcd.Image, []string{"app:/test"})
	if err != nil {
		log.Terror(conf.Tracer, "CreateEtcdContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "CreateEtcdContainer Complate!")
	EtcdContainerID = id
	return nil
}

// StartEtcdContainer used to Start etcd container
func startEtcdContainer() error {
	err := api.StartContainer(EtcdContainerID, conf.Config.Bridge)
	if err != nil {
		log.Terror(conf.Tracer, "StartEtcdContainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StartEtcdContainer Complate!")
	return nil
}

// StopEtcdContainer used to Stop etcd container
func stopEtcdContainer() error {
	err := api.StopContainer(EtcdContainerID, 20)
	if err != nil {
		log.Terror(conf.Tracer, "StopEtcdContainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StopEtcdContainer Complate!")
	return nil

}

// RemoveEtcdContainer used to Remobve etcd container
func removeEtcdContainer() error {
	err := api.RemoveContainer(EtcdContainerID, false)
	if err != nil {
		log.Terror("RemoveEtcdContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "RemoveEtcdContainer Complate!")
	return nil
}

func etcdInit() error {
	etcdHost := "http://" + conf.Tracer + "-etcd" + ":2379"
	machines := []string{etcdHost}
	client := etcd.NewClient(machines)
	pgsqlHost := conf.Tracer + "-pgsql"
	redisHost := conf.Tracer + "-redis"
	keyValue := map[string]string{"/17mei/pgsql/host": pgsqlHost, "/17mei/pgsql/port": "5432", "/17mei/pgsql/name": Postgres.Dname, "/17mei/pgsql/user": Postgres.Duser, "/17mei/pgsql/password": Postgres.Passwd, "17mei/redis/host": redisHost, "/17mei/redis/port": "6379", "/17mei/redis/password": Redis.Passwd, "/17mei/mediastore/mode": "test", "/17mei/payment/mode": "test", "/17mei/push/apns": "flase"}
	for k, v := range keyValue {
		if _, err := client.Set(k, v, 0); err != nil {
			return err
		}
	}
	return nil

}
