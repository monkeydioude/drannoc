package body

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// LoginData struct embodies data received from gin.Context.Request.Body
type LoginData struct {
	Login    string `json:"l"`
	Password string `json:"p"`
}

// IsValid verifies that LoginData were parsed and set correctly
func (ldata LoginData) IsValid() bool {
	return ldata.Login != "" && ldata.Password != ""
}

// NewLoginData should return a LoginData struct from a gin.Context.Request.Body
func NewLoginData(body io.ReadCloser) (LoginData, error) {
	data, err := ioutil.ReadAll(body)
	loginData := LoginData{}
	if err != nil {
		return loginData, err
	}

	err = json.Unmarshal(data, &loginData)
	if err != nil {
		return loginData, err
	}

	return loginData, nil
}
