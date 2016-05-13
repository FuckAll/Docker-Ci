package conf

import (
    "flag"
    "encoding/json"
    
    "io/ioutil"
    "github.com/wothing/log"
)

var (
    Tracer string

	Concurrent int

	REPO        string
	ProjectPath string // This is a absolute PATH
    Branch      string

	DockerRegistryPosition string
	REV                    string

	Services    []Service
	ServicesRun []Service //TODO use
    
    CGO_ENABLED string
    GOOS string
    GOARCH string
)

var (
    Push  = flag.Bool("p", false, "Push image to Registry")
    BuildList  = flag.String("b", "all", "Building list such as: appway, interway..... split by ,")
    TestOnly = flag.Bool("t", false, "after build")
)

type Service struct {
	Name string
	Path string
	Bin string
}

func init(){
    log.SetFlags(log.LstdFlags | log.Llevel)
    
    flag.Parse()
    data, err := ioutil.ReadFile("config.json")
    if err != nil{
        log.Tfatal(Tracer, "config.json unmarshal error: %v", err)
    }
    cm := make(map[string]interface{})
    err = json.Unmarshal(data, &cm)
    if err != nil{
       log.Tfatalf(Tracer, "config.json unmarshal error: %v", err) 
    }
    defer func(){
        if r:= recover(); r!=nil{
            log.Tfatal(Tracer, "config.json file illegal --> %v", r)
        }
    }()
    Concurrent = int(cm["Concurrent"].(float64))
	REPO = cm["REPO"].(string)
    Branch = cm["Branch"].(string)
	ProjectPath = cm["ProjectPath"].(string) // This is a absolute PATH
    
	DockerRegistryPosition = cm["DockerRegistryPosition"].(string)
    CGO_ENABLED = cm["CGO_ENABLED"].(string)
    GOOS = cm["GOOS"].(string)
    GOARCH = cm["GOARCH"].(string)
	services := cm["Services"].([]interface{})
	for _, v := range services {
		s := Service{
			Name: v.(map[string]interface{})["Name"].(string),
			Path: v.(map[string]interface{})["Path"].(string),
			Bin: v.(map[string]interface{})["Bin"].(string),
		}
		Services = append(Services, s)
	}

	log.Tinfo(Tracer, "load config.json succeed")
}

