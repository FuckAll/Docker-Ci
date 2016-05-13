package ci

import (
	"bytes"
	"net/http"
	"time"

	"github.com/wothing/log"

	. "github.com/FuckAll/Docker-Ci/cmdrun"
	"github.com/FuckAll/Docker-Ci/conf"
)

var client = &http.Client{}

func Consul() {
	CMD(FMT("docker run -it -d --net=ci --name %s-consul %s agent -dev -bind=0.0.0.0 -client=0.0.0.0", conf.Tracer, conf.ConsulImage))

	for i := 0; ; i++ {
		if i > 30 {
			log.Tfatal("After for a long time we can't connection to consul")
		}

		if consulCheck() {
			log.Tinfof(conf.Tracer, "connection to consul success")
			for _, s := range append(conf.Services, conf.ServicesRun...) {
				consulRegister(s.Name, conf.Tracer+"-"+s.Name+".test")
			}
			break
		} else {
			log.Tinfof(conf.Tracer, "Try connection to consul %d time(s)", i+1)
			time.Sleep(time.Second)
		}
	}
}

func consulCheck() bool {
	url := "http://" + conf.Tracer + "-consul.test:8500/v1/agent/services"
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

func consulRegister(Name, Address string) {
	url := "http://" + conf.Tracer + "-consul.test:8500/v1/agent/service/register"

	var jsonStr = []byte(`{"Name":"` + Name + `", "Port": 3000, "Address":"` + Address + `" }`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		log.Tfatalf(conf.Tracer, "REG service error %s %s", Name, Address)
	}
}
