package agent

import (
	"strings"
)

// GeminiAgent detects and parses Gemini CLI instances.
type GeminiAgent struct{}

func (a *GeminiAgent) Type() Type {
	return TypeGemini
}

func (a *GeminiAgent) Icon() string {
	return "♊"
}

// Match checks if the pane title and current command indicate a Gemini CLI instance.
func (a *GeminiAgent) Match(paneVars map[string]string) bool {
	title := paneVars["pane_title"]
	currentCommand := paneVars["pane_current_command"]
	// If the process is "gemini-cli" or "gemini", we trust it regardless of title.
	if currentCommand == "gemini-cli" || currentCommand == "gemini" {
		return true
	}
	// If the process is "node", we check title OR the full command line.
	if currentCommand == "node" {
		if strings.Contains(strings.ToLower(title), "gemini") {
			return true
		}
		// Fallback: check command line
		if strings.Contains(strings.ToLower(GetCommandLine(paneVars["pane_pid"])), "gemini") {
			return true
		}
	}
	return false
}

// ExtractSummary extracts the task summary from the pane title.
func (a *GeminiAgent) ExtractSummary(title string) string {
	return strings.TrimSpace(title)
}

// ParseStatus parses the pane content and determines the Gemini CLI status.
func (a *GeminiAgent) ParseStatus(content string) Status {
	return parseGeminiStatus(content)
}
