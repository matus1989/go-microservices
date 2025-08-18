package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct{
	Action string `json:"action"`
	Auth AuthPayload `json:"action,omitempty"`
	

}

type AuthPayload struct {

}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker"}
	_ = app.writeJSON(w, http.StatusOK, payload)
	
}

func (app *Config) HandlerSubmission(w http.ResponseWriter, r *http.Request){
	var RequestPayload RequestPayload
	err := app.readJSON(w,r, &RequestPayload)
	if err !=nil{
		app.errorJson(w,err)
		return
	}

	switch RequestPayload.Action{
	case "auth":
	default: 
	app.errorJson(w, errors.New("uknowo action"))

	}
}


func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload){
	// creae some json send to auth microsercvie
	jsonData, _:= json.MarshalIndent(a,"","\t")

	// call the service
	request, err := http.NewRequest("POST","http://authentication-service/authenticate",bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJson(w,err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)

		if err != nil {
		app.errorJson(w,err)
		return
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusUnauthorized{
		app.errorJson(w, errors.New("Invalid credentials"))
	}
	//make sure we get bacck correct status code
}