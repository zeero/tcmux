package agent

import "testing"

func TestParseGeminiStatus(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		wantState string
	}{
		{
			"Idle state",
			"Welcome to Gemini CLI\n> ",
			StateIdle,
		},
		{
			"Idle state with history",
			"Previous output\n\n> ",
			StateIdle,
		},
		{
			"Running state",
			"Checking files...\nThinking…",
			StateRunning,
		},
		{
			"Waiting state - Allow once",
			"Do you want to run this command?\n1. Allow once\n2. Allow for this session\n3. No, suggest changes (esc)",
			StateWaiting,
		},
		{
			"Waiting state - Allow for this session",
			"2. Allow for this session",
			StateWaiting,
		},
		{
			"Waiting state - No, suggest changes",
			"3. No, suggest changes (esc)",
			StateWaiting,
		},
		{
			"Unknown state",
			"Some other output",
			StateUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseGeminiStatus(tt.content)
			if got.State != tt.wantState {
				t.Errorf("parseGeminiStatus() State = %v, want %v", got.State, tt.wantState)
			}
		})
	}
}
