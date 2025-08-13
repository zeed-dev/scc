package utils

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("JWT_SECRET", "test-secret")
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
	code := m.Run()
	os.Unsetenv("JWT_SECRET")
	os.Exit(code)
}

func TestGenerateAndValidateJWT(t *testing.T) {
	token, err := GenerateJWT(1, "test@example.com", "admin")
	if err != nil {
		t.Fatalf("GenerateJWT returned error: %v", err)
	}
	if token == "" {
		t.Fatalf("GenerateJWT returned empty token")
	}

	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("ValidateJWT returned error: %v", err)
	}
	if claims.UserID != 1 || claims.Email != "test@example.com" || claims.Role != "admin" {
		t.Fatalf("claims mismatch: %+v", claims)
	}
}

func TestValidateJWTMalformedToken(t *testing.T) {
	if _, err := ValidateJWT("malformed.token"); err == nil {
		t.Fatalf("expected error for malformed token")
	}
}
