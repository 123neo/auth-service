package main

import (
	"auth-service/config"
	"auth-service/repository"
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// Server start for go

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	conn := connectToDB()

	if conn == nil {
		logger.Panic("Failed trying to connect Postgres...")
	}

	repo := repository.NewRepo(conn, logger)

	app := config.NewConfig(conn, repo, logger)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      routes(app),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// wrapping ListenAndServe in gofunc so it's not going to block
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.Error("Error:", zap.Any("Server start error", err))
		}
	}()

	// make a new channel to notify on os interrupt of server (ctrl + C)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// This blocks the code until the channel receives some message
	sig := <-sigChan
	logger.Info("Received terminate, graceful shutdown", zap.Any("channel", sig))

	// Once message is consumed shut everything down
	// Gracefully shuts down all client requests. Makes server more reliable
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

var counts int64

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)

		if err != nil {
			log.Println("Postgres is not ready....")
			counts++
		} else {
			log.Println("Connected to postgres....")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for 2 seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
