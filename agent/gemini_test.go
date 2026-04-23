package agent

import "testing"

func TestGeminiAgent_MayBeTitle(t *testing.T) {
	agent := &GeminiAgent{}
	tests := []struct {
		name  string
		title string
		want  bool
	}{
		{"Gemini title", "Gemini CLI", true},
		{"Lowercase gemini title", "gemini-cli", true},
		{"Empty title", "", false},
		{"zsh", "zsh", false},
		{"Random title", "some random title", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := agent.MayBeTitle(tt.title)
			if got != tt.want {
				t.Errorf("GeminiAgent.MayBeTitle(%q) = %v, want %v", tt.title, got, tt.want)
			}
		})
	}
}

func TestGeminiAgent_MayBeProcess(t *testing.T) {
	agent := &GeminiAgent{}
	tests := []struct {
		name           string
		currentCommand string
		want           bool
	}{
		{"node", "node", true},
		{"gemini-cli", "gemini-cli", true},
		{"gemini", "gemini", true},
		{"zsh", "zsh", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := agent.MayBeProcess(tt.currentCommand)
			if got != tt.want {
				t.Errorf("GeminiAgent.MayBeProcess(%q) = %v, want %v", tt.currentCommand, got, tt.want)
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
