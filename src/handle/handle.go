package handle

import (
	"log"
	"net/http"
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

func Page404(w http.ResponseWriter, r *http.Request) {
	log.Println( r.RemoteAddr, r.Method, r.URL)
	w.Write([]byte(get404()))
}