package main

// TODO:
// https://templ.guide/server-side-rendering/example-counter-application
// organize into a templ folder and use recursive build to see if it works
j
import (
    "fmt"
    "net/http"
    "time"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        textComponent("Here is text").Render(r.Context(), w)
        timeComponent(time.Now()).Render(r.Context(), w)



    })

    fmt.Println("Server is listening on port 3000...")
    http.ListenAndServe(":3000", nil)
}
