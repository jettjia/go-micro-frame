package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	SigningKey = "whowiewuwoo111gwoehwou"
)

// 获取token
func TestJWT_CreateToken(t *testing.T) {
	claims := jwt.MapClaims{
		"uuid":     111,
		"username": "www",
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}
	token, err := NewJWT(SigningKey).CreateToken(claims)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(token)
}

// 刷新token
func TestJWT_CreateTokenByOldToken(t *testing.T) {
	claims := jwt.MapClaims{
		"uuid":     111,
		"username": "www",
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}

	oldToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzU0OTg2NjcsInVzZXJuYW1lIjoid3d3IiwidXVpZCI6MTExfQ.MyM2oLRABePWZco_GmjzOFtCna9ClbejBOTHvlWLjdA"
	token, err := NewJWT(SigningKey).CreateTokenByOldToken(oldToken,claims)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(token)
}

// 解析token
func TestJWT_ParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzU1MTYyNDAsInVzZXJuYW1lIjoid3d3IiwidXVpZCI6MTExfQ.Awk7Bvik2ADLVwfrygkQRdKOjQIA9sVYabEdtI2uenY"
	mapClaims, err := NewJWT(SigningKey).ParseToken(tokenString)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(mapClaims)
}

// 解析token
func TestJWT_ParseTokenBearer(t *testing.T) {
	tokenString := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzU1MTYyNDAsInVzZXJuYW1lIjoid3d3IiwidXVpZCI6MTExfQ.Awk7Bvik2ADLVwfrygkQRdKOjQIA9sVYabEdtI2uenY"
	mapClaims, err := NewJWT(SigningKey).ParseTokenBearer(tokenString)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(mapClaims)
}