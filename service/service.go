package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"github.com/FuckAll/Docker-Ci/service/webhook"
)

var (
	ip string = "127.0.0.1"
	port string = "4000"
	hookURL string = "/dockerci/{id}"
)

func Start() {

	l := negroni.NewLogger()
	webhookLog := log.New(os.Stderr, "[webhook] ", log.Ldate | log.Ltime)

	negroniRecovery := &negroni.Recovery{
		Logger:     webhookLog,
		PrintStack: true,
		StackAll:   false,
		StackSize:  1024 * 8,
	}

	n := negroni.New(negroniRecovery, l)

	router := mux.NewRouter()

	router.HandleFunc(hookURL, webhook.HookHandler)
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", ip, port), n))

}
