package eightball

type ResponseType string

const (
	ResponseTypePositive ResponseType = "Positive"
	ResponseTypeNeutral  ResponseType = "Neutral"
	ResponseTypeNegative ResponseType = "Negative"
)

type Response struct {
	Message      string
	ResponseType ResponseType
}

var Responses []Response

var PositiveResponses = []Response{
	{
		Message:      "It is certain",
		ResponseType: ResponseTypePositive,
	},
	{
		Message:      "It is decidedly so",
		ResponseType: ResponseTypePositive,
	},
	{
		Message:      "Without a doubt",
		ResponseType: ResponseTypePositive,
	},
	{
		Message:      "Yes – definitely",
		ResponseType: ResponseTypePositive,
	},
	{
		Message:      "You may rely on it",
		ResponseType: ResponseTypePositive,
	},
	{
		Message:      "As I see it, yes",
		ResponseType: ResponseTypePositive,
	},
	{
		Message:      "Most likely",
		ResponseType: ResponseTypePositive,
	},
	{
		Message:      "Outlook good",
		ResponseType: ResponseTypePositive,
	},
	{
		Message:      "Yes",
		ResponseType: ResponseTypePositive,
	},
	{
		Message:      "Signs point to yes",
		ResponseType: ResponseTypePositive,
	},
}
var NeutralResponses = []Response{
	{
		Message:      "Reply hazy, try again",
		ResponseType: ResponseTypeNeutral,
	},
	{
		Message:      "Ask again later",
		ResponseType: ResponseTypeNeutral,
	},
	{
		Message:      "Better not tell you now",
		ResponseType: ResponseTypeNeutral,
	},
	{
		Message:      "Cannot predict now",
		ResponseType: ResponseTypeNeutral,
	},
	{
		Message:      "Concentrate and ask again",
		ResponseType: ResponseTypeNeutral,
	},
}
var NegativeResponses = []Response{
	{
		Message:      "Don’t count on it",
		ResponseType: ResponseTypeNegative,
	},
	{
		Message:      "My reply is no",
		ResponseType: ResponseTypeNegative,
	},
	{
		Message:      "My sources say no",
		ResponseType: ResponseTypeNegative,
	},
	{
		Message:      "Outlook not so good",
		ResponseType: ResponseTypeNegative,
	},
	{
		Message:      "Very doubtful",
		ResponseType: ResponseTypeNegative,
	},
}

func init() {
	Responses = make([]Response, 0, len(PositiveResponses)+len(NeutralResponses)+len(NegativeResponses))
	Responses = append(Responses, PositiveResponses...)
	Responses = append(Responses, NeutralResponses...)
	Responses = append(Responses, NegativeResponses...)
}
