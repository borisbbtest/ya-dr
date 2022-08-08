package postgres

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Error string

func (e Error) Error() string { return string(e) }

const (
	errorTooManyParameters               = Error("Too many parameters.")
	errorCannotFetchData                 = Error("Cannot fetch data.")
	errorPostgresPing                    = Error("PostgreSQL ping failed: got wrong data.")
	errorCannotParseData                 = Error("Cannot parse data.")
	errorUnsupportedQuery                = Error("Unsupported query.")
	errorEmptyResult                     = Error("Empty result.")
	errorCannotConvertPostgresVersionInt = Error("Cannot convert Postgres version to integer.")
	errorFourthParamEmpty                = Error("The key requires database name as fourth parameter")
	errorFourthParamLen                  = Error("Expected database name as fourth parameter for the key, got empty string")
	errorUnknownSession                  = Error("Unknown session")
)

// formatError formats a given error text. It capitalizes the first letter and adds a dot to the end.
// TBD: move to the agent's core
func formatError(errText string) string {
	if errText[len(errText)-1:] != "." {
		errText = errText + "."
	}
	return cases.Title(language.Und, cases.NoLower).String(errText)
}

func sanitizeError(errText, connString string) (err error) {
	// default error string shows user credentials used to establish connection
	// this can be potentialy dangerous
	// create error message string without sensitive info
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
