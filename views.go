package main

import (
	"bytes"
	"net/http"
	"html/template"
)

func parseTemplate(basefile string, file string, data map[string]interface{}) []byte {
	var buf bytes.Buffer

	t, err := template.ParseFiles(basefile, "templates/"+file)
	checkError(err)

	if data != nil {
		utils := &Utils{}
		data["utils"] = utils
	}

	err = t.Execute(&buf, data)
	checkError(err)

	return buf.Bytes()
}

func renderTemplate(w http.ResponseWriter, r *http.Request, file string, data map[string]interface{}) {
	page := parseTemplate("templates/base.html", file, data)
	w.Write(page)
}

func renderTemplate2(w http.ResponseWriter, r *http.Request, file string, data map[string]interface{}) {
	var buf bytes.Buffer

	t, err := template.ParseFiles("templates/"+file)
	checkError(err)

	if data != nil {
		utils := &Utils{}
		data["utils"] = utils
	}

	err = t.Execute(&buf, data)
	checkError(err)

	w.Write(buf.Bytes())
}
