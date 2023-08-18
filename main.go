//Home page func
//an blog page
//handle req function
//main func - entry

package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
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

func handleRequests(){
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/home/blogPage", blogPage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequests()
}