package main

import (
	"net/http"
	"blog/content"

)
func main() {
	http.HandleFunc("/",content.HomePageTemplate)
	http.HandleFunc("/whoami",content.WhoamiPageTemplate)
	http.HandleFunc("/contents",content.ContentsTemplate)
	http.HandleFunc("/links", content.LinksPageTemplate)
	http.ListenAndServe(":8080",nil)
}