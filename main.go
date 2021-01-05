// To be entirely honest, I’d recommend against fiddling with CORS at all.
//If your goal is to make back-end requests from a front-end (presumably running on different ports), I’d consider using something like webpack’s proxy server,
//which will proxy requests that make it to the front-end server to a back-end.
// Adding all this extra logic can go wrong if you forget to remove something once you actually make it to production.

//It’s pretty neat. It supports TLS with a self-signed cert as well.
//It essentially allows you to change absolutely nothing in your code; you just have to insert a few extra lines into your webpack config and everything works fine.

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
			posts = append(posts[:i], posts[i+1:]...)
			var post Posts
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(post)

		}
	}
	json.NewEncoder(w).Encode(posts)
}

//Delete Post
func deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:i], posts[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func main() {
	//Init router
	router := mux.NewRouter()

	//Router handlers/endpoints
	router.HandleFunc("/api/posts", getPosts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/posts", createPost).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/posts/{id}", getPost).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE", "OPTIONS")
	log.Fatal(http.ListenAndServe(":5000", router))
}
