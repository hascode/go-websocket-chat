package main

import (
	"log"
	"net/http"

	"./src/chat"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	for _, name := range []string{"arduino", "java", "groovy", "scala"} {
		r := chat.NewRoom(name)
		http.Handle("/chat/"+name, r)
		go r.Run()
	}

	log.Printf("starting chat server on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("starting server failed: ", err)
	}

}
