package build

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
)

var FMT = fmt.Sprintf
var GoPath string
var CurentPath string

func init() {
	log.SetFlags(log.LstdFlags | log.Llevel)
	var err error
	CurentPath, err = os.Getwd()
	if err != nil {
		log.Tfatal(conf.Tracer, err)
	}
	command := conf.Config.InitCommand
	if err := os.Chdir(conf.Config.ProjectPath); err != nil {
		log.Tfatalf(conf.Tracer, "cd %s Error ", conf.Config.ProjectPath)
	}
	_, err = CMD(command)
	if err != nil {
		log.Tfatalf(conf.Tracer, "Run %s Error", command)
	}
	GoPath = os.Getenv("GOBIN")
	if GoPath == "" {
		log.Tfatal(conf.Tracer, "GOBIN is empty")
	}

}

func CMD(order string) (string, error) {
	log.Tinfof(conf.Tracer, "CMD: %s", order)
	cmd := exec.Command("bash")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	in := bytes.NewBuffer(nil)
	cmd.Stdin = in

	in.WriteString(order)
	err := cmd.Run()
	if err != nil {
		log.Infof(conf.Tracer, "%v --> %v, CMD STDERR --> %v\n", order, err.Error(), stderr.String())
		log.Infof(conf.Tracer, "Stdout: %s", stdout.String())
		return stderr.String(), err
	}
	return stdout.String(), nil
}

//BuildApp build app
//GoRoutine To Build App, 火力全开
func BuildApp() (string, error) {
	// 开启最大的CPU并行。计时开始
	t1 := time.Now().UnixNano()
	log.Info(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	services := conf.Config.Services
	apps := make(chan string, len(services))
	f := func(name, cmd string) {
		_, err := CMD(cmd)
		if err != nil {
			log.Tfatalf(conf.Tracer, "Run %s Error", cmd)
		}
		apps <- name
	}
	for _, s := range services {
		go f(s.Name, s.BuildCommand)
	}
	for i := 0; i < len(services); i++ {
		<-apps
	}
	t2 := time.Now().UnixNano()
	time := string(t2 - t1)
	return time, nil
}
