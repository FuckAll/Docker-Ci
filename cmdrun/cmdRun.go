package cmdrun

import(
    "bytes"
    "fmt"
    "os/exec"
    
    
    "github.com/wothing/log"
    "github.com/FuckAll/Docker-Ci/conf"
)

var FMT = fmt.Sprintf


func CMD(order string) string{
    log.Tinfo(conf.Tracer, "CMD: %s", order)
    cmd := exec.Command("/usr/bin/script", "-f", "-e", "-q", "-c", order, "ci.log")
    
    var stdout bytes.Buffer
    cmd.Stdout = &stdout
    
    var stderr bytes.Buffer
    cmd.Stderr = &stderr
    
    err := cmd.Run()
    if err != nil{
        log.Tfatal(conf.Tracer, "%s -- > %s, CMD STDERR --> %s\n", order, err.Error, stderr.String())
    }
    return stdout.String()
}
