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
	Title string `json:"Title`
	Desc string `json:"desc"`
	Content string `json:"content"`
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

	myRouter:=mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/home", homePage).Methods("GET")
	myRouter.HandleFunc("/home/blogPage", blogPage).Methods("GET")
	myRouter.HandleFunc("/home/blogPage/add", postBlogPage).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){
	handleRequests()
}