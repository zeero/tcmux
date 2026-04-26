package agent

import "testing"

func TestGeminiAgent_Match(t *testing.T) {
	agent := &GeminiAgent{}

	// Mock GetCommandLine
	originalGetCommandLine := GetCommandLine
	defer func() { GetCommandLine = originalGetCommandLine }()

	tests := []struct {
		name           string
		title          string
		currentCommand string
		commandLine    string
		want           bool
	}{
		{name: "Gemini process", title: "zsh", currentCommand: "gemini", want: true},
		{name: "Gemini-cli process", title: "zsh", currentCommand: "gemini-cli", want: true},
		{name: "Node with Gemini title", title: "Gemini CLI", currentCommand: "node", want: true},
		{name: "Node with lowercase gemini title", title: "gemini-cli", currentCommand: "node", want: true},
		{name: "Node with no gemini in title", title: "zsh", currentCommand: "node", want: false},
		{name: "Random process", title: "Gemini CLI", currentCommand: "zsh", want: false},
		{name: "Node with Gemini in command line", title: "zsh", currentCommand: "node", commandLine: "node /usr/local/bin/gemini", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetCommandLine = func(pid string) string {
				return tt.commandLine
			}
			vars := map[string]string{
				"pane_title":           tt.title,
				"pane_current_command": tt.currentCommand,
				"pane_pid":             "1234",
			}
			got := agent.Match(vars)
			if got != tt.want {
				t.Errorf("GeminiAgent.Match(%q, %q) = %v, want %v (commandLine: %q)", tt.title, tt.currentCommand, got, tt.want, tt.commandLine)
			}
		})
	}
}

func TestGeminiAgent_ExtractSummary(t *testing.T) {
	agent := &GeminiAgent{}
	tests := []struct {
		name  string
		title string
		want  string
	}{
		{"Normal title", "Fix bug", "Fix bug"},
		{"With spaces", "  Fix bug  ", "Fix bug"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := agent.ExtractSummary(tt.title)
			if got != tt.want {
				t.Errorf("GeminiAgent.ExtractSummary(%q) = %q, want %q", tt.title, got, tt.want)
			}
		})
	}
}
