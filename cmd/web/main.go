package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := openDB(ctx, "mongodb://localhost:27017")
	if err != nil {
		errorLog.Fatal(err)
	}
	collection := client.Database("testing").Collection("numbers")
	collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})

	addr := flag.String("addr", ":4000", "HTTP network adress")
	flag.Parse()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Print(fmt.Printf("Starting server on %s", *addr))
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(ctx context.Context, url string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, url)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}
