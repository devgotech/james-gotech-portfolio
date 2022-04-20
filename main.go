package main

import "net/http"

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.url.Path != "/html" {
		http.Error(w, "404 not found".http.StatusNotFound)
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./html"))
	http.Handle("/", fileServer)
	http.HandleFunc("/", homePage)
}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "index")
// }

// func handleRequests() {
// 	http.HandleFunc("/", homePage)
// 	log.Fatal(http.ListenAndServe(":8082", nil))
// }

// func main() {
// 	handleRequests()

// }
