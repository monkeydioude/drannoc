package entity

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func init() {
	rand.Seed(3)
}

func TestICanGenerateANewAuthToken(t *testing.T) {
	token := NewAuthToken("could you be loved", time.Unix(0, 0), 10*time.Second)
	goal := &AuthToken{
		Created:  0,
		Expires:  10,
		LastUsed: 0,
		token:    "24e1dc2fae12912da7a56a87edd1b447",
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
		token:    "a",
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
		token:    "b",
	}
	if token.IsValid(time.Unix(11, 0)) {
		t.Fail()
	}
}

func TestIShouldRemakeToken(t *testing.T) {
	token := &AuthToken{
		Created:  0,
		Expires:  10,
		LastUsed: 0,
		token:    "c",
	}
	if !token.ShouldRemake(time.Unix(9, 0)) {
		t.Fail()
	}
}

func TestIShouldNotRemakeToken(t *testing.T) {
	token := &AuthToken{
		Created:  0,
		Expires:  10,
		LastUsed: 0,
		token:    "d",
	}
	if token.ShouldRemake(time.Unix(15, 0)) {
		t.Fail()
	}
}
