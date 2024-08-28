package format

import (
	"fmt"
	"github-checker/internal/request"
)

func GithubEventMsg(e request.GithubUserEvent) string {
	var msg string
	switch e.Type {
	case "PushEvent":
		msg = fmt.Sprintf("Pushed %d commits to %s", len(e.Payload.Commits), e.Repo.Name)
	case "WatchEvent":
		msg = fmt.Sprintf("Starred %s", e.Repo.Name)
	case "IssuesEvent":
		msg = makeIssuesEventMsg(e)
	case "CreateEvent":
		msg = makeCreateEventMsg(e)
	default:
		msg = "unknown"
	}

	return msg
}

func makeIssuesEventMsg(e request.GithubUserEvent) string {
	if e.Payload.Action == "opened" {
		return fmt.Sprintf("Opened a new issue in %s", e.Repo.Name)
	}
	return fmt.Sprintf("Closed an issue in %s", e.Repo.Name)
}

func makeCreateEventMsg(e request.GithubUserEvent) string {
	if e.Payload.RefType == "branch" {
		return fmt.Sprintf("Created branch %s in %s", e.Payload.Ref, e.Repo.Name)
	}
	return fmt.Sprintf("Created %s", e.Repo.Name)
}
