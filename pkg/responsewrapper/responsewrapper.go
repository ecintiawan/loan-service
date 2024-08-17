package responsewrapper

import (
	"net/http"

	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
	"github.com/labstack/echo/v4"
)

func Response(message, status, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"status":  status,
		"data":    data,
	}
}

func response(message string, httpStatus int, data interface{}) map[string]interface{} {
	return Response(message, http.StatusText(httpStatus), data)
}

func ErrorHandler(err error, c echo.Context) error {
	errx, ok := err.(*errorwrapper.Error)
	if ok {
		switch errx.Code {
		case errorwrapper.CodeInvalid:
			err = badRequest(c, err, nil)
		case errorwrapper.CodeNotFound:
			err = notFound(c, err, nil)
		default:
			err = internalServerError(c, err, nil)
		}
	}

	return err
}

func OK(c echo.Context, message string, data interface{}) error {
	body := response(message, http.StatusOK, data)
	return c.JSON(http.StatusOK, body)
}

func Created(c echo.Context, message string, data interface{}) error {
	body := response(message, http.StatusCreated, data)
	return c.JSON(http.StatusCreated, body)
}

func badRequest(c echo.Context, err error, data interface{}) error {
	body := response(err.Error(), http.StatusBadRequest, data)
	return c.JSON(http.StatusBadRequest, body)
}

// notFound is used to denote that the requested resource does not exist
func notFound(c echo.Context, err error, data interface{}) error {
	body := response(err.Error(), http.StatusNotFound, data)
	return c.JSON(http.StatusNotFound, body)
}

func internalServerError(c echo.Context, err error, data interface{}) error {
	body := response(err.Error(), http.StatusInternalServerError, data)
	return c.JSON(http.StatusInternalServerError, body)
}
