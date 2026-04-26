package agent

import "strings"

// CopilotAgent detects and parses GitHub Copilot CLI instances.
type CopilotAgent struct{}

func (a *CopilotAgent) Type() Type {
	return TypeCopilot
}

func (a *CopilotAgent) Icon() string {
	return "⬢"
}

// Match checks if the pane title and current command indicate a Copilot CLI instance.
func (a *CopilotAgent) Match(paneVars map[string]string) bool {
	return paneVars["pane_current_command"] == "copilot"
}

// ExtractSummary extracts the task summary from the pane title.
func (a *CopilotAgent) ExtractSummary(title string) string {
	// Remove common emoji prefixes
	title = strings.TrimSpace(title)
	if len(title) > 0 {
		r := []rune(title)
		// Skip leading emoji (if any)
		if len(r) > 0 && r[0] > 0x1F000 {
			title = strings.TrimSpace(string(r[1:]))
		}
	}
	// "GitHub Copilot" is the default title, return empty
	if title == "GitHub Copilot" {
		return ""
	}
	return title
}

// ParseStatus parses the pane content and determines the Copilot CLI status.
func (a *CopilotAgent) ParseStatus(content string) Status {
	return parseCopilotStatus(content)
}
