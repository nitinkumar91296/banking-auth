package domain

import "github.com/dgrijalva/jwt-go"

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"

type Claims struct {
	CustomerId string   `json:"customer_id"`
	Accounts   []string `json:"accounts"`
	Username   string   `json:"username"`
	Expiry     int64    `json:"exp"`
	Role       string   `json:"role"`
}

func (c Claims) IsUserRole() bool {
	if c.Role == "user" {
		return true
	}
	return false
}

func BuildClaimsFromJwtMapClaims(mapClaims jwt.MapClaims) (*Claims, error) {
	return nil, nil
}
