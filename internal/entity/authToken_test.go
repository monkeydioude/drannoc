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
	token := NewAuthToken("could you be loved", 10*time.Second, time.Unix(0, 0))
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
