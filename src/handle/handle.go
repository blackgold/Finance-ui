package handle

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func get404() string {
	str404 := `
	 _  _    ___  _  _
	| || |  / _ \| || |
	| || |_| | | | || |_
	|__   _| | | |__   _|
           | | | |_| |  | |
	   |_|  \___/   |_|`
	return str404
}

func Api(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.Method, r.URL)
	vars := mux.Vars(r)
	api := vars["key"]
	if api != "info" {
		w.Write([]byte(get404()))
	} else {
		data := `{
	  		"UP50dayavg" : {
	      			"AAPL","HPE"
	  		},
	  		"DOWN50dayavg" : {
	      			"CHK","HPE"
	  		}
		}`
		w.Write([]byte(data))
	}
}

func Page404(w http.ResponseWriter, r *http.Request) {
	log.Println( r.RemoteAddr, r.Method, r.URL)
	w.Write([]byte(get404()))
}