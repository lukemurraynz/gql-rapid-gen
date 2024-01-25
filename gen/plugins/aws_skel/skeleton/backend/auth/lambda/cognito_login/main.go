package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
	"lib/data"
	"log"
	"strings"
)

func handler(ctx context.Context, event events.CognitoEventUserPoolsPostAuthentication) (ret events.CognitoEventUserPoolsPostAuthentication, err error) {
	fmt.Printf("Received event %s\n", spew.Sdump(event))

	email := strings.TrimSpace(strings.ToLower(event.Request.UserAttributes["email"]))
	phone := event.Request.UserAttributes["phone_number"]

	userProvider := data.GetUserProvider()

	user, err := userProvider.GetUser(ctx, strings.ToLower(event.UserName))
	if err != nil {
		return ret, fmt.Errorf("failed retrieving user: %w", err)
	}
	if user == nil {
		log.Printf("Creating user")
		user = &data.User{
			Id: event.UserName,
		}
		if email != "" {
			user.Email = &email
		}
		if phone != "" {
			user.PhoneNumber = &phone
		}
		err = userProvider.PutUser(ctx, user)
		if err != nil {
			return ret, fmt.Errorf("failed creating user: %w", err)
		}
	} else {
		changed := false
		if email != "" && (user.Email == nil || *user.Email == "") {
			log.Println("Need to add email")
			user.Email = &email
			changed = true
		}
		if phone != "" && (user.PhoneNumber == nil || *user.PhoneNumber == "") {
			log.Println("Need to add phone")
			user.PhoneNumber = &phone
			changed = true
		}
		if changed {
			// TODO move to an UpdateItem call to avoid races
			spew.Dump(user)
			err = userProvider.PutUser(ctx, user)
			if err != nil {
				return ret, fmt.Errorf("failed updating user: %w", err)
			}
		}
	}

	return event, nil

}

func main() {
	data.InitProvidersDynamoDB()
	lambda.Start(handler)
}
