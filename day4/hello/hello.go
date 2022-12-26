package hello

import "net/http"

func Hello() {
	http.ListenAndServe("localhost:8000", nil)
	http.HandleFunc("/hello", helloHandler)
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello"))
}
