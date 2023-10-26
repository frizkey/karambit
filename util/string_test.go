package util

import "testing"

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

			if got := CoallesceString(tt.args.s...); got != tt.want {
				t.Errorf("CoallesceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
