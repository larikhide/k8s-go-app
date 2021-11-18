package main

import (
	"context"
	"fmt"
	"github.com/larikhide/k8s-go-app/server"
	"github.com/larikhide/k8s-go-app/version"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := "8080"

	info := server.VersionInfo{
		Version: version.Version,
		Commit: version.Commit,
		Build: version.Build,
	}

	srv := server.New(info, port)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := srv.Serve(ctx)
		if err != nil {
			log.Println(fmt.Errorf("server: %w", err))
			return
		}
	}()
	osSigChan := make(chan os.Signal, 1)
	signal.Notify(osSigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-osSigChan
	log.Println("OS interrupting signal has received")

	cancel()
	}