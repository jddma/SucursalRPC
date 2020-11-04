package model

import (
	"context"
	"github.com/gorilla/securecookie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"

	"../serverUtilities/auth"
)

type Worker struct{
	Email string
	Role string
	Password string
}

func (w *Worker) Login(cookieHandler *securecookie.SecureCookie, wr http.ResponseWriter) bool {

	var result bool

	mgdUser := os.Getenv("MGD_USER")
	mgdPassword := os.Getenv("MGD_PASSWORD")
	mgdHost := os.Getenv("MGD_HOST")

	uri := "mongodb://" + mgdUser + ":" + mgdPassword + "@" + mgdHost + ":27017"

	connection, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = connection.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{
		"email", w.Email,
	}}

	collection := connection.Database("branch").Collection("workers")
	var queryResult map[string]string
	err = collection.FindOne(context.TODO(), filter).Decode(&queryResult)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(queryResult["password"]), []byte(w.Password))
	if err != nil {
		result = false
	} else {
		result = true
		auth.SetUserSession(w, cookieHandler, wr)
	}

	return result

}

func (w *Worker) GetIdentifier() string {

	return w.Email

}

func (w *Worker) GetPassword() string {

	return w.Password

}

func (w *Worker) GetRole() string {

	return w.Role

}