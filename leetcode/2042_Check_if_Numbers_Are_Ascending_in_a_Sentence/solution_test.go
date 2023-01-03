package _042_Check_if_Numbers_Are_Ascending_in_a_Sentence

import "testing"

func Test_areNumbersAscending(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "succ",
			args: args{s: "a puppy has 2 eyes 4 legs"},
			want: true,
		},
		{
			name: "fail",
			args: args{s: "a puppy has 4 eyes 4 legs"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areNumbersAscending(tt.args.s); got != tt.want {
				t.Errorf("areNumbersAscending() = %v, want %v", got, tt.want)
			}
		})
	}
}
