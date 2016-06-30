package build

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
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

func BuildApp() (string, error) {
	t1 := time.Now().UnixNano()
	services := conf.Config.Services

	for _, s := range services {
		_, err := CMD(s.BuildCommand)
		if err != nil {
			log.Tfatalf(conf.Tracer, "Run %s Error", s.BuildCommand)
			return "", err
		}
	}
	t2 := time.Now().UnixNano()
	time := string(t2 - t1)
	return time, nil

}
