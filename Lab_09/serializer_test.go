package main

import (
	"strings"
	"testing"
)

// ==========================================
// Табличні юніт-тести для ToYAML
// ==========================================
func TestToYAML(t *testing.T) {

	tests := []struct {
		name      string
		input     any
		wantError bool
		wantSub   string
	}{
		{
			name: "Успішна серіалізація структури Server",
			input: Server{
				Host:       "localhost",
				Port:       8080,
				Debug:      true,
				AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
			},
			wantError: false,
			wantSub:   "host: \"localhost\"\nport: 8080\ndebug: true\nallowed_ips: \n  - \"192.168.1.1\"\n  - \"10.0.0.1\"",
		},
		{
			name:      "Помилка при передачі не структури (рядок)",
			input:     "just a string",
			wantError: true,
		},
		{
			name:      "Помилка при передачі числа",
			input:     42,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToYAML(tt.input)

			if (err != nil) != tt.wantError {
				t.Errorf("ToYAML() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError && !strings.Contains(got, tt.wantSub) {
				t.Errorf("ToYAML() результат:\n%v\n\nОчікувалося наявність:\n%v", got, tt.wantSub)
			}
		})
	}
}

// ==========================================
// Бенчмарки (Вимірювання продуктивності)
// ==========================================

var benchData = Server{
	Host:       "localhost",
	Port:       8080,
	Debug:      true,
	AllowedIPs: []string{"192.168.1.1", "10.0.0.1", "10.0.0.2", "127.0.0.1"},
}

func BenchmarkToYAML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ToYAML(benchData)
	}
}

func BenchmarkToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ToJSON(benchData)
	}
}
