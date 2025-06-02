package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitializeFirebase() {
	credentialFiles := os.Getenv("FIREBASE_CREDENTIALS")

	opt := option.WithCredentialsFile(credentialFiles)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing Firebase app: %v\n", err)
	}

	FirebaseApp = app
	log.Println("✅ Firebase initialized successfully")
}
