package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/maxfelker/openai-poc-api/src/handlers"

	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	openai "github.com/sashabaranov/go-openai"
)

func main() {

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}

	PORT := "8000"

	client := openai.NewClient(apiKey)
	router := mux.NewRouter()
	router.HandleFunc("/chat", handlers.ChatCompletion(client)).Methods("POST")

	corsObj := muxHandlers.CORS(muxHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
		muxHandlers.AllowedOrigins([]string{"*"}),
		muxHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))

	http.Handle("/", corsObj(router))

	fmt.Println("Starting openai-poc-api on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
