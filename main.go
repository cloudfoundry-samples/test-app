package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jbayer/lattice-app/handlers"
	"github.com/jbayer/lattice-app/helpers"
	"github.com/jbayer/lattice-app/routes"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/rata"
)

var message string
var quiet bool

func init() {
	flag.StringVar(&message, "message", "Hello", "The Message to Log and Display")
	flag.BoolVar(&quiet, "quiet", false, "Less Verbose Logging")
	flag.Parse()
}

func main() {
	logger := lager.NewLogger("lattice-app")
	if quiet {
		logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.INFO))
	} else {
		logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("lattice-app.starting", lager.Data{"port": port})
	handler, err := rata.NewRouter(routes.Routes, handlers.New(logger))
	if err != nil {
		logger.Fatal("router.creation.failed", err)
	}

	index, err := helpers.FetchIndex()
	appName := fetchAppName()
	go func() {
		t := time.NewTicker(time.Second)
		for {
			<-t.C
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to fetch index: %s\n", err.Error())
			} else {
				fmt.Println(fmt.Sprintf("%s. Says %s. on index: %d", appName, message, index))
			}
		}
	}()

	server := ifrit.Envoke(http_server.New(":"+port, handler))
	logger.Info("lattice-app.up", lager.Data{"port": port})
	err = <-server.Wait()
	if err != nil {
		logger.Error("farewell", err)
	}
	logger.Info("farewell")
}

func fetchAppName() string {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		return "Lattice-app"
	}
	return appName
}
