package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/nunoki/gorest/internal/connstr"
	"github.com/nunoki/gorest/internal/gorest"
	"github.com/nunoki/gorest/internal/gorest/postgres"
)

const DefaultPayloadLimit = 5000 // bytes #default_payload_limit

var (
	ctx = context.Background()
)

func main() {
	// connect to database
	dbU, dbPW, db, dbH, dbP, dbSSL := connstr.FromEnv()
	connStr := postgres.ConnectionString(dbU, dbPW, db, dbH, dbP, dbSSL)
	pg, err := postgres.NewClient(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		pg.Close()
	}()

	port := getPort()
	plimit := getPayloadLimit()
	s := gorest.NewServer(pg, port, plimit, true)
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, s))
}

// getPort returns the port number that the app is specified to run on. It will try to read from
// the environment, or return "80" by default.
func getPort() string {
	port := os.Getenv("APP_PORT")
	if len(port) == 0 {
		port = "80"
	}
	return port
}

// getPayloadLimit returns the payload limit set in the environment
// If empty or negative, the default value will be returned.
func getPayloadLimit() int64 {
	ls := os.Getenv("PAYLOAD_BYTE_LIMIT")
	limit, err := strconv.Atoi(ls)

	if err != nil || limit < 0 {
		limit = DefaultPayloadLimit
	}

	return int64(limit)
}
