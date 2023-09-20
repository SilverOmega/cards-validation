package cards_validation

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http"
	"time"
)

func main() {
	///add feature for link just use http.HandleFunc:
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/api/facts", CatFacts)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error createing server %s\n", err.Error())
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// net/http just support GET, POST, HEAD " http.Get, http.Post, http.Post"
type CatFact struct{}

func CatFacts(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://cat-fact.herokuapp.com/facts/random?animal_type=cat&amount=29")

	//use http.NewRequest to create new Request
	var client = &http.Client{}
	req, err := http.NewRequest("GET", "...", nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err = client.Do(req)
	// create Clienst and call method Get, Post, Head
	var clients = &http.Client{Timeout: 10 * time.Second}
	resp, err = clients.Get("...")
	if err != nil {
		log.Printf("Error getting data %s", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error parsing data %s", err.Error())
	}

	var catFacts []CatFact
	err = json.Unmarshal(body, &catFacts)
	if err != nil {
		log.Printf("Error parsing json %s", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(catFacts)
	if err != nil {
		log.Printf("Error parsing json %s", err.Error())
	}
}
