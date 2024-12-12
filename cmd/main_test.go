package main

import (
	"testing"

	emailverifier "github.com/AfterShip/email-verifier"
)

func TestEmailVerif(t *testing.T) {
	v := emailverifier.NewVerifier().EnableDomainSuggest()
	t.Run("martavoi.@gmail.com", func(t *testing.T) {
		ret, err := v.Verify("martavoi.@gmail.com")
		if err != nil {
			t.Errorf("Unable to verify: %v", err)
		}

		if ret.Syntax.Valid {
			t.Errorf("Syntax should be invalid for %s", ret.Email)
		}

		if len(ret.Suggestion) != 0 {
			t.Logf("Suggestion: %s -> %s", ret.Email, ret.Suggestion)
		}
	})

	t.Run("martavoi@gnail.com", func(t *testing.T) {
		ret, err := v.Verify("martavoi@gnail.com")
		if err != nil {
			t.Errorf("Unable to verify: %v", err)
		}

		if len(ret.Suggestion) != 0 {
			t.Logf("Suggestion: %s -> %s", ret.Email, ret.Suggestion)
		}
	})

	t.Run("martavoi@hotmail-com", func(t *testing.T) {
		ret, err := v.Verify("martavoi@hotmail-com")
		if err != nil {
			t.Errorf("Unable to verify: %v", err)
		}

		if ret.Syntax.Valid {
			t.Errorf("Syntax should be invalid for %s", ret.Email)
		}

		if len(ret.Suggestion) != 0 {
			t.Logf("Suggestion: %s -> %s", ret.Email, ret.Suggestion)
		}
	})
}
