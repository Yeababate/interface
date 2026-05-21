package main
import ( 
	"net/http"
	"html/template"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))
type Data struct {
	OutputDecode string
	StatusDecode int
	OutputEncode string
	StatusEncode int
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/decoder", decoderHandler)
	// http.HandleFunc("/encoder", encoderHandler)
	http.ListenAndServe(":8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, Data{})
}
func decoderHandler(w http.ResponseWriter, r *http.Request) {
	var status int
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	input := r.FormValue("encoded")
	action := r.FormValue("action")
	if action == "decode" {
		result,ok := decoder(input)
		if !ok {
			status = http.StatusBadRequest
		}else {
			status = http.StatusAccepted
		}
		data := Data{OutputDecode: result, StatusDecode: status}
		tmpl.Execute(w, data)
	}else if action == "encode" {
		result,ok := encoder(input)
		if !ok {
			status = http.StatusBadRequest
		}else {
			status = http.StatusAccepted
		}
		data := Data{OutputDecode: result, StatusDecode: status}
		tmpl.Execute(w, data)
	}
}
