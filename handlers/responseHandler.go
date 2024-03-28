package handlers

import (
	"apirest-gorm/models"
	"apirest-gorm/resources"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

func Send(w http.ResponseWriter, re models.Response, format string) {
	re.Message = resources.ErrorMessage
	w.WriteHeader(re.Status)
	if re.Status < http.StatusNotFound {
		re.IsSuccess = true
		re.Message = resources.SuccessMessage
	}
	re.StatusCode = http.StatusText(re.Status)
	if re.Data == nil{
		re.Data = struct{}{}
	}
	re.Metadata = []string{}
	var output []byte

	switch format {
	case "xml":
		output, _ = xml.Marshal(&re)
		w.Header().Set("Content-Type", "application/xml")
	case "yaml":
		output, _ = yaml.Marshal(&re)
		w.Header().Set("Content-Type", "application/yaml")
	default:
		output, _ = json.Marshal(&re)
		w.Header().Set("Content-Type", "application/json")
	}

	fmt.Fprintln(w, string(output))
}
