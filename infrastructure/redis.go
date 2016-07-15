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
	Image  string
	Name   string
	Passwd string
	Init   string
}

func init() {
	redismap := (conf.Config.Infrastructure["redis"]).(map[string]interface{})
	Redis.Name = (redismap["name"]).(string)
	Redis.Image = (redismap["image"]).(string)
	Redis.Passwd = (redismap["passwd"]).(string)
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

func StopRedis() error {
	err := StopRedisContainer()
	if err != nil {
		return err
	}
	err = RemoveRedisContainer()
	if err != nil {
		return err
	}
	return nil
}

func CreateRedisContainer() error {
	name := conf.Tracer + "-" + Redis.Name
	passwd := "REDIS_PASS=" + Redis.Passwd
	redisContainerId, err := api.CreateContainer(name, Redis.Image, []string{"app:/test"}, passwd)
	if err != nil {
		log.Terror(conf.Tracer, "CreateRedisContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "CreateRedisContainer Complate!")
	RedisContainerId = redisContainerId
	return nil
}

func StartRedisContainer() error {
	err := api.StartContainer(RedisContainerId, conf.Config.Bridge)
	if err != nil {
		log.Terror(conf.Tracer, "StartRedis Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StartRedisContainer Complate!")
	return nil
}

func StopRedisContainer() error {
	err := api.StopContainer(RedisContainerId, 20)
	if err != nil {
		log.Terror(conf.Tracer, "StopRedisContainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StopRedisContainer Complate!")
	return nil

}

func RemoveRedisContainer() error {
	err := api.RemoveContainer(RedisContainerId, false)
	if err != nil {
		log.Terror(conf.Tracer, "RemoveRedis Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "RemoveRedisContainer Complate!")
	return nil
}
