package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/api"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/infrastructure"
	"github.com/go-chi/chi"
)

func main() {
	routes := chi.NewRouter()
	r := api.RegisterRoutes(routes)
	log.Println("Server ready at 7000")
	log.Fatal(http.ListenAndServe(":7000", r))
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitMongoDB()
}
