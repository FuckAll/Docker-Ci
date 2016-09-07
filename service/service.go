package service

import "github.com/urfave/negroni"

func Start()  {

	n := negroni.New()


	n.Run()
}