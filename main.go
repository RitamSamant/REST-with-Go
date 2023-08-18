//Home page func
//an blog page
//handle req function
//main func - entry
//Using gorilla mux

package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Blog struct {
	Title string
	Desc string
	Content string 
}

type Blogs []Blog

func homePage(w http.ResponseWriter, r *http.Request){

	//on webpage
	fmt.Fprintf(w, "Welcome to the HomePage!")

	//for terminal
	fmt.Println("Endpoint Hit: homePage")
}

func blogPage(w http.ResponseWriter, r *http.Request){

	blogs:=Blogs{
		Blog{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}
	//for terminal
	fmt.Println("Endpoint Hit: blogPage")
	json.NewEncoder(w).Encode(blogs);

}
func postBlogPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Post Blog Page")
	fmt.Println("Endpoint Hit: postBlogPage")
}

func handleRequests(){

	router:=mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/home", homePage).Methods("GET")
	router.HandleFunc("/home/blogPage", blogPage).Methods("GET")
	router.HandleFunc("/home/blogPage/add", postBlogPage).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main(){
	handleRequests()
}