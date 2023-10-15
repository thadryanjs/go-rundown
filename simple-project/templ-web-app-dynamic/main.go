package main

// TODO:
// organize into a templ folder and use recursive build to see if it works

import (
    "fmt"
    "net/http"
    "time"
)

type GlobalState struct {
    count int
}

// https://templ.guide/server-side-rendering/example-counter-application
var globalState GlobalState

func getHandler(w http.ResponseWriter, r *http.Request) {

    // initialize a page
    component := page(globalState.count)
    component.Render(r.Context(), w)

}

func postHandler(w http.ResponseWriter, r *http.Request) {

    // This should update the global state
    r.ParseForm()

    // if the form has a global field, increment the global state
    if r.Form.Has("increase") {
        globalState.count++
    } else if r.Form.Has("decrease") {
        globalState.count--
    }

    // pass the reader and writer to the get handler
    getHandler(w, r)

}

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        // if it's a post operation then handle it
        if r.Method == http.MethodPost {
            postHandler(w, r)
            return
        }
        // otherwise it's a get operation
        getHandler(w, r)

        textComponent("Here is text").Render(r.Context(), w)
        timeComponent(time.Now()).Render(r.Context(), w)

    })

    fmt.Println("Server is listening on port 3000...")
    http.ListenAndServe(":3000", nil)
}
