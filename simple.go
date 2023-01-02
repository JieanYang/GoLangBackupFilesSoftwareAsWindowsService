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
	fmt.Println("os.Args: ",os.Args)
	fmt.Println("os.Args[0]: ",os.Args[0])
	fmt.Println("os.Args[1]: ",os.Args[1])
	fmt.Println("os.Args[2]: ",os.Args[2])

	actionFlag := os.Args[2]

	svcConfig := &service.Config{
		Name:        "GoServiceTest",
		DisplayName: "Go Service Test",
		Description: "This is a test Go service.",
		Executable: "simple.exe service start",
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

	switch actionFlag {
	case "install":
		err := s.Install()
		if err != nil {
			logger.Error("Fail to install service", err)
		} else {
			fmt.Println("Succeed to install service")
		}
	case "uninstall":
		s.Stop()
		err := s.Uninstall()
		if err != nil {
			logger.Error("Fail to uninstall service", err)
		} else {
			fmt.Println("Succeed to uninstall service")
		}
	case "start": 
		err := s.Start()
		if err != nil {
			logger.Error("Failed to start service: ", err)
			fmt.Println("Failed to start service: ", err)
		} else {
			fmt.Println("Succeed to start service go-supervisord")
		}
	case "stop":
		err := s.Stop()
		if err != nil {
			logger.Error("Failed to stop service: ", err)
			fmt.Println("Failed to stop service: ", err)
		} else {
			fmt.Println("Succeed to stop service go-supervisord")
		}
	default:
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
}