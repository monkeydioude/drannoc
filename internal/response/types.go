package response

import "github.com/monkeydioude/drannoc/internal/entity"

// UserIndex is an intermediary type for handling request/response
// with UserIndex endpoint
type UserIndex struct {
	User        *entity.User            `json:"user"`
	Preferences *entity.UserPreferences `json:"preferences"`
}
