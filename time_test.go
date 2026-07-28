package timed_test

import (
	"testing"
	"time"

	"github.com/ItsMeViipeR/timed"
)

func TestSince(t *testing.T) {
	// 1. Définition des cas de test
	tests := []struct {
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
			wantErr: true, // On s'attend à une erreur
		},
		{
			name:    "Type non supporté (int)",
			input:   12345,
			want:    "",
			wantErr: true, // On s'attend à une erreur
		},
	}

	// 2. Exécution de chaque cas
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timed.Since(tt.input)

			// Vérification de la présence d'une erreur
			if (err != nil) != tt.wantErr {
				t.Fatalf("Since() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Vérification du résultat obtenu
			if got != tt.want {
				t.Errorf("Since() = %q, want %q", got, tt.want)
			}
		})
	}
}
