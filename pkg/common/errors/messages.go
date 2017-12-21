package errors

import "errors"

var (
	NotLoggedMessage   = errors.New("You are currntly not logged in to system , to get proper access create a new user or login with an exiting user.")
	LoginErrorMessage  = errors.New("Incorrect login or password")
	LogoutErrorMessage = errors.New("Some problems with logout")
	UnknownMessage     = errors.New("Ooops, error occurred: Please provide bug to github: https://github.com/pososi/moi/yaiki")
)
