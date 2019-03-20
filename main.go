package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/JulzDiverse/feedelphia/api"
	"github.com/JulzDiverse/feedelphia/photobase"
)

const defaultPort = "8080"

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = defaultPort
	}

	photobase := photobase.NewInMemoryPhotobase()
	handler := api.New(&photobase)

	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
