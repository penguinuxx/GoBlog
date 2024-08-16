package content 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
	"io"
	"os"
)

type whoami struct {
	Par string `json:"Par"`
}


func WhoamiPageTemplate(w http.ResponseWriter, r *http.Request)  {
	file, err := os.Open("./src/whoami.json")
	if err != nil {
		fmt.Println("Error :", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error :" ,err)
	}
	
	var content whoami
	err = json.Unmarshal(byteValue, &content)
	
	if err != nil {
		fmt.Println("Error :" ,err)
	}
	

	temp ,err:= template.ParseFiles("./templates/whoami.html")
	
	if err != nil {
		fmt.Println("Check Template Files", err)
	}
	
    temp.Execute(w,content)

}