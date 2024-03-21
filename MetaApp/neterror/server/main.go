package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	host := flag.String("host", "127.0.0.1", "listen host")
	port := flag.String("port", "8088", "listen port")

	http.HandleFunc("/hello", Hello)

	err := http.ListenAndServe(*host+":"+*port, nil)

	if err != nil {
		panic(err)
	}
}

func Hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
	fmt.Printf(" req.header=%+v,\n req.usr.string=%+v,\n req.url.path=%+v\n", req.Header, req.URL.String(), req.URL.Path)
	req.ParseForm()
	d := req.Form
	fmt.Println(d)
}
