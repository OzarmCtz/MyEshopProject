package firebase_service

import (
	"time"

	f "firebase.google.com/go"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8082",
		Timeout: 200 * time.Millisecond,
	}
)

type FirebaseService struct{ FirebaseApp *f.App }
type firebaseServiceInterface interface{}
