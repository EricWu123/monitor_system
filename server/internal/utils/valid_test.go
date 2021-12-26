package utils

import (
	"testing"
)

func TestCheckStrWhite(t *testing.T) {
	type args struct {
		str     string
		pattern string
		maxLen  int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"test-1", args{"peach", `p([a-z]+)ch`, 100}, true, false},
		{"test-2", args{"wyq*", `^[a-z-A-Z0-9]+$`, 100}, false, false},
		{"test-3", args{"wyqAKD09123", `^[a-z-A-Z0-9]+$`, 100}, true, false},
		{"test-4", args{"wyq", `^[a-z-A-Z0-9]+$`, 2}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckStrWhite(tt.args.str, tt.args.pattern, tt.args.maxLen)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckStrWhite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckStrWhite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckStrBlack(t *testing.T) {
	type args struct {
		str     string
		pattern string
		maxLen  int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"test-1", args{"*abc&^%", `[~!@#$%^&*]+`, 100}, false, false},
		{"test-2", args{"abc-()_+=123ASD", `[~!@#$%^&*]+`, 100}, true, false},
		{"test-2", args{"abc-()_+=123ASD", `[~!@#$%^&*]+`, 2}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckStrBlack(tt.args.str, tt.args.pattern, tt.args.maxLen)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckStrBlack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckStrBlack() = %v, want %v", got, tt.want)
			}
		})
	}
}
