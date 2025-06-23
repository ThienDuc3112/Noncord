package entities

import "regexp"

var emailReg = regexp.MustCompile(`^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$`)

func IsValidEmail(email string) bool {
	return emailReg.MatchString(email)
}

var usernameReg = regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9_-]{1,30}[a-zA-Z0-9])$`)

func IsValidUsername(username string) bool {
	return usernameReg.MatchString(username)
}

var urlReg = regexp.MustCompile(`(?:http[s]?:\/\/.)?(?:www\.)?[-a-zA-Z0-9@%._\+~#=]{2,256}\.[a-z]{2,6}\b(?:[-a-zA-Z0-9@:%_\+.~#?&\/\/=]*)`)

func IsValidUrl(url string) bool {
	return urlReg.MatchString(url)
}
