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

var Redis CreateRedisOpts

var RedisContainerId string

type CreateRedisOpts struct {
	Image string
	Name  string
	Init  string
}

func init() {
	redismap := (conf.Config.Infrastructure["redis"]).(map[string]interface{})
	Redis.Name = (redismap["name"]).(string)
	Redis.Image = (redismap["image"]).(string)
	Redis.Init = (redismap["init"]).(string)
}

func StartRedis() error {
	err := CreateRedisContainer()
	if err != nil {
		return err
	}
	err = StartRedisContainer()
	if err != nil {
		return err
	}
	return nil

}

func CreateRedisContainer() error {
	name := conf.Tracer + "-" + Redis.Name
	redisContainerId, err := api.CreateContainer(name, Redis.Image, []string{"app:/test"})
	if err != nil {
		log.Terror("CreateRedisContainer Error:", err)
		return err
	}
	log.Info("CreateRedisContainer Complate!")
	RedisContainerId = redisContainerId
	return nil
}

func StartRedisContainer() error {
	err := api.StartContainer(RedisContainerId, conf.Config.Bridge)
	if err != nil {
		log.Terror("StartRedis Error:", err)
		return err

	}
	log.Info("StartRedisContainer Complate!")
	return nil
}

func StopRedisContainer() error {
	err := api.StopContainer(RedisContainerId, 20)
	if err != nil {
		log.Terror("StopRedisContainer Error:", err)
		return err

	}
	log.Info("StopRedisContainer Complate!")
	return nil

}

func RemoveRedisContainer() error {
	err := api.RemoveContainer(RedisContainerId, false)
	if err != nil {
		log.Terror("RemoveRedis Error:", err)
		return err
	}
	log.Info("RemoveRedisContainer Complate!")
	return nil
}
