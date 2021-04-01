package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Firstname string             `bson:"Firstname"`
	Lastname  string             `bson:"Lastname"`
	email     string             `bson:"email"`
}
*/

var mongoStringCo string
var mongoDatabase string
var mongoCollection string
var wait time.Duration

func init() {
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	var mongoServer = flag.String("s", "localhost", "MongoDB address")
	var mongoPort = flag.String("p", "27017", "MongoDB port")
	mongoDatabase = *flag.String("d", "datasample", "MongoDB database")
	mongoStringCo = fmt.Sprintf("mongodb://%s:%s", *mongoServer, *mongoPort)
	mongoCollection = *flag.String("c", "users", "MongoDB collection")
}
func getall(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>
	<html>
	
	<head>
		<title>getall</title>
		<meta charset="utf-8">
	</head>
	
	<body>
	getall
	</body>
	
	</html>`))
}

func post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = vars["key"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>
	<html>
	
	<head>
		<title>post</title>
		<meta charset="utf-8">
	</head>
	
	<body>
	post
	</body>
	
	</html>`))
}

func get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = vars["key"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>
	<html>
	
	<head>
		<title>get</title>
		<meta charset="utf-8">
	</head>
	
	<body>
	get
	</body>
	
	</html>`))
}

func put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = vars["key"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>
	<html>
	
	<head>
		<title>put</title>
		<meta charset="utf-8">
	</head>
	
	<body>
	put
	</body>
	
	</html>`))
}

func patch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = vars["key"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>
	<html>
	
	<head>
		<title>patch</title>
		<meta charset="utf-8">
	</head>
	
	<body>
	patch
	</body>
	
	</html>`))
}
func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = vars["key"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>
	<html>
	
	<head>
		<title>delete</title>
		<meta charset="utf-8">
	</head>
	
	<body>
	delete
	</body>
	
	</html>`))
}

func main() {
	flag.Parse()

	// MongoDB connexion
	mctx, mcancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer mcancel()
	client, err := mongo.Connect(mctx, options.Client().ApplyURI(mongoStringCo))
	if err != nil {
		log.Println(err)
	}
	defer client.Disconnect(mctx)

	// collection := client.Database(mongoDatabase).Collection(mongoCollection)

	// Initialise Router
	r := mux.NewRouter()
	// Add your routes as needed
	r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(getall))).Methods("GET")
	r.Handle("/{key}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(post))).Methods("POST")
	r.Handle("/{key}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(get))).Methods("GET")
	// Reminder PUT method must include all the parameters of the record
	r.Handle("/{key}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(put))).Methods("PUT")
	// Reminder PATCH method must include only the parameter updated
	r.Handle("/{key}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(patch))).Methods("PATCH")
	r.Handle("/{key}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(delete))).Methods("DELETE")

	// Extra routes
	r.Handle("/workinprogress", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(workInProgress)))
	r.Handle("/healthz", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(healthz)))

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Server listen on :8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
