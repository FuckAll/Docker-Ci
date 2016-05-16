/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 16:30
 */

package ci

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	_ "github.com/lib/pq"
	"github.com/wothing/log"

	. "github.com/FuckAll/Docker-Ci/cmdrun"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/FuckAll/Docker-Ci/filewalk"
)

// DBStage main entry point of this module
func Pgsql() {
	CMD(FMT("docker run -it -d --net=test -v log:/log --name %s-pgsql -e POSTGRES_DB=meidb -e POSTGRES_PASSWORD=wothing %s", conf.Tracer, conf.PGImage))
	pgInit(filewalk.WalkDir(conf.ProjectPath+"/"+conf.SQLDir, "sql").FileList()...)
}

func pgInit(files ...string) {
	dsn := FMT("postgres://postgres:wothing@%s-pgsql.test:5432/meidb?sslmode=disable", conf.Tracer)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Tfatalf(conf.Tracer, "error on connecting to db: %s", err)
	}
	defer db.Close()

	for i := 0; ; i++ {
		if i > 30 {
			log.Tfatal("After for a long time we can't connection to database")
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
		log.Fatal("sql error:", err)
	}

	fmt.Println("sql list:", files)

	for _, f := range files {
		sql, err := ioutil.ReadFile(f)
		if err != nil {
			log.Tfatalf(conf.Tracer, "reading sql file error : %s", f)
		}
		_, err = db.Exec(string(sql))
		if err != nil {
			log.Fatal("sql error:", err)
		}
	}
}
