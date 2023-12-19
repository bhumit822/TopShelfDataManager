package clients

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	option "google.golang.org/api/option"
)

var ClientOption option.ClientOption
var App *firebase.App

func InitFirebase() {

	ClientOption = option.WithCredentialsFile("appKey.json")
	firebaseApp, err := firebase.NewApp(context.Background(), &firebase.Config{StorageBucket: "topshelf-d392c.appspot.com"}, ClientOption)

	if err != nil {
		fmt.Fprintln(os.Stdout, []any{"error initializing app: %v", err}...)
	}

	if err == nil {
		App = firebaseApp
	}

}
