package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
        "strconv"
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
			Slots interface{} `json:"slots"`
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

func (a *AlexaResponse) EndSession() {
	a.Response.ShouldEndSession = true
}









func LowerBrightnessIntent(a *AlexaRequest) string {
	
	amount, err := strconv.Atoi(a.Request.Intent.Slots.amount)
	if err != nil {
		return fmt.Sprintf("The amount you wanted was not valid. Amount was %s", a.Request.Intent.Slots.amount)
	}

	// code to lower the bulb brightness here

	return fmt.Sprintf("Bulb successfully lowered by %d amount!", amount)


}	


func HandleRequest(context context.Context, request AlexaRequest) (AlexaResponse, error) {

	alexa_response := CreateAlexaResponse()
	fmt.Printf("THIS IS THE REQUEST TYPE! -> %s <-\n", request.Request.Type)

	request_type := request.Request.Type

	if request_type == "LaunchRequest" {
		alexa_response.Say("Welcome back, Victor.")
	} else if request_type == "IntentRequest"{

		intent := request.Request.Intent.Name

		switch intent {
		case "HelloIntent":
			alexa_response.Say("Well hello there!")
		case "LowerBrightnessIntent":
			alexa_response.Say(LowerBrightnessIntent(&request))
		default: 
			alexa_response.Say("I have absolutely no idea what you mean or will ever mean by that")
		}
	}
	
	return alexa_response, nil
}


func main() {
    lambda.Start(HandleRequest)
}