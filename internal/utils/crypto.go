package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with a plaintext password.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Generate a code challenge from a code verifier using the S256 method
func GenerateCodeChallengeS256(codeVerifier string) string {
	hash := sha256.Sum256([]byte(codeVerifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

// PKCE - checks if the code challenge matches the code verifier based on the given method.
// The method used for hashing (e.g., "S256") default to "plain" if not specified.
// Returns a boolean indicating if the code challenge matches the code verifier.
func VerifyCodeChallenge(codeVerifier, codeChallenge, method string) bool {
	if method == "S256" {
		return GenerateCodeChallengeS256(codeVerifier) == codeChallenge
	}
	// Plain method
	return codeVerifier == codeChallenge
}

// generates a secure, random authorization code.
func GenerateAuthorizationCode() (string, error) {
	// Create a byte slice with enough length
	code := make([]byte, 32) // 256 bits, which gives us a 43-character string in Base64 URL encoding

	// Read random bytes
	if _, err := rand.Read(code); err != nil {
		return "", err
	}

	// Encode to Base64 URL encoding (using URL encoding to make it safe for URLs)
	return base64.RawURLEncoding.EncodeToString(code), nil
}
