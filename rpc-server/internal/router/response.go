package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/islu/ethereum_private_chain/rpc_server/internal/domain/common"
)

// @description	Error message response for 4xx and 5xx errors
type ErrorMessage struct {
	Name       string                 `json:"name" example:"PARAMETER_INVALID"`       // Error name (key)
	Code       int                    `json:"code" example:"400"`                     // HTTP status code
	Message    string                 `json:"message,omitempty" swaggerignore:"true"` // Client message
	RemoteCode int                    `json:"remoteCode,omitempty" swaggerignore:"true"`
	Detail     map[string]interface{} `json:"detail,omitempty" swaggerignore:"true"`
} // @name ErrorMessageResponse

// @description	Success message response for 200
type SuccessMessage struct {
	Name string `json:"name" example:"SUCCESS"`
	Code int    `json:"code" example:"200"`
} // @name SuccessMessageResponse

func respondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}

func respondWithoutBody(c *gin.Context, code int) {
	c.Status(code)
}

func respondWithSuccess(c *gin.Context) {
	respondWithJSON(c, http.StatusOK, SuccessMessage{Name: "SUCCESS", Code: http.StatusOK})
}

func respondWithError(c *gin.Context, err error) {
	errMessage := parseError(err)

	// TODO: Add logger
	// ctx := c.Request.Context()
	// zerolog.Ctx(ctx).Error().Err(err).Str("component", "handler").Msg(errMessage.Message)

	_ = c.Error(err)
	c.AbortWithStatusJSON(errMessage.Code, errMessage)
}

func parseError(err error) ErrorMessage {
	var domainError common.DomainError
	// We don't check if errors.As is valid or not
	// because an empty common.DomainError would return default error data.
	_ = errors.As(err, &domainError)

	return ErrorMessage{
		Name:       domainError.Name(),
		Code:       domainError.HTTPStatus(),
		Message:    domainError.ClientMsg(),
		RemoteCode: domainError.RemoteHTTPStatus(),
		Detail:     domainError.Detail(),
	}
}
