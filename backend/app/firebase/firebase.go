package firebase

import (
	"context"

	"firebase.google.com/go/auth"
	adf "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/firebase"
)

func GetAuthToken(token string) (*auth.Token, error) {
	authToken, err := adf.FirebaseClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, err
	}

	return authToken, nil
}

func AuthByUid(uid string) (*auth.UserRecord, error) {
	user, err := adf.FirebaseClient.GetUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUIDFromToken(token string) (string, error) {
	authToken, err := GetAuthToken(token)
	if err != nil {
		return "", err
	}

	return authToken.UID, nil
}

/*
	 func CreateUserInFb(userParams adm.CreateUserParams) (string, error) {

		ctx := context.Background()

		newUser := &auth.UserToCreate{}

		newUser.Email(userParams.UEmail).Password(userParams.UPassword)

		fbUser, err := adf.FirebaseClient.CreateUser(ctx, newUser)
		if err != nil {
			return "", err
		}

		customToken, err := adf.FirebaseClient.CustomToken(ctx, fbUser.UID)
		if err != nil {
			return "", err
		}

		return customToken, nil
	}
*/
func GetUIDFromEmail(email string) (string, error) {
	user, err := adf.FirebaseClient.GetUserByEmail(context.Background(), email)
	if err != nil {
		return "", err
	}

	return user.UID, nil
}

func DeleteUser(uid string) error {
	ctx := context.Background()

	err := adf.FirebaseClient.DeleteUser(ctx, uid)
	if err != nil {
		return err
	}

	return nil
}
