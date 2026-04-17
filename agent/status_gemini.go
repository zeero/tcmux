package agent

import (
	"regexp"
	"strings"
)

var (
	// Gemini idle pattern: prompt line
	geminiIdlePattern = regexp.MustCompile(`(?m)^\s*> `)

	// Gemini running pattern
	geminiRunningPattern = regexp.MustCompile(`(?m)Thinking…`)

	// Gemini waiting patterns
	geminiWaitingPatterns = []string{
		"Allow once",
		"Allow for this session",
		"No, suggest changes",
	}
)

func parseGeminiStatus(content string) Status {
	lines := strings.Split(content, "\n")
	lastLines := lastNonEmptyLines(lines, 10)
	combined := strings.Join(lastLines, "\n")

	status := Status{
		State: StateUnknown,
	}

	if geminiRunningPattern.MatchString(combined) {
		status.State = StateRunning
		return status
	}

	for _, pattern := range geminiWaitingPatterns {
		if strings.Contains(combined, pattern) {
			status.State = StateWaiting
			return status
		}
	}

	if geminiIdlePattern.MatchString(combined) {
		status.State = StateIdle
		return status
	}

	return status
}
