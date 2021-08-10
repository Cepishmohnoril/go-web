package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/get", get)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visits")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visits",
			Value: "0",
			Path:  "/",
		}
	}

	counter, _ := strconv.Atoi(cookie.Value)
	counter++
	cookie.Value = strconv.Itoa(counter)

	http.SetCookie(res, cookie)
	io.WriteString(res, "Visited!")
}

func get(res http.ResponseWriter, req *http.Request) {
	visits, err := req.Cookie("visits")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(res, "YOUR COOKIE:", visits)
}
