package main 

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"    
)

type Request struct {
    Name string
}

type Response struct {
    Greeting string
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func creator(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    request := Request{}
    response := Response{}
    json.NewDecoder(req.Body).Decode(&request)
    response.Greeting = "Hello, " + request.Name
    responseJson, _ := json.Marshal(response)
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(rw, "%s", responseJson)
}


func main() {
    mux := httprouter.New()    
    mux.GET("/hello/:name", hello)
    mux.POST("/hello", creator)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
