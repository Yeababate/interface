package main
import (
	"fmt"
	"net/http"
	"html/template"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))
type Data struct {
	OutputDecode string
	StatusDecode int

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
	r.ParseForm()
	input := r.FormValue("artToCode")
	action := r.FormValue("action")
	switch action {
	case "decode":
		result,ok := decoder(input)
		if !ok {
			result = ""
			status = http.StatusBadRequest
		}else {
			status = http.StatusAccepted
		}
		data := Data{OutputDecode: result, StatusDecode: status}
		fmt.Println(data)
		w.WriteHeader(status)
		tmpl.Execute(w, data)
	case "encode":
		result,ok := encoder(input)
		if !ok {
			result = ""
			status = http.StatusBadRequest
		}else {
			status = http.StatusAccepted
		}
		data := Data{OutputDecode: result, StatusDecode: status}
		fmt.Println(data)
		w.WriteHeader(status)
		tmpl.Execute(w, data)
	}
}
