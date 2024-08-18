package content

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
	"io"
	"os"
)

type ContentsTemplates struct {
	Topic       string `json:"Topic"`
	Posttitle   string `json:"Posttitle"`
	Url 		string `json:"Url"`
	Date        string `json:"Date"`
	ContextText string `json:"ContextText"`
}




func ContentsTemplate(w http.ResponseWriter, r *http.Request)  {
	file, err := os.Open("./src/contents.json")
	if err != nil {
		fmt.Println("Error :", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error :" ,err)
	}
	
	
	var Posts []ContentsTemplates

	err = json.Unmarshal(byteValue, &Posts)
	
	if err != nil {
		fmt.Println("Error :" ,err)
	}
	

	temp ,err:= template.ParseFiles("./templates/contents.html")
	
	if err != nil {
		fmt.Println("Check Template Files", err)
	}
	
    temp.Execute(w,Posts)

}