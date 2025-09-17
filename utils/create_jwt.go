package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub          int `json:"sub"`
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	IsShopeOwner bool   `json:"isShopeOwner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	// Header encode
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := base64UrlEncode(byteArrHeader)

	// Payload encode
	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadB64 := base64UrlEncode(byteArrData)

	// Message
	message := headerB64 + "." + payloadB64

	// Signature
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)

	// Final JWT
	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt, nil
}

// Base64 URL encode without padding
func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
