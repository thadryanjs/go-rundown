package main

import (
    "fmt"
    "net/http"
    "github.com/a-h/templ"
)

func main() {

    /* create pages */
    homePage := homePage()
    page1 := page1("But look, I made you some content!")
    page2 := page2("Daddy made you your favorite,")
    page3 := page3("Open wide!")
    page4 := page4("I have a button.")

    /* handle pages  */
    http.Handle("/", templ.Handler(homePage))
    http.Handle("/page1", templ.Handler(page1))
    http.Handle("/page2", templ.Handler(page2))
    http.Handle("/page3", templ.Handler(page3))
    http.Handle("/page4", templ.Handler(page4))

    /* serve */
    fmt.Println("Listening on :3000")
    http.ListenAndServe(":3000", nil)
}
