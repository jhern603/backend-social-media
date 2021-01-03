package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//Posts Structs (Model)
type Posts struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

//Init book vars as slice
var posts []Posts

//Get All Posts
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w, "GET")

	json.NewEncoder(w).Encode(posts)
}

//Get Post
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	//Loop through books and return ID match
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Posts{})
}

//Create Post
func createPost(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, "POST")
	w.Header().Set("Content-Type", "application/json")
	var post Posts
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000)) //Mock ID
	posts = append(posts, post)
	json.NewEncoder(w).Encode(post)
}

//Update Book
func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range posts {
		if item.ID == params["id"] {
			//Slices out the old book
			posts = append(posts[:i], posts[i+1:]...)
			//Adds the new post
			var post Posts
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = strconv.Itoa(rand.Intn(1000)) //Mock ID
			posts = append(posts, post)
			json.NewEncoder(w).Encode(post)

		}
	}
	json.NewEncoder(w).Encode(posts)
}

//Delete Post
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, text/plain")
	enableCors(&w, "GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH")
	params := mux.Vars(r)
	for i, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:i], posts[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}

//Allow CORS
func enableCors(w *http.ResponseWriter, methods string) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", methods)
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, x-requested-with, Access-Control-Allow-Headers, Access-Control-Allow-Methods, Access-Control-Allow-Origin, Date, Content-Length")
}

func main() {
	//Init router
	router := mux.NewRouter()

	//Router handlers/endpoints
	router.HandleFunc("/api/posts", getPosts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/posts", createPost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", getPost).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", router))
}
