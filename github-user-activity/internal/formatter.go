package internal

import "fmt"

func FormatEvent(e GitHubEvent) string {
	switch e.Type {

	case "PushEvent":
		if len(e.Payload.Commits) == 0 {
			return fmt.Sprintf(
				"Pushed to %s",
				e.Repo.Name,
			)
		}
		return fmt.Sprintf(
			"Pushed %d commits to %s",
			len(e.Payload.Commits),
			e.Repo.Name,
		)

	case "IssuesEvent":
		if e.Payload.Action == "opened" {
			return fmt.Sprintf(
				"Opened a new issue in %s",
				e.Repo.Name,
			)
		}

	case "WatchEvent":
		return fmt.Sprintf(
			"Starred %s",
			e.Repo.Name,
		)

	case "PullRequestEvent":
		if e.Payload.Action == "opened" {
			return fmt.Sprintf(
				"Opened a pull request in %s",
				e.Repo.Name,
			)
		}
	}

	return ""
}
