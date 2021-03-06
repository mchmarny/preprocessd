package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	ev "github.com/mchmarny/gcputil/env"
	pj "github.com/mchmarny/gcputil/project"
)

var (
	logger = log.New(os.Stdout, "", 0)
	prgID  = pj.GetIDOrFail()
	topic  = ev.MustGetEnvVar("TOPIC", "preprocessd")
	port   = ev.MustGetEnvVar("PORT", "8080")
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	http.HandleFunc("/v1/push", pushHandler)
	http.HandleFunc("/v1/api", apiHandler)

	hostPort := net.JoinHostPort("0.0.0.0", port)

	if err := http.ListenAndServe(hostPort, nil); err != nil {
		logger.Fatal(err)
	}

}
