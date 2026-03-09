package agent

import "strings"

// CodexAgent detects and parses Codex CLI instances.
type CodexAgent struct{}

func (a *CodexAgent) Type() Type {
	return TypeCodex
}

func (a *CodexAgent) Icon() string {
	return "❂"
}

// MayBeTitle checks if the pane title may indicate a Codex CLI instance.
// Title check is permissive - process name is the primary signal.
func (a *CodexAgent) MayBeTitle(title string) bool {
	return true
}

// MayBeProcess checks if the current command may be a Codex CLI process.
func (a *CodexAgent) MayBeProcess(currentCommand string) bool {
	currentCommand = strings.ToLower(strings.TrimSpace(currentCommand))
	return currentCommand == "codex" || strings.HasPrefix(currentCommand, "codex-")
}

// ExtractSummary extracts the task summary from the pane title.
func (a *CodexAgent) ExtractSummary(title string) string {
	title = strings.TrimSpace(title)
	if title == "" {
		return ""
	}
	if title == "Codex" {
		return ""
	}
	return title
}

// ParseStatus parses the pane content and determines the Codex CLI status.
func (a *CodexAgent) ParseStatus(content string) Status {
	return parseCodexStatus(content)
}
