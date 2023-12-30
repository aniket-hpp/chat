package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	buff, err := os.ReadFile("./home.html")
	if err != nil {
		log.Printf("error: home.html > %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := string(buff)
	fmt.Fprintf(w, data, time.Now().UTC().String())
}
