package main

import (
        "net/http"
    "log"
    "strconv"
	"time"
)

func main() {

	logMiddleware := func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request){
			log.Printf("Entering %s", request.URL)
			handlerFunc(writer, request)
			log.Printf("Exiting %s", request.URL)
		}

	}

    // health check
    healthy := true

    http.HandleFunc("/healthz", logMiddleware(func(writer http.ResponseWriter, request *http.Request) {

    	log.Println(healthy)
        if !healthy {
            writer.WriteHeader(http.StatusInternalServerError)
        } else {
            writer.WriteHeader(http.StatusOK)
        }
    }))

    // toggle health on or off
    http.HandleFunc("/togglehealth", logMiddleware(func(writer http.ResponseWriter, request *http.Request) {
       	healthy = !healthy

       	log.Printf("Health Status %s\n", healthy)

       	writer.Write([]byte(strconv.FormatBool(healthy)))
    }))

	http.HandleFunc("/ready", logMiddleware(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))

    // say hello
    http.HandleFunc("/", logMiddleware(func(writer http.ResponseWriter, request *http.Request) {

        writer.Header().Add("content-type", "text/html")
        writer.Write([]byte("<html><body><h1>Hello World</h1></body></html>"))
    }))

    // perform startup operations like connecting to DB, caching data, etc.
    log.Print("Startup up...")
    time.Sleep(10 * time.Second)

    log.Print("Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}