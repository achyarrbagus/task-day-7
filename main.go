package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)
var Data = map[string]interface{}{
	"title": "Personal Web",
}


func main()  {

	route := mux.NewRouter()
	// crated static files
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/",http.FileServer(http.Dir("./public"))))



	route.HandleFunc("/helloworld",helloWorld).Methods("GET")
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/add-project",addProject).Methods("GET")
    route.HandleFunc("/contact-me",contact).Methods("GET")
	route.HandleFunc("/blog-content/{id}",blogContent).Methods("GET")
	

	

	fmt.Println("server running on port : 5000")
	http.ListenAndServe("localhost:5000",route)
}
// blog content handler
func blogContent(w http.ResponseWriter, r*http.Request){
	w.Header().Set("content-type","text-html;charset=utf-8")
	id, _ :=   strconv.Atoi(mux.Vars(r)["id"]) 

	 rsp := map[string]interface{}{
		"id": id,
		"data": Data,
	}

	var tmpl, err = template.ParseFiles("./views/blog-content.html")

	if  err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w,rsp)
}

// add project handler
func addProject(w http.ResponseWriter, r*http.Request){
	w.Header().Set("content-type","text-html;charset=utf-8")

	var tmpl , err = template.ParseFiles("./views/add-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w,Data)
}

// contact me handler 
func contact(w http.ResponseWriter, r*http.Request){
	w.Header().Set("content-type","text/html;charset=utf-8")

	 var tmpl , err = template.ParseFiles("./views/contact-me.html")
	 if  err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	 }
	 w.WriteHeader(http.StatusOK)
	 tmpl.Execute(w,Data)
	
}

// home handler
func home(w http.ResponseWriter,r*http.Request){
w.Header().Set("Content-Type" , "text/html; charset=utf-8")

var tmpl, err = template.ParseFiles("./views/home.html")
if err != nil {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("message :" + err.Error()))
	return
}
w.WriteHeader(http.StatusOK)
tmpl.Execute(w, Data)
}


//handler hello world
func helloWorld(w http.ResponseWriter,r*http.Request){
	w.Header().Set("content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))

	fmt.Println(http.StatusOK)
}