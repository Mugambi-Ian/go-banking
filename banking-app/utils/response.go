package utils

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func SendJSONResponse(res http.ResponseWriter, code int, data interface{}) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(code)
	if err := json.NewEncoder(res).Encode(data); err != nil {
		panic(err.Error())
	}
}

func SendXMLResponse(res http.ResponseWriter, code int, data interface{}) {
	res.Header().Add("Content-Type", "application/xml")
	res.WriteHeader(code)
	if err := xml.NewEncoder(res).Encode(data); err != nil {
		panic(err.Error())
	}
}
