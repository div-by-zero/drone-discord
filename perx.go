package main

import (
	"fmt"
	"strings"
)

// PerxTemplate is custom template for Drone CI.
func (p *Plugin) PerxTemplate() EmbedObject {
	description := ""
	fields := []EmbedFieldObject{
		{
			Name:   "Build",
			Value:  fmt.Sprintf("[#%d](%s)", p.Build.Number, p.Build.Link),
			Inline: true,
		},
	}
	if p.Build.PreviewURL != "" {
		fields = append(fields, EmbedFieldObject{
			Name:   "Preview",
			Value:  fmt.Sprintf("[Open](%s)", p.Build.PreviewURL),
			Inline: true,
		})
	}
	switch p.Build.Event {
	case "push":
		description = fmt.Sprintf("pushed to %s.", p.Build.Branch)
		if strings.Contains(p.Build.PrevRefSpec, "pull") {
			prNumber := strings.TrimPrefix(strings.TrimSuffix(p.Build.PrevRefSpec, "/merge"), "refs/pull/")
			description = fmt.Sprintf("merged PR [#%s](%s) to %s.",
				prNumber, p.Build.PrevCommitLink, p.Build.Branch)
		}
	case "pull_request":
		branch := ""
		if p.Build.RefSpec != "" {
			branch = p.Build.RefSpec
		} else {
			branch = p.Build.Branch
		}
		description = fmt.Sprintf("updated pull request %s", branch)
	case "tag":
		description = fmt.Sprintf("pushed tag %s", p.Build.Branch)
	}

	return EmbedObject{
		Title:       p.Build.Message,
		Description: description,
		URL:         p.Build.CommitLink,
		Color:       p.Color(),
		Author: EmbedAuthorObject{
			Name:    p.Build.Author,
			IconURL: p.Build.Avatar,
		},
		Fields: fields,
		Footer: EmbedFooterObject{
			Text:    DroneDesc,
			IconURL: DroneIconURL,
		},
	}
}
