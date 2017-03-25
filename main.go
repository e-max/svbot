package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

const SECRET = "mutHupGhemopholOwdoocHuckmidBor4"

const BotToken = "760c6348f92cd0885737cf2e4a78dd0582915941"

func main() {
	fmt.Println("vim-go")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Printf("r.URL.Path = %+v\n", r.URL.Path)
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("body = %+v\n", string(body))
		event := map[string]interface{}{}
		err = json.Unmarshal(body, &event)
		if err != nil {
			log.Print(err)
		}
		res, err := json.MarshalIndent(event, "", "    ")
		if err != nil {
			log.Print(err)
		}
		fmt.Println(string(res))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
