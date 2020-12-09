package encrypt

import "testing"

func TestICanGenerateHashedPasswd(t *testing.T) {
	goal := "7ae75468d284ad9401b8242be67790fd"
	if Password("une", "bonnepuree") != goal {
		t.Fail()
	}
}
