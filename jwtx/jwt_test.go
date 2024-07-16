package jwtx

import (
	"fmt"
	"testing"
)

type testPayload struct {
	Id   int
	Name string
}

func TestJWT(t *testing.T) {
	jwt := NewJWT(JWTConfig{
		Method:     "HS256",
		Key:        "test+test",
		Expiration: 720,
	})

	jwtExpired := NewJWT(JWTConfig{
		Method:     "HS256",
		Key:        "test+test",
		Expiration: 0,
	})

	t.Run("sign and parse", func(t *testing.T) {
		sign, err := jwt.Sign(&testPayload{1, "hello_world"})
		if err != nil {
			t.Error(err)
		}

		payload := &testPayload{}
		err = jwt.Payload(sign, payload)
		if err != nil {
			t.Error(err)
		}

		if payload.Name != "hello_world" || payload.Id != 1 {
			t.Error("payload name or id is wrong")
		}

	})

	t.Run("verify expired token", func(t *testing.T) {
		sign, err := jwtExpired.Sign(&testPayload{1, "hello_world"})
		if err != nil {
			t.Error(err)
		}

		payload := &testPayload{}
		err = jwt.Payload(sign, payload)
		if err == nil {
			t.Errorf("expect error but got nil")
		}
		fmt.Println(err)
	})
}
