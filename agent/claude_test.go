package agent

import "testing"

func TestClaudeAgent_Match(t *testing.T) {
	agent := &ClaudeAgent{}
	tests := []struct {
		name           string
		title          string
		currentCommand string
		want           bool
	}{
		{"Idle with node", "✳ summary", "node", true},
		{"Spinner with node", "⠋ summary", "node", true},
		{"Idle with claude", "✳ summary", "claude", true},
		{"Version with idle title", "✳ summary", "2.1.34", true},
		{"Generic title with node", "zsh", "node", false},
		{"Generic title with claude", "zsh", "claude", false},
		{"Correct title but wrong process", "✳ summary", "zsh", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vars := map[string]string{
				"pane_title":           tt.title,
				"pane_current_command": tt.currentCommand,
			}
			got := agent.Match(vars)
			if got != tt.want {
				t.Errorf("ClaudeAgent.Match(%q, %q) = %v, want %v", tt.title, tt.currentCommand, got, tt.want)
			}
		})
	}
}

func TestClaudeAgent_ExtractSummary(t *testing.T) {
	agent := &ClaudeAgent{}
	tests := []struct {
		name  string
		title string
		want  string
	}{
		{"Normal title", "✳ Task summary", "Task summary"},
		{"Japanese title", "✳ 日本語タスク", "日本語タスク"},
		{"With extra spaces", "✳  Multiple  spaces", "Multiple  spaces"},
		{"Only prefix", "✳", ""},
		{"Only prefix with space", "✳ ", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := agent.ExtractSummary(tt.title)
			if got != tt.want {
				t.Errorf("ClaudeAgent.ExtractSummary(%q) = %q, want %q", tt.title, got, tt.want)
			}
		})
	}
}
