package conf

import (
    "flag"
    "fmt"
    "encoding/json"
    
    "io/ioutil"
    "github.com/pborman/uuid"
    "github.com/wothing/log"
)

var (
    Tracer string

	Concurrent int

	REPO        string
	ProjectPath string // This is a absolute PATH
	SQLDir      string

	DockerRegistryPosition string
	REV                    string

	PGImage string

	RedisImage string

	ConsulImage string

	Services    []Service
	ServicesRun []Service //TODO use
)

var (
    Push  = flag.Bool("p", false, "Push image to Registry")
    BuildList  = flag.String("b", "all", "Building list such as: appway, interway..... split by ,")
    TestOnly = flag.Bool("t", false, "after build")
)

type Service struct {
	Name string
	Path string
	Para string
}

func init(){
    log.SetFlags(log.LstdFlags | log.Llevel)
    
    flag.Parse()
    Tracer = uuid.New()[:8]
    fmt.Println(Tracer)
    
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
	ProjectPath = cm["ProjectPath"].(string) // This is a absolute PATH
	SQLDir = cm["SQLDir"].(string)

	DockerRegistryPosition = cm["DockerRegistryPosition"].(string)

	//REV = cm["REV"].(string)

	PGImage = cm["PGImage"].(string)

	RedisImage = cm["RedisImage"].(string)

	ConsulImage = cm["ConsulImage"].(string)

	services := cm["Services"].([]interface{})
	for _, v := range services {
		s := Service{
			Name: v.(map[string]interface{})["Name"].(string),
			Path: v.(map[string]interface{})["Path"].(string),
			Para: v.(map[string]interface{})["Para"].(string),
		}
		Services = append(Services, s)
	}

	log.Tinfo(Tracer, "load config.json succeed")
}