package stringutil

import (
	"testing"
)

func Test_IsCamelCase(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple example",
			args: args{
				val: "fooBar",
			},
			want: true,
		},
		{
			name: "snakecase",
			args: args{
				val: "foo_bar",
			},
			want: false,
		},
		{
			name: "titlecase",
			args: args{
				val: "Foo",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCamelCase(tt.args.val); got != tt.want {
				t.Errorf("IsCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTrimmed(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "trimmed",
			args: args{
				in: "foo",
			},
			want: true,
		},
		{
			name: "leading whitespace",
			args: args{
				in: "  foo",
			},
			want: false,
		},
		{
			name: "trailing whitespace",
			args: args{
				in: "foo  ",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTrimmed(tt.args.in); got != tt.want {
				t.Errorf("IsTrimmed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ContainsIgnoreString(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty string",
			args: args{
				in: "",
			},
			want: false,
		},
		{
			name: "simple example",
			args: args{
				in: "nolint: jsonstructlint",
			},
			want: true,
		},
		{
			name: "comma separated list",
			args: args{
				in: "nolint: foo,jsonstructlint,bar",
			},
			want: true,
		},
		{
			name: "whitespace variant",
			args: args{
				in: "nolint:jsonstructlint",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsIgnoreString(tt.args.in); got != tt.want {
				t.Errorf("ContainsIgnoreString() = %v, want %v", got, tt.want)
			}
		})
	}
}
