package cmdrun

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
)

var FMT = fmt.Sprintf

func CMD(order string) (string, error) {
	log.Tinfof(conf.Tracer, "CMD: %s", order)
	// cmd := exec.Command("/usr/bin/script", "-f", "-e", "-q", "-c", order, "ci.log")
	// cmd := exec.Command("/usr/bin/script ", order)
	cmd := exec.Command("sh")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	in := bytes.NewBuffer(nil)
	cmd.Stdin = in

	in.WriteString(order)
	err := cmd.Run()
	if err != nil {
		log.Infof(conf.Tracer, "%s --> %s, CMD STDERR --> %s\n", order, err.Error(), stderr.String())
		return stderr.String(), err
	}
	return stdout.String(), nil
}
