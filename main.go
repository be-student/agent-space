package main

import (
	"agent-harness/agent"
	"agent-harness/api"
	"agent-harness/cli"
	"flag"
	"log"
	"net/http"
)

func main() {
	address := flag.String("http", "127.0.0.1:8080", "HTTP API listen address")
	flag.Parse()
	cli.StartCli(func(assistant *agent.Agent) {
		handler := api.Handler(nil)
		if assistant != nil {
			handler = api.Handler(assistant)
		}
		go func() {
			if err := http.ListenAndServe(*address, handler); err != nil {
				log.Printf("HTTP API stopped: %v", err)
			}
		}()
	})
}
