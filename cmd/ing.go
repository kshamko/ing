package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/InVisionApp/go-health"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/kshamko/ing/internal/datasource"
	"github.com/kshamko/ing/internal/debug"
	"github.com/kshamko/ing/internal/handler"
	"github.com/kshamko/ing/internal/restapi"
	"github.com/kshamko/ing/internal/restapi/operations"
	"golang.org/x/sync/errgroup"
)

func main() { //nolint: funlen
	var opts = struct {
		HTTPListenHost string `long:"http.listen.host" env:"HTTP_LISTEN_HOST" default:"" description:"http server interface host"`
		HTTPListenPort int    `long:"http.listen.port" env:"HTTP_LISTEN_PORT" default:"8080" description:"http server interface port"`
		DebugListen    string `long:"debug.listen" env:"DEBUG_LISTEN" default:":6060" description:"Interface for serve debug information(metrics/health/pprof)"`
		Verbose        bool   `long:"v" env:"VERBOSE" description:"Enable Verbose log output"`
	}{}

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	log.SetLevel(log.ErrorLevel)

	if opts.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Infof("Launching Application with: %+v", opts)

	gr, appctx := errgroup.WithContext(context.Background())
	gr.Go(func() error {
		healthd := health.New()
		d := debug.New(healthd)

		return d.Serve(appctx, opts.DebugListen)
	})

	gr.Go(func() error {
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			return err
		}
		api := operations.NewSwaggerIngRoutesAPI(swaggerSpec)

		routesDS, err := datasource.NewRoutesOSRM()
		if err != nil {
			return err
		}
		api.RoutesRoutesHandler = handler.NewRoutes(
			routesDS,
		)

		server := restapi.NewServer(api)
		// nolint: errcheck
		defer server.Shutdown()

		server.Host = opts.HTTPListenHost
		server.Port = opts.HTTPListenPort

		return server.Serve()
	})

	errCanceled := errors.New("Canceled")

	gr.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		cusr := make(chan os.Signal, 1)
		signal.Notify(cusr, syscall.SIGUSR1)
		for {
			select {
			case <-appctx.Done():
				return nil
			case <-sigs:
				log.Info("Caught stop signal. Exiting ...")

				return errCanceled
			case <-cusr:
				if log.GetLevel() == log.DebugLevel {
					log.SetLevel(log.ErrorLevel)
					log.Info("Caught SIGUSR1 signal. Log level changed to INFO")

					continue
				}
				log.Info("Caught SIGUSR1 signal. Log level changed to DEBUG")
				log.SetLevel(log.DebugLevel)
			}
		}
	})

	err = gr.Wait()
	if err != nil && errors.Is(err, errCanceled) {
		log.Fatal(err)
	}
}
