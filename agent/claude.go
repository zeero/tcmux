package agent

import (
	"regexp"
	"strings"
	"unicode"
)

// Claude Code title prefix
const claudePrefixIdle = "✳" // Idle state

// ClaudeAgent detects and parses Claude Code instances.
type ClaudeAgent struct{}

func (a *ClaudeAgent) Type() Type {
	return TypeClaude
}

func (a *ClaudeAgent) Icon() string {
	return "✻"
}

// Match checks if the pane title and current command indicate a Claude Code instance.
func (a *ClaudeAgent) Match(paneVars map[string]string) bool {
	title := paneVars["pane_title"]
	currentCommand := paneVars["pane_current_command"]
	if currentCommand != "node" && currentCommand != "claude" && !claudeVersionPattern.MatchString(currentCommand) {
		return false
	}
	if strings.HasPrefix(title, claudePrefixIdle) {
		return true
	}
	// Check for Braille pattern dots (U+2800-U+28FF) used as spinner
	if len(title) > 0 {
		r := []rune(title)
		if isBraillePattern(r[0]) {
			return true
		}
	}
	return false
}

// isBraillePattern checks if a rune is a Braille pattern character (U+2800-U+28FF).
func isBraillePattern(r rune) bool {
	return unicode.In(r, unicode.Braille)
}

// claudeVersionPattern matches semver-like version strings (e.g., "2.1.34").
// Native Install of Claude Code reports its version as pane_current_command.
var claudeVersionPattern = regexp.MustCompile(`^\d+\.\d+`)

// ExtractSummary extracts the task summary from the pane title.
func (a *ClaudeAgent) ExtractSummary(title string) string {
	// Remove the "✳ " prefix
	if strings.HasPrefix(title, claudePrefixIdle) {
		summary := strings.TrimPrefix(title, claudePrefixIdle)
		return strings.TrimSpace(summary)
	}
	// Remove Braille pattern prefix (spinner)
	if len(title) > 0 {
		r := []rune(title)
		if isBraillePattern(r[0]) {
			summary := string(r[1:])
			return strings.TrimSpace(summary)
		}
	}
	return strings.TrimSpace(title)
}

// ParseStatus parses the pane content and determines the Claude Code status.
func (a *ClaudeAgent) ParseStatus(content string) Status {
	return parseClaudeStatus(content)
}
