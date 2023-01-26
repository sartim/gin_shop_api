package core

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func LogEvent(logLevel string) *log.Logger {
	// Info writes logs in the color blue with "INFO: " as prefix
	var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)

	// Warning writes logs in the color yellow with "WARNING: " as prefix
	var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)

	// Error writes logs in the color red with "ERROR: " as prefix
	var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)

	// Debug writes logs in the color cyan with "DEBUG: " as prefix
	var Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)

	switch logLevel {
	case "INFO":
		return Info
	case "WARNING":
		return Warning
	case "ERROR":
		return Error
	case "DEBUG":
		return Debug
	}

	return nil
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Log(logLevel string) *log.Logger {
	var logEvent = LogEvent(logLevel)

	return logEvent
}

func MsgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	}
	return ""
}

func ValidateSchema(ctx *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		body := make(gin.H)
		for _, fe := range ve {
			body[strings.ToLower(fe.Field())] = MsgForTag(fe.Tag())
		}
		ctx.JSON(http.StatusBadRequest, body)
	}
}
