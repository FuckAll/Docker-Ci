//package buildlog
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"strings"
)

func showLog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filename := "/log/ci.log"
	//for k, v := range r.Form {
	//if k != "id" {
	//http.NotFound(w, r)
	//} else {
	//filename = "/log/" + strings.Join(v, "") + ".log"
	//}
	//}
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		http.NotFound(w, r)
	}
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		http.NotFound(w, r)
	}
	fmt.Fprintf(w, string(fd))

}

//func BuildLog() {
func main() {
	http.HandleFunc("/logs", showLog)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
