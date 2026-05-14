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

// Match checks if the pane title and current command indicate a Codex CLI instance.
func (a *CodexAgent) Match(paneVars map[string]string) bool {
	currentCommand := strings.ToLower(strings.TrimSpace(paneVars["pane_current_command"]))
	if currentCommand == "codex" || strings.HasPrefix(currentCommand, "codex-") {
		return true
	}

	if currentCommand == "node" {
		title := strings.TrimSpace(paneVars["pane_title"])
		if title == "Codex" {
			return true
		}
		pid := paneVars["pane_pid"]
		if pid != "" {
			cmdLine := GetCommandLine(pid)
			if strings.Contains(cmdLine, "codex") {
				return true
			}
		}
	}

	return false
}

// ExtractSummary extracts the task summary from the pane title.
func (a *CodexAgent) ExtractSummary(title string) string {
	// Remove common emoji prefixes.
	title = strings.TrimSpace(title)
	if title == "" {
		return ""
	}
	if len(title) > 0 {
		r := []rune(title)
		if len(r) > 0 && r[0] > 0x1F000 {
			title = strings.TrimSpace(string(r[1:]))
		}
	}

	// Treat default Codex title / host-like local title as empty summary.
	if title == "Codex" || strings.HasSuffix(strings.ToLower(title), ".local") {
		return ""
	}

	return title
}

// ParseStatus parses the pane content and determines the Codex CLI status.
func (a *CodexAgent) ParseStatus(content string) Status {
	return parseCodexStatus(content)
}
