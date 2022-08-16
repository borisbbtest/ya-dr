package postgres

import (
	"fmt"
	"strings"
)

type Error string

func (e Error) Error() string { return string(e) }

func sanitizeError(errText, connString string) (err error) {
	var (
		originSchema, sanitizedErrorString string
		startingIndexOfError               int
	)
	if strings.Contains(connString, "host=") {
		originSchema = "unix://"
	} else {
		originSchema = "tcp://"
	}
	// want to get only uri from conn string for the error message
	splitted := strings.Split(connString, "@")
	if len(splitted) > 1 {
		// create error message without username and password and with the right schemaname
		startingIndexOfError = strings.Index(errText, connString) + len(connString) + 2
		sanitizedErrorString = fmt.Sprintf("%s%s  %s ", originSchema, splitted[1], errText[startingIndexOfError:])
	} else {
		// should never happen
		sanitizedErrorString = errText[len(connString):]
	}

	return fmt.Errorf("invalid connection string: %s", sanitizedErrorString)
}
