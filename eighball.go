package eightball

import (
	"math/rand"
	"time"
)

func GetPositive() Response {
	return PositiveResponses[rand.Intn(len(PositiveResponses))]
}

func GetNeutral() Response {
	return NeutralResponses[rand.Intn(len(NeutralResponses))]
}

func GetNegative() Response {
	return NegativeResponses[rand.Intn(len(NegativeResponses))]
}

func GetRandom() Response {
	return Responses[rand.Intn(len(Responses))]
}

func AddMessage(message string, responseType ResponseType) {
	response := Response{
		Message:      message,
		ResponseType: responseType,
	}
	Responses = append(Responses, response)
	switch responseType {
	case ResponseTypePositive:
		PositiveResponses = append(PositiveResponses, response)
	case ResponseTypeNeutral:
		NeutralResponses = append(NeutralResponses, response)
	case ResponseTypeNegative:
		NegativeResponses = append(NegativeResponses, response)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
