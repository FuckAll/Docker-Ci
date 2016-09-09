package webhook

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/FuckAll/Docker-Ci/ci"
)

type Payload struct {
	Ref     string
	Compare string
	Commits []Commit
}

type Commit struct {
	Id      string
	Message string
}

func HookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading the request body. %+v\n", err)

	}

	var payLoad map[string]interface{}
	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "json") {
		decoder := json.NewDecoder(strings.NewReader(string(body)))
		decoder.UseNumber()

		err := decoder.Decode(&payLoad)
		if err != nil {
			log.Printf("error parsing JSON payload %+v\n", err)
		}
	} else if strings.Contains(contentType, "form") {

		fd, err := url.ParseQuery(string(body))
		if err != nil {
			log.Printf("error parsing form payload %+v\n", err)
		} else {
			payLoad = valuesToMap(fd)
		}
	}

	// commit Message中出现dockerci则进行Ci操作

	var payLoads Payload
	payloadMap := payLoad["payload"].(string)
	err = json.Unmarshal([]byte(payloadMap), &payLoads)
	if err != nil {
		log.Printf("error Unmarshal. %+v\n", err)
	}
	if strings.Contains(payLoads.Commits[0].Message, "add") {
		ci.CiRun("TestClean", "")
	}
}

