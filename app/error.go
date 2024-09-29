package app

import "errors"

const (
	ErrorEmptyGitLabProjectIdPage = "empty page of project id"
)

type errorEmptyGitLabProjectIdPage struct {
	error
}

func IsErrorEmptyProjectIdPage(err error) bool {
	return errors.As(err, &errorEmptyGitLabProjectIdPage{})
}
