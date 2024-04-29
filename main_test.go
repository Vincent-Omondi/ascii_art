package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Empty String",
			input:  "",
			output: "",
		},
		{
			name:   "Newline",
			input:  "\\n",
			output: "\n",
		},
		{
			name:  "Simple Text",
			input: "Hello",
			output: `      
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
																			
`,
		},
		{
			name:   "Unsupported Character",
			input:  "€uro",
			output: "Error: Character '€' is not accepted\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Set arguments and run main
			os.Args = []string{"go run .", tt.input}
			main()

			// Restore stdout and capture output
			w.Close()
			os.Stdout = oldStdout
			out, _ := io.ReadAll(r)
			output := string(out)

			// Compare output
			if output != tt.output {
				t.Errorf("Expected output:\n%s\nGot:\n%s", tt.output, output)
			}
		})
	}
}

// Helper function to capture stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
