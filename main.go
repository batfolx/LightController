package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type AlexaResponse struct {
	Version  string `json:"version"`
	Response struct {
		OutputSpeech struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"outputSpeech"`
		ShouldEndSession bool `json:"shouldEndSession"`
	} `json:"response"`

	
}


type AlexaRequest struct {
	Version string `json:"version"`
	Request struct {
		Type   string `json:"type"`
		Time   string `json:"timestamp"`
		Intent struct {
			Name               string `json:"name"`
			ConfirmationStatus string `json:"confirmationstatus"`
		} `json:"intent"`
	} `json:"request"`
}



func CreateAlexaResponse() AlexaResponse {

	 var a AlexaResponse
	 a.Version = "1.0"
	 a.Response.OutputSpeech.Type = "PlainText"
	 a.Response.OutputSpeech.Text = "Welcome default Victor!"
	 a.Response.ShouldEndSession = false
	 return a


}

func (a *AlexaResponse) Say(message string) {
	a.Response.OutputSpeech.Text = message
}





func HandleRequest(context context.Context, request AlexaRequest) (AlexaResponse, error) {

	alexa_response := CreateAlexaResponse()
	fmt.Printf("THIS IS THE REQUEST TYPE! -> %s <-\n", request.Request.Type)

	if request.Request.Type == "LaunchRequest" {
		alexa_response.Say("I am sorry for being here!")
	} else {

		switch request.Request.Intent.Name {
		case "HelloIntent":
			alexa_response.Say("Well hello there you little chicken!")
		default: 
			alexa_response.Say("I have absolutely no idea what you mean or will ever mean by that")
		}
	}
	
	return alexa_response, nil
}


func main() {
    lambda.Start(HandleRequest)
}