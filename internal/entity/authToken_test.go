package entity

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func init() {
	rand.Seed(3)
}

func TestICanGenerateAuthToken(t *testing.T) {
	token := GenerateAuthToken(time.Unix(0, 0), 10*time.Second, "", 5)
	fmt.Println(token)
	goal := &AuthToken{
		Created:  0,
		Expires:  10,
		LastUsed: 0,
		Duration: 10,
		Consumer: "",
		Life:     5,
		Token:    "cc2c8f0a084657a50fa37c58520b40cb",
	}
	if !reflect.DeepEqual(token, goal) {
		t.Fail()
	}
}

func TestICanVerifyTokenValidity(t *testing.T) {
	token := &AuthToken{
		Created:  0,
		Expires:  10,
		LastUsed: 0,
		Token:    "a",
	}
	if !token.IsValid(time.Unix(9, 0)) {
		t.Fail()
	}
}

func TestICanVerifyTokenExpired(t *testing.T) {
	token := &AuthToken{
		Created:  0,
		Expires:  10,
		LastUsed: 0,
		Token:    "b",
	}
	if token.IsValid(time.Unix(11, 0)) {
		t.Fail()
	}
}

func TestIShouldNotRemakeToken(t *testing.T) {
	token := &AuthToken{
		Created:  0,
		Expires:  10,
		LastUsed: 0,
		Token:    "d",
		Life:     1,
	}
	if token.ShouldRemake(time.Unix(15, 0)) {
		t.Fail()
	}
	if token.ShouldRemake(time.Unix(5, 0)) {
		t.Fail()
	}
}

func TestIShouldRemakeToken(t *testing.T) {
	token := &AuthToken{
		Created:  0,
		Expires:  10,
		LastUsed: 0,
		Token:    "d",
		Life:     1,
	}

	// inside remake threshold
	if !token.ShouldRemake(time.Unix(9, 0)) {
		t.Fail()
	}

	token.Life = 0
	// no more life but validity is still ok
	if !token.ShouldRemake(time.Unix(5, 0)) {
		t.Fail()
	}
}
