package content

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
	"io"
	"os"
)

type HomePageCon struct {
	TabHeader  string `json:"TabHeader"`
	Header     string `json:"Header"`
	NavHeader1 string `json:"NavHeader1"`
	NavHeader2 string `json:"NavHeader2"`
	NavHeader3 string `json:"NavHeader3"`
	Heading1   string `json:"Heading1"`
	Heading2   string `json:"Heading2"`
	Heading3   string `json:"Heading3"`
	Heading4   string `json:"Heading4"`
	Heading5   string `json:"Heading5"`
	Heading6   string `json:"Heading6"`
	Par1       string `json:"Par1"`
	Par2       string `json:"Par2"`
	Par3       string `json:"Par3"`
	Par4       string `json:"Par4"`
	Par5       string `json:"Par5"`
	Par6       string `json:"Par6"`
	Muted1     string `json:"Muted1"`
	Muted2     string `json:"Muted2"`
	Muted3     string `json:"Muted3"`
	Img        string `json:"img"`
	Img2       string `json:"img2"`
	Img3       string `json:"img3"`
	Img4	   string `json:"img4"`
	Img5	   string `json:"img5"`	
	Img6	   string `json:"img6"`	

}


func HomePageTemplate(w http.ResponseWriter, r *http.Request)  {
	file, err := os.Open("./src/homepage.json")
	if err != nil {
		fmt.Println("Error :", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error :" ,err)
	}
	
	var content HomePageCon
	err = json.Unmarshal(byteValue, &content)
	
	if err != nil {
		fmt.Println("Error :" ,err)
	}
	

	temp ,err:= template.ParseFiles("./templates/index.html")
	
	if err != nil {
		fmt.Println("Check Template Files", err)
	}
	
    temp.Execute(w,content)

}