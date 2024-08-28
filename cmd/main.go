package cmd

import (
	"errors"
	"fmt"
	"github-checker/internal/format"
	"github-checker/internal/request"
	"os"
)

func Execute() error {
	args := os.Args[1:]
	if len(args) == 0 {
		return errors.New("missing github username")
	}

	username := args[0]
	events, err := request.FetchGithubUserEvents(username)
	if err != nil {
		return err
	}

	var messages []string
	for _, e := range events {
		messages = append(messages, format.GithubEventMsg(e))
	}

	for _, m := range messages {
		fmt.Printf("- %s\n", m)
	}

	return nil
}
