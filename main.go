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
	} `json:"response"`

	ShouldEndSession bool `json:"shouldEndSession"`
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
	 a.ShouldEndSession = false
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
		case "LaunchRequest":
			alexa_response.Say("Good evening you fucking bitch")
		case "SessionEndedRequest":
			alexa_response.Say("Goodbye you trog")
		default: 
			alexa_response.Say("I don't know what that means.")

		}
	}



	
	return alexa_response, nil
}


func main() {
    lambda.Start(HandleRequest)
}