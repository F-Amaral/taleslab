package apierror

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apilogger"
)

type ApiError interface {
	ErrorStatusCode() int
	Error() string
}

func New(httpStatusCode int, message string) *apiError {
	return &apiError{
		HttpCode: httpStatusCode,
		Message:  message,
	}
}

type apiError struct {
	Message  string `json:"message"`
	HttpCode int    `json:"http_code"`
}

func (self apiError) ErrorStatusCode() int {
	return self.HttpCode
}

func (self apiError) Error() string {
	return self.Message
}

func Log(ctx context.Context, err ApiError) {
	apilogger.Error(ctx, fmt.Sprintf("status_code: %d, message: %s", err.ErrorStatusCode(), err.Error()))
}
