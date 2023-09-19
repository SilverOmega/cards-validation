package cards_validation

import (
	"fmt"
	_ "fmt"
	"log"
	"net/http"
	_ "net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error createing server %s\n", err.Error())
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
