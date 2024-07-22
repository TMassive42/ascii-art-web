package asciiart

import (
	"strings"
	"testing"
)

func TestGenerateASCIIArt(t *testing.T) {
	type args struct {
		input  string
		banner string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// Test case 1: Valid input and banner
		{
			name: "Valid input and banner",
			args: args{
				input:  "Hello",
				banner: "banners/standard",
			},
			// Replace this with expected ASCII art output
			want: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
								
								`,
			wantErr: false,
		},

		// Test case 2: Non-existent banner
		{
			name: "Non-existent banner",
			args: args{
				input:  "World",
				banner: "notreal",
			},
			want:    "",
			wantErr: true, // Expect error for missing banner file
		},

		// Test case 3: Invalid banner format (wrong line count)
		{
			name: "Invalid banner format (wrong line count)",
			args: args{
				input:  "Test",
				banner: "broken", // Create a test file with wrong line count
			},
			want:    "",
			wantErr: true, // Expect error for invalid banner format
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateASCIIArt(tt.args.input, tt.args.banner)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateASCIIArt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotTrimmed := strings.TrimSpace(got)
			if gotTrimmed != strings.TrimSpace(tt.want) {
				t.Errorf("GenerateASCIIArt() = %v, want %v", got, tt.want)
			}
		})
	}
}
