package agent

import (
	"os/exec"
	"strings"
)

// Type identifies the type of coding agent.
type Type string

const (
	TypeClaude  Type = "claude"
	TypeCopilot Type = "copilot"
	TypeCodex   Type = "codex"
	TypeGemini  Type = "gemini"
)

// Status represents the status of a coding agent instance.
type Status struct {
	State       string // Idle, Running, Waiting, Unknown
	Mode        string // plan mode, accept edits, or empty
	Description string // Additional description (e.g., time elapsed)
}

// Status state constants
const (
	StateIdle    = "Idle"
	StateRunning = "Running"
	StateWaiting = "Waiting" // Agent is waiting for user input/selection
	StateUnknown = "Unknown"
)

// Mode constants
const (
	ModePlan        = "plan mode"
	ModeAcceptEdits = "accept edits"
)

// Detector defines the interface for detecting and parsing coding agents.
type Detector interface {
	Type() Type
	Icon() string
	Match(paneVars map[string]string) bool
	ExtractSummary(title string) string
	ParseStatus(content string) Status
}

// All registered detectors
var detectors = []Detector{
	&ClaudeAgent{},
	&CopilotAgent{},
	&CodexAgent{},
	&GeminiAgent{},
}

// Detect checks if a pane might be running a coding agent.
// Returns the detected agent or nil if no agent is detected.
func Detect(paneVars map[string]string) Detector {
	for _, d := range detectors {
		if d.Match(paneVars) {
			return d
		}
	}
	return nil
}

// GetCommandLine returns the full command line of a process.
var GetCommandLine = func(pid string) string {
	if pid == "" {
		return ""
	}
	// ps -p <pid> -o command=
	cmd := exec.Command("ps", "-p", pid, "-o", "command=")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

