package internal

import "fmt"

func GroupAndFormat(events []GitHubEvent) []string {
	var result []string

	for i := 0; i < len(events); {
		e := events[i]

		// Group PushEvents by repo
		if e.Type == "PushEvent" {
			count := 1
			j := i + 1

			for j < len(events) &&
				events[j].Type == "PushEvent" &&
				events[j].Repo.Name == e.Repo.Name {
				count++
				j++
			}

			if count == 1 {
				result = append(
					result,
					fmt.Sprintf("Pushed to %s", e.Repo.Name),
				)
			} else {
				result = append(
					result,
					fmt.Sprintf(
						"Pushed %d times to %s",
						count,
						e.Repo.Name,
					),
				)
			}

			i = j
			continue
		}

		// Non-push events
		line := FormatEvent(e)
		if line != "" {
			result = append(result, line)
		}

		i++
	}

	return result
}
