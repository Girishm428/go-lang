package renders

import (
	"html/template"
	"net/http"
)

func RenderTemp(w http.ResponseWriter, tempName string) {
	parsedfiles, _ := template.ParseFiles("./templates/" + tempName)
	_ = parsedfiles.Execute(w, nil)
}
