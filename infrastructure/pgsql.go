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
	"net/http"
	"strings"
	"time"

	"github.com/FuckAll/Docker-Ci/api"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/FuckAll/Docker-Ci/filewalk"
	_ "github.com/lib/pq"
	"github.com/wothing/log"
)

var Postgres CreatePostgresOpts

var PostgresContainerId string

type CreatePostgresOpts struct {
	Image  string
	Name   string
	Passwd string
	Dname  string
	Duser  string
	Init   string
}

func init() {
	postgresmap := (conf.Config.Infrastructure["pgsql"]).(map[string]interface{})
	Postgres.Name = (postgresmap["name"]).(string)
	Postgres.Image = (postgresmap["image"]).(string)
	Postgres.Passwd = (postgresmap["passwd"]).(string)
	Postgres.Dname = (postgresmap["dname"]).(string)
	Postgres.Duser = (postgresmap["duser"]).(string)
	Postgres.Init = (postgresmap["init"]).(string)
}

func StartPostgres() error {
	err := CreatePostgresContainer()
	if err != nil {
		return err
	}
	err := StartPostgresContainer()
	if err != nil {
		return err
	}
	dirs := strings.Split(Postgres.Init, ",")
	var projectPath string
	if strings.HasSuffix(conf.Config.ProjectPath, "/") {
		projectPath = conf.Config.ProjectPath
	} else {
		projectPath = conf.Config.ProjectPath + "/"
	}
	for _, dir := range dirs {
		sqlDir := projectPath + dir
		err := PostgresInit(filewalk.WalkDir(sqlDir, "sql").FileList()...)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreatePostgresContainer() error {
	name := conf.Tracer + "-" + Postgres.Name
	postgresContainerId, err := api.CreateContainer(name, Postgres.Image, []string{"app:/test"}, "POSTGRES_DB=meidb", "POSTGRES_PASSWORD=wothing")
	if err != nil {
		log.Terror(conf.Tracer, "CreatePostgresContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "CreatePostgresContainer Complate!")
	PostgresContainerId = postgresContainerId
	return nil
}

func StartPostgresContainer() error {
	err := api.StartContainer(PostgresContainerId, conf.Config.Bridge)
	if err != nil {
		log.Terror(conf.Tracer, "StartPostgresContainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StartPostgresContainer Complate!")
	return nil
}

func StopPostgresContainer() error {
	err := api.StopContainer(PostgresContainerId, 20)
	if err != nil {
		log.Terror(conf.Tracer, "StopPostgresContainer Error:", err)
		return err

	}
	log.Tinfo(conf.Tracer, "StopPostgresContainer Complate!")
	return nil

}

func RemovePostgresContainer() error {
	err := api.RemoveContainer(PostgresContainerId, false)
	if err != nil {
		log.Terror("RemovePostgresContainer Error:", err)
		return err
	}
	log.Tinfo(conf.Tracer, "RemovePostgresContainer Complate!")
	return nil
}

func PostgresCheck() bool {
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

func PostgresInit(files ...string) error {
	dsn := FMT("postgres://" + Postgres.Duser + ":" + Postgres.Passwd + "@" + conf.Tracer + "-" + Postgres.Name + "." + conf.Config.Bridge + ":5432/" + Postgres.Dname + "?sslmode=disable")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Tinfof(conf.Tracer, "error on connecting to db: %s", err)
		return err
	}
	defer db.Close()
	for i := 0; ; i++ {
		if i > 30 {
			log.Info("After for a long time we can't connection to database")
			return errors.New("After for a long time we can't connection to database")
		}
		if db.Ping() != nil {
			log.Tinfof(conf.Tracer, "Try connection to database %d time(s)", i+1)
			time.Sleep(time.Second)
		} else {
			log.Tinfof(conf.Tracer, "connection to postgresql success")
			break
		}
	}

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`)
	if err != nil {
		log.Tinfo("sql error:", err)
		return err
	}

	fmt.Println("sql list:", files)

	for _, f := range files {
		sql, err := ioutil.ReadFile(f)
		if err != nil {
			log.Tinfof(conf.Tracer, "reading sql file error : %s", f)
			return err
		}

	}
}
