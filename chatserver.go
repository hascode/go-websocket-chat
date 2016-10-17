package main

import (
	"log"
	"net/http"

	"./src/chat"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	roomNames := []string{"arduino", "java"}

	for _, name := range roomNames {
		r := chat.NewRoom(name)
		http.Handle("/chat/"+name, r)
		go r.Run()
	}

	log.Printf("starting server on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("starting server failed: ", err)
	}

}
