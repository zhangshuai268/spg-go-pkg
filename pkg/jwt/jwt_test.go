package jwt

import (
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	type claim struct {
		UserId int `json:"user_id"`
		MyClaims
	}
	req := claim{
		UserId: 1,
	}
	token, err := GetToken("dadasdasd", &req)
	if err != nil {
		return
	}
	fmt.Println(token)

}

func TestParseToken(t *testing.T) {
	type claim struct {
		UserId int `json:"user_id"`
		MyClaims
	}
	req := claim{}
	err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.5u9FpJMCaJild6v1PoYtm-mnQCsKeBFzT-nX3Ck6VgQ", "dadasdasd", &req)
	if err != nil {
		return
	}
	fmt.Println(req.UserId)

}
