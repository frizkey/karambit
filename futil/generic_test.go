package futil

import (
	"reflect"
	"testing"
)

func TestCoallesceString(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name    string
		prepare func() args
		args    args
		want    string
	}{
		{
			name: "INITIAL",
			prepare: func() args {
				return args{s: []string{"", "a", "b"}}
			},
			args: args{},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args = tt.prepare()
			if got := CoalesceString(tt.args.s...); got != tt.want {
				t.Errorf("CoallesceString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoalesce(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "INITIAL",
			args: args{values: []interface{}{nil, "a", "b"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Coalesce(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Coalesce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		s   []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "INITIAL",
			args: args{s: []string{"a", "b"}, str: "a"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.str); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "INITIAL",
			args: args{s: "a"},
			want: 3826002220,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.s); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "INITIAL",
			args: args{s: "a"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.s); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonEscape(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "INITIAL",
			args: args{i: "a"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JsonEscape(tt.args.i); got != tt.want {
				t.Errorf("JsonEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrEscape(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "INITIAL",
			args: args{str: "a"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrEscape(tt.args.str); got != tt.want {
				t.Errorf("StrEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToJSONString(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "INITIAL",
			args: args{data: "a"},
			want: "\"a\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToJSONString(tt.args.data); got != tt.want {
				t.Errorf("ToJSONString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "INITIAL",
			args: args{s: "a"},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.args.s); got != tt.want {
				t.Errorf("ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
