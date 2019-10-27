package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	h "github.com/NandanSatheesh/go-lang-playground/url-shortener/api"
	rr "github.com/NandanSatheesh/go-lang-playground/url-shortener/repository/redis"

	"github.com/NandanSatheesh/go-lang-playground/url-shortener/shortener"
)

// https://www.google.com -> 98sj1-293
// http://localhost:8000/98sj1-293 -> https://www.google.com

// repo <- service -> serializer  -> http

func main() {

	repo := getRedisRepository()
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8080")
		errs <- http.ListenAndServe(httpPort(), r)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func httpPort() string {
	port := "8080"
	return fmt.Sprintf(":%s", port)
}

func getRedisRepository() shortener.RedirectRepository {

	redisURL := "redis-13858.c8.us-east-1-4.ec2.cloud.redislabs.com:13858"
	redisPassword := "0OhEKCh7XLllhES2aO1jU79EEvdRNWRf"
	repo, err := rr.NewRedisRepository(redisURL, redisPassword)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
