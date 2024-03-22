package helper

import (
	"agit-test/model/web"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func ReadFromJSON(c *gin.Context, data interface{}) {
	err := c.BindJSON(data)
	if err != nil {
		PanicIfError(err)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func WriteResponseJSON(c *gin.Context, result interface{}, err error) {
	var code int = 200
	var status string = "OK"
	var message string = "success"
	if err != nil {
		code = 500
		status = "ERROR"
		message = err.Error()
	} else {

	}
	webResponse := web.WebResponse{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    result,
	}

	c.JSON(code, webResponse)
}

// func WriteToResponseBody(c *gin.Context, result interface{}) {
// 	// writer.Header().Add("Content-Type", "application/json")

// 	encoder := json.NewEncoder(writer)
// 	err := encoder.Encode(response)
// 	PanicIfError(err)
// }
