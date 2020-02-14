package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
        "strconv"
)

const (
	red = "red"
	green = "green"
	blue = "blue"
	yellow = "yellow"
	purple = "purple"
	orange = "orange"
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

			Slots map[string]struct { // map of strings to other structs

				Name string 	          `json:"name"`
				ConfirmationStatus string `json:"confirmationStatus,omitempty"`
				Value string    	      `json:"value"`	

			} `json:"slots"`
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



func HigherBrightnessIntent(a *AlexaRequest) string {

	response := ""
	slots := a.Request.Intent.Slots

	val := slots["amount"].Value; 

	if val == "" {
		val = "10"
	}

	value, err := strconv.Atoi(val)

	if err != nil {
		response = "I could not understand your value, please try again."
		return response
	}

	if value < 0 || value > 100 {
		response = fmt.Sprintf("How can I higher it by %d percent?", value)
		return response
	}


	response = fmt.Sprintf("Okay, highering bulb brightness by %d percent! Enjoy!", value)
	return response
}



func LowerBrightnessIntent(a *AlexaRequest) string {

	
	response := ""
	slots := a.Request.Intent.Slots

	value_str := slots["amount"].Value

	if value_str == "" {
		value_str = "10"
	}

	value, err := strconv.Atoi(value_str)

	if err != nil {
		response = fmt.Sprintf("I can't understand how much you want me to lower it by.")
		return response
	}

	if value < 0 {
		response = fmt.Sprintf("You can't lower the brightness below 0 are you crazy pills?")
		return response
	}

	if value > 100 {
		response = fmt.Sprintf("You can't lower the brightness more than 100 percent are you insane pills?")
		return response		
	}


	/* code to lower the bulb brightness here */


	response = fmt.Sprintf("Okay! Lowering brightness by %d percent!", value)

	return response


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
		case "HigherBrightnessIntent":
			alexa_response.Say(HigherBrightnessIntent(&request))
		default: 
			alexa_response.Say("I have absolutely no idea what you mean or will ever mean by that")
		}
	} else if request_type == "SessionEndedRequest" {
		alexa_response.Say("Goodbye! I will miss you.")
		alexa_response.EndSession()
	} else {
		alexa_response.Say("Huh?")
	}
	
	return alexa_response, nil
}


func main() {
    lambda.Start(HandleRequest)
}