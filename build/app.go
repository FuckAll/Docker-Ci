package build

import (
	"bytes"
	"fmt"
	"net/http"
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
	http.Request
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

// CMD Used To Do Some Shell Command
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

// BuildApp Used To Build App
// GoRoutine To Build App,
// Create 3 Builder To Build App Use Goroutine
func BuildApp() (string, error) {
	// 开启最大的CPU并行。计时开始
	t1 := time.Now().UnixNano()
	runtime.GOMAXPROCS(runtime.NumCPU())
	services := conf.Config.Services

	// 创建信道，长度为App个数,并且把信道中填充
	apps := make(chan string, len(services))
	for _, s := range services {
		apps <- s.BuildCommand

	}

	// 关闭channel,只读不可写
	close(apps)

	// 开启Builder的数量,默认为3
	builderNum := 3
	complete := make(chan bool, builderNum)
	for i := 0; i < builderNum; i++ {
		go builder(apps, complete)

	}
	for i := 0; i < builderNum; i++ {
		<-complete
	}

	//创建goroutine
	t2 := time.Now().UnixNano()
	time := string(t2 - t1)
	return time, nil
}

// builder Used By BuildApp
func builder(apps chan string, complete chan bool) {
	for cmd := range apps {
		_, err := CMD(cmd)
		if err != nil {
			log.Tfatalf(conf.Tracer, "Run %s Error", cmd)
		}
		if len(apps) <= 0 {
			complete <- true
		}
	}

}
