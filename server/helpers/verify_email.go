package helpers

import (
	"fmt"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.NewVerifier().DisableCatchAllCheck()
)

func VerifyEmail(email string) (bool, error) {

	ret, err := verifier.Verify(email)
	if err != nil {
		return false, fmt.Errorf("Verification Failedddd: %v", err)
	}

	if !ret.Syntax.Valid {
		return false, fmt.Errorf("Verification failed: %v", "Syntax Invalid")
	}

	return true, nil
}
