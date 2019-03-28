package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

//Reqeust is
type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

//Response is
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

//Handler lambda req handler
func Handler(req Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Processed the id : %f", req.ID),
		Ok:      true,
	}, nil
}

func main() {
	log.Println("Lambda Main Start....")
	lambda.Start(Handler)
	log.Println("Lambda Main End....")
}
