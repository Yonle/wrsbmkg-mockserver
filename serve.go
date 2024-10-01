package main

import (
	"net/http"
)

func main() {
	print("Listening on port 1945\n")
	print("In your code, Set <Penerima>.API_URL = \"http://localhost:1945\"")
	print("\n\nTo stop this server, Press CTRL + C.")
	http.ListenAndServe(":1945", http.FileServer(http.Dir("static")))
}
