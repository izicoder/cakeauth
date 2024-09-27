package core

import (
	"fmt"
	"io"
	"net/http"
)

func handleHello(resp http.ResponseWriter, req *http.Request) {
	response := `
	<body>
		<h1>Hello World</h1>
	</body>
	`
	io.WriteString(resp, response)
}

func handleIndex(resp http.ResponseWriter, req *http.Request) {
	clientIP := req.RemoteAddr
	response := `
	<body>
		Welcome<br>
		<a href="/hello">say hello</a><br>
	`
	fmt.Fprint(resp, response, clientIP, "</html>")
}

func StartServer() error {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/hello", handleHello)
	return http.ListenAndServe(":9999", nil)
}
