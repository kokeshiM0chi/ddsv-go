package do_test

import (
	"testing"

	"github.com/kokeshiM0chi/ddsv-go/deadlock/rule/do"
	"github.com/kokeshiM0chi/ddsv-go/deadlock/rule/vars"
)

func TestNothing(t *testing.T) {

	tests := []struct {
		name string
		want vars.Shared
	}{
		{
			name: "no var",
			want: vars.Shared{},
		},
		{
			name: "single var",
			want: vars.Shared{"x": 42},
		},
		{
			name: "multiple vars",
			want: vars.Shared{"x": 42, "y": 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := do.Nothing()(tt.want)
			if err != nil {
				t.Fatalf("want no error, but has error %v", err)
			}
			if !eqVars(got, tt.want) {
				t.Fatalf("want %+v, but %+v", tt.want, got)
			}
		})
	}

}

func TestCopyVar(t *testing.T) {

	tests := []struct {
		name      string
		from      vars.Name
		to        vars.Name
		in        vars.Shared
		want      vars.Shared
		wantError bool
	}{
		{
			name: "defined to defined", from: "x", to: "y",
			in:        vars.Shared{"x": 42, "y": 1},
			want:      vars.Shared{"x": 42, "y": 42},
			wantError: false,
		},
		{
			name: "defined to undefined", from: "x", to: "y",
			in:        vars.Shared{"x": 42},
			want:      vars.Shared{},
			wantError: true,
		},
		{
			name: "undefined to defined", from: "x", to: "y",
			in:        vars.Shared{"y": 42},
			want:      vars.Shared{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := do.CopyVar(tt.from).ToVar(tt.to)(tt.in)
			if tt.wantError && err == nil {
				t.Fatalf("want error, but has no error")
			}
			if !tt.wantError && err != nil {
				t.Fatalf("want no error, but has error %v", err)
			}
			if !tt.wantError && !eqVars(got, tt.want) {
				t.Fatalf("want %+v, but %+v", tt.want, got)
			}
		})
	}

}

func TestSet(t *testing.T) {

	tests := []struct {
		name      string
		val       int
		to        vars.Name
		in        vars.Shared
		want      vars.Shared
		wantError bool
	}{
		{
			name: "to defined", val: 42, to: "x",
			in:        vars.Shared{"x": 1},
			want:      vars.Shared{"x": 42},
			wantError: false,
		},
		{
			name: "to undefined", val: 42, to: "x",
			in:        vars.Shared{},
			want:      vars.Shared{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := do.Set(tt.val).ToVar(tt.to)(tt.in)
			if tt.wantError && err == nil {
				t.Fatalf("want error, but has no error")
			}
			if !tt.wantError && err != nil {
				t.Fatalf("want no error, but has error %v", err)
			}
			if !tt.wantError && !eqVars(got, tt.want) {
				t.Fatalf("want %+v, but %+v", tt.want, got)
			}
		})
	}

}

func TestAdd(t *testing.T) {

	tests := []struct {
		name      string
		val       int
		to        vars.Name
		in        vars.Shared
		want      vars.Shared
		wantError bool
	}{
		{
			name: "to defined", val: 42, to: "x",
			in:        vars.Shared{"x": 1},
			want:      vars.Shared{"x": 43},
			wantError: false,
		},
		{
			name: "to undefined", val: 42, to: "x",
			in:        vars.Shared{},
			want:      vars.Shared{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := do.Add(tt.val).ToVar(tt.to)(tt.in)
			if tt.wantError && err == nil {
				t.Fatalf("want error, but has no error")
			}
			if !tt.wantError && err != nil {
				t.Fatalf("want no error, but has error %v", err)
			}
			if !tt.wantError && !eqVars(got, tt.want) {
				t.Fatalf("want %+v, but %+v", tt.want, got)
			}
		})
	}

}

func eqVars(got, want vars.Shared) bool {
	if len(got) != len(want) {
		return false
	}
	for x, _ := range got {
		if got[x] != want[x] {
			return false
		}
	}
	return true
}
