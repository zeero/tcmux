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

// MayBeTitle checks if the pane title may indicate a Gemini CLI instance.
func (a *GeminiAgent) MayBeTitle(title string) bool {
	// If the process is "gemini-cli" or "gemini", we trust it.
	// If the process is "node", we need some hint from the title.
	// Common practice for these CLI tools is to set the title.
	// If it's just "node" and title is generic like "zsh" or empty, it's likely not it.
	// Gemini CLI usually has "Gemini" in the title.
	return strings.Contains(strings.ToLower(title), "gemini")
}

// MayBeProcess checks if the current command may be a Gemini CLI process.
func (a *GeminiAgent) MayBeProcess(currentCommand string) bool {
	// User said: node gemini_path
	// tmux's pane_current_command usually only shows the executable name (e.g., "node")
	return currentCommand == "node" || currentCommand == "gemini-cli" || currentCommand == "gemini"
}

// ExtractSummary extracts the task summary from the pane title.
func (a *GeminiAgent) ExtractSummary(title string) string {
	return strings.TrimSpace(title)
}

// ParseStatus parses the pane content and determines the Gemini CLI status.
func (a *GeminiAgent) ParseStatus(content string) Status {
	return parseGeminiStatus(content)
}
