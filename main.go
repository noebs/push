//Package push implements push notifications in cashq and noebs.
//
// How it works
// We collect data and analysis from through our network so it can inform our analysis regarding
// EBS reliability. Currently we are interested in these classes of errors:
// - Excessive number of errors (like more than 5 consecutive SYSTEM_ERROR, 196 or 696)
// - Helm specific errors
// - we also omit specific class of errors
// 		- insufficient funds
// - inelgiblile accounts
// - billers-down specific errors
// - and other non systematic errors
// there is all source of heuristics we are going to implement to ensure a reliable source of results.
package push

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"

	_ "embed"

	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

//go:embed cashq-273817-firebase-adminsdk-rtubx-d5799a788f.json
var f []byte

var app *firebase.App

func init() {
	var err error
	opt := option.WithCredentialsJSON(f)
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("Error in initializing firebase: %v", err)
	}
}

func NewMessage(title, body, topic string) error {

	// Obtain a messaging.Client from the App.
	ctx := context.Background()

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Printf("error getting Messaging client: %v\n", err)
		return err
	}

	// This registration token comes from the client FCM SDKs.
	// registrationToken := "dSBNe2BcQR68iLlGk6us7B:APA91bFk9ZC5Brjr7du_v3KPDyI2TANtdvhkqfP3bbyVjVn6M0Buoyfx-6u_7yrnZ3ko1YQVfEdqQmq4zWil5AHZoghq_pBC19-zbXpXGTgzkhUL2TGuvqRotk9FFY1zSc7IAuQb7nsv"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"title": title,
			"body":  body,
		},
		Topic: topic,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Print(err)
		return err
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	return nil
}
