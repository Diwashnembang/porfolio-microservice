package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
}

func main() {
	app := &application{}
	addr := flag.String("addr", ":4000", "HTTP Network address")
	if err := godotenv.Load(); err != nil {
		slog.Error("No .env file found", err)
	}
	uri := os.Getenv("MONGODB_URI")
	dsn := flag.String("dsn", uri, "mongoes connetction string")
	flag.Parse()
	//TODO make this  a function
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(*dsn))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	// Ping MongoDB to ensure the connection is alive
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Successfully connected to MongoDB")
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	err = srv.ListenAndServe()
	slog.Error(err.Error())

}

// func openDb(dsn string) *mongo.Client {

// }
