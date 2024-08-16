package content

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
	"io"
	"os"
)



type LinkTemplates struct {
	LinkTitle string `json:"LinkTitle"`
	Link string `json:"Link"`
	Png  string `json:"Png"`
}


func LinksPageTemplate ( w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./src/links.json")
	if err != nil {
		fmt.Println("Error :", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error :" ,err)
	}
	
	
	var Links []LinkTemplates
	
	err = json.Unmarshal(byteValue, &Links)
	
	if err != nil {
		fmt.Println("Error :" ,err)
	}
	

	temp ,err:= template.ParseFiles("./templates/links.html")
	
	if err != nil {
		fmt.Println("Check Template Files", err)
	}
	
    temp.Execute(w,Links)


}