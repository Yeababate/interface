package main
import (
	"fmt"
	"net/http"
	"html/template"
)
var tmpl = template.Must(template.ParseFiles("templates/index.html"))
type Data struct {
	Output string
	OutputStatus int
	Input string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("POST /art", codehandler)
	fmt.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
    tmpl.Execute(w, Data{})
}

func codehandler(w http.ResponseWriter, r *http.Request) {
	var status int

	input := r.FormValue("input")
	action := r.FormValue("action")
	var operation func (input string)(string, bool)
	
	if action == "decode" {
		operation = decoder
	}else if action == "encode" {
		operation = encoder
	}
	result, ok := operation(input)
	if !ok {
		result = ""
		status = http.StatusBadRequest
	}else {
		status = http.StatusAccepted
	}
	w.WriteHeader(status)
	tmpl.Execute(w, Data{Input: input, Output: result, OutputStatus: status})
}
