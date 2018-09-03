package main

import (
	"testing"
)

func Test_isCamelCase(t *testing.T) {
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
		{
			name: "whitespace",
			args: args{
				val: "foo bar",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCamelCase(tt.args.val); got != tt.want {
				t.Errorf("isCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trim(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple example",
			args: args{
				in: " foo  bar ",
			},
			want: "foobar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trim(tt.args.in); got != tt.want {
				t.Errorf("trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
