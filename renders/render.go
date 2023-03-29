package renders

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemp(w http.ResponseWriter, tempName string) {
	parsedfiles, _ := template.ParseFiles("D:/ShareX Captures/Go Lang/WorkSpace/templatesudemy/"+tempName, "D:/ShareX Captures/Go Lang/WorkSpace/templatesudemy/base.layout.tmpl")
	err := parsedfiles.Execute(w, nil)
	if err != nil {
		fmt.Println("Parsed files error:", err)
	}

}
