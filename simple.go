package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go - Hello World</h1>")
}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	fmt.Println("simple test")
	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		panic(err)
	}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	actionCommendArg := "start" 
	if len(os.Args)>=2 {
		actionCommendArg = os.Args[1]
	} 

	svcConfig := &service.Config{
		Name:        "GoServiceTest",
		DisplayName: "Go Service Test",
		Description: "This is a test Go service.",
		// Executable: "simple.exe",
		// Arguments: []string{"service", "start"},
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	switch actionCommendArg {
	case "install":
		err := s.Install()
		if err != nil {
			logger.Error("Fail to install service", err)
		} else {
			fmt.Println("Succeed to install service")
		}
		os.Exit(0)
	case "uninstall":
		s.Stop()
		err := s.Uninstall()
		if err != nil {
			logger.Error("Fail to uninstall service", err)
		} else {
			fmt.Println("Succeed to uninstall service")
		}
		os.Exit(0)
	case "start": 
		err := s.Start()
		if err != nil {
			logger.Error("Failed to start service: ", err)
			fmt.Println("Failed to start service: ", err)
		} else {
			fmt.Println("Succeed to start service go-supervisord")
		}
		os.Exit(0)
	case "stop":
		// fmt.Println("Run stop function")
		err := s.Stop()
		if err != nil {
			logger.Error("Failed to stop service: ", err)
			fmt.Println("Failed to stop service: ", err)
		} else {
			fmt.Println("Succeed to stop service go-supervisord")
		}
		os.Exit(0)
	case "run":
		// Run
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	default:
		os.Exit(0)
	}
}