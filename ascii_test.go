package main

import (
	"reflect"
	"testing"
)

func TestLoadAsciiChars(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    map[byte][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadAsciiChars(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadAsciiChars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadAsciiChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
