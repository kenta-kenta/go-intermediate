package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/auth/credentials/idtoken"
)

func main() {
	googleClientID := "[myClientID]"
	idToken := `[myIDToken]`

	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
	if err != nil {
		fmt.Println("validate error:", err)
		return
	}

	fmt.Println(payload.Claims["name"])
}
