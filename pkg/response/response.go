package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	log "github.com/sirupsen/logrus"
)

// Response is the endpoint's HTTP reponse defining struct
type Response struct {
	err  string
	body map[string]interface{}
	code int
}

func (r *Response) Error() string {
	return r.err
}

// AddBody adds a field in the response's body
func (r *Response) AddBody(key, value string) {
	r.body[key] = value
}

// Body returns the response body
func (r *Response) Body() map[string]interface{} {
	return r.body
}

// Code returns the response HTTP status code
func (r *Response) Code() int {
	return r.code
}

// New returns the most simple Response:
// a HTTP status code and an empty body
func New(code int) *Response {
	res := &Response{
		code: code,
		body: make(map[string]interface{}),
	}
	return res
}

// NewWithBody returns a simple New Response and
// adds a body
func NewWithBody(code int, key, value string) *Response {
	res := New(code)
	res.AddBody(key, value)
	return res
}

// NewWithError returns a the most simple type of error-like
// Response: a HTTP status code, a response and an error message
func NewWithError(code int, err, msg string) *Response {
	return &Response{
		code: code,
		body: map[string]interface{}{
			"message": msg,
		},
		err: err,
	}
}

// NewWithMessage returns a the most simple type of info-like
// Response: a HTTP status code and a response message
func NewWithMessage(code int, msg string) *Response {
	return &Response{
		code: code,
		body: map[string]interface{}{
			"message": msg,
		},
	}
}

// Redirect = 302
func Redirect(url string) *Response {
	return NewWithBody(302, "url", url)
}

// BadRequest = 400 response code
func BadRequest(msg string) *Response {
	log.Error(msg)
	return NewWithError(400, msg, msg)
}

// ServiceUnavailable = 503 response code
func ServiceUnavailable(msg, err string) *Response {
	log.Error(err)
	return NewWithError(503, err, msg)
}

// Ok = 200 response code
func Ok(c *gin.Context, res map[string]interface{}) {
	data, _ := json.Marshal(res)
	c.Render(http.StatusOK, render.Data{
		ContentType: "application/json",
		Data:        data,
	})
}

// Write handles the output response of a JSON endpoint
func Write(c *gin.Context, res *Response) {
	c.JSON(res.Code(), res.Body())
}
