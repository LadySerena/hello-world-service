package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodPost:
		defer req.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(req.Body)
		parsedBody := map[string]string{}
		json.Unmarshal(bodyBytes, &parsedBody)
		fmt.Fprintf(w, "{\"message\":\"%s\"}", parsedBody["message"])
	default:
		fmt.Fprint(w, "{\"message\":\"hello world\"}")
	}

}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
