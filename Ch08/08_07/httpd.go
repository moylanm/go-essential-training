/*
	Key/Value database

$ curl -d'hello' http://localhost:8080/k1
$ curl http://localhost:8080/k1
hello
$ curl -i http://localhost:8080/k2
404 not found

Limit value size to 1k
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Error string `json:"error"`
	Key	  string `json:"key"`
	Value []byte `json:"value"`
}

type Server struct {
	db DB
}

// POST /key Store request body as value
// GET /<key> Send back value, or 404 if key not found
func (s *Server) PostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	reader := io.LimitReader(r.Body, 1 << 10)
	key := r.PathValue("key")
	body, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	s.db.Set(key, body)
	log.Printf("set %s -> %v", key, body)
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	value := s.db.Get(key)
	response := Response{}

	if value == nil {
		response.Error = fmt.Sprintf("no such key [%s]", key)
	} else {
		response.Key = key
		response.Value = value
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Printf("cannot encode %v: %s", response, err)
	}
}

func main() {
	var server Server
	http.HandleFunc("GET /{key}", server.GetHandler)
	http.HandleFunc("POST /{key}", server.PostHandler)

	addr := "localhost:8080"
	log.Printf("starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
