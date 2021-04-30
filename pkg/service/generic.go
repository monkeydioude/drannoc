package service

import (
	"encoding/json"
	"io"
	"io/ioutil"

	pkgEntity "github.com/monkeydioude/drannoc/pkg/entity"
)

// EntityFromRequestBody instantiate an entity from a request's body
func EntityFromRequestBody(body io.ReadCloser, entity pkgEntity.Entity) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, entity)
	if err != nil {
		return err
	}

	return nil
}
