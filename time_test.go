package timed_test

import (
	"testing"
	"time"

	"github.com/ItsMeViipeR/timed"
)

func TestSince(t *testing.T) {
	testsSince := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:    "Date time.Time d'il y a 10 jours",
			input:   time.Now().Add(-240 * time.Hour),
			want:    "10 days ago",
			wantErr: false,
		},
		{
			name:    "String valide d'il y a 2 heures",
			input:   time.Now().Add(-2 * time.Hour),
			want:    "2 hours ago",
			wantErr: false,
		},
		{
			name:    "String avec mauvais format",
			input:   "2026/01/01",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Type non supporté (int)",
			input:   12345,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range testsSince {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timed.Since(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Since() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("Since() = %q, want %q", got, tt.want)
			}
		})
	}

	testsUntil := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:    "In 2 days",
			input:   time.Now().Add(2 * 24 * time.Hour),
			want:    "In 2 days",
			wantErr: false,
		},
		{
			name:    "In 1 days and 3 hours",
			input:   time.Now().Add(24 * time.Hour).Add(3 * time.Hour),
			want:    "In 1 days",
			wantErr: false,
		},
	}

	for _, tt := range testsUntil {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timed.Until(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Until() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("Until() = %q, want %q", got, tt.want)
			}
		})
	}
}
