package firebase

import (
	"context"
	"log"

	f "firebase.google.com/go"
	"firebase.google.com/go/auth"

	aasf "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/firebase"
	"google.golang.org/api/option"
)

var (
	FirebaseService aasf.FirebaseService
	FirebaseClient  *auth.Client
)

func FirebaseInit() {
	sa := option.WithCredentialsFile("./appli/datasources/firebase/firebase-service-account-key.json")
	firebaseNewApp, err := f.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalf("Firebase service unabled to start:%v", err)
		panic(err)
	}

	log.Println("Firebase service successfully started")
	FirebaseService.FirebaseApp = firebaseNewApp

	client, err := FirebaseService.FirebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("Unabled to get Firebase client :%v", err)
		panic(err)
	}
	FirebaseClient = client
	log.Println("Firebase client successfully initialized")
}
