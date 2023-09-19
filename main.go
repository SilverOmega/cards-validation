package cards_validation

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http"
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

func CatFacts(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://cat-fact.herokuapp.com/facts/random?animal_type=cat&amount=29")
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
