package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

func handler(ctx context.Context, event events.CognitoEventUserPoolsPreSignup) (ret events.CognitoEventUserPoolsPreSignup, err error) {
	fmt.Printf("Received event %s\n", spew.Sdump(event))

    // TODO implement

	return event, nil

}

func main() {
	lambda.Start(handler)
}
