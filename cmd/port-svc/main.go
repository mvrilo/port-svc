package main

import (
	"context"
	"fmt"
	"log"
	gohttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/mvrilo/go-port-svc/adapter/http"
	"github.com/mvrilo/go-port-svc/adapter/inmemory"
	"github.com/mvrilo/go-port-svc/domain"
	"github.com/mvrilo/go-port-svc/usecase/portupsert"
)

func main() {
	_ = godotenv.Load()

	var conf domain.Config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err)
	}

	if conf.ServerAddress == "" {
		conf.ServerAddress = ":8000"
	}

	db := domain.PortMap{}
	store := inmemory.NewPortUpsertInMemoryStorage(db)
	svc := portupsert.New(store)

	handler := http.NewPortUpsertHandler(svc)

	mux := gohttp.NewServeMux()
	mux.Handle("/ports", handler)

	server := gohttp.Server{
		Addr:    conf.ServerAddress,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != gohttp.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	fmt.Println("Start port-svc")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	_ = server.Shutdown(ctx)
}
