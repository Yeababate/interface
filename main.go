package main
import (
	"net/http"
	"html/template"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))
type Data struct {
	Output string
	OutputStatus int
	ArtToCode string
	
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("POST /art", codehandler)
	http.ListenAndServe(":8080", mux)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, Data{})
}
func codehandler(w http.ResponseWriter, r *http.Request) {
	var status int

	input := r.FormValue("artToCode")
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
	data := Data{ArtToCode: input, Output: result, OutputStatus: status}
	w.WriteHeader(status)
	tmpl.Execute(w, data)

}
