package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r* http.Request) {
    fmt.Fprintln(w, "Welcom!")
}

func TodoIndex(w http.ResponseWriter, r* http.Request) {
    todos := Todos{
        Todo{ Name: "Foo" },
        Todo{ Name: "Bar" },
    }

    if err:= json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r* http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintf(w, "Todo show:", todoId)
}
