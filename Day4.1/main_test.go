package main

import (
	"regexp"
	"testing"
)

func Test_winner(t *testing.T) {
	type args struct {
		board string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Vertical Win First",
			args: args{
				board: `   24 66 11 51
   32 37 83 85
   46 59 14 76
   71 28 94 98
   16 40 74  6
`,
			},
			want: true,
		},
		{
			name: "Vertical Win Second",
			args: args{
				board: `24    66 11 51
32    37 83 85
46    59 14 76
71    28 94 98
16    40 74  6
`,
			},
			want: true,
		},
		{
			name: "Vertical Win Third",
			args: args{
				board: `24 66    11 51
32 37    83 85
46 59    14 76
71 28    94 98
16 40    74  6
`,
			},
			want: true,
		},
		{
			name: "Vertical Win Fourth",
			args: args{
				board: `24 66 11    51
32 37 83    85
46 59 14    76
71 28 94    98
16 40 74     6
`,
			},
			want: true,
		},
		{
			name: "Vertical Win Last",
			args: args{
				board: `24 66 11 51   
32 37 83 85   
46 59 14 76   
71 28 94 98   
16 40 74  6   
`,
			},
			want: true,
		},
		{
			name: "Horizontal Win First",
			args: args{
				board: `              
32 37 83 85 24
46 59 14 76 66
71 28 94 98 11
16 40 74  6 51
`,
			},
			want: true,
		},
		{
			name: "Horizontal Win Second",
			args: args{
				board: `32 37 83 85 24
              
46 59 14 76 66
71 28 94 98 11
16 40 74  6 51
`,
			},
			want: true,
		},
		{
			name: "Horizontal Win Third",
			args: args{
				board: `32 37 83 85 24
46 59 14 76 66
              
71 28 94 98 11
16 40 74  6 51
`,
			},
			want: true,
		},
		{
			name: "Horizontal Win Fourth",
			args: args{
				board: `32 37 83 85 24
46 59 14 76 66
71 28 94 98 11
              
16 40 74  6 51
`,
			},
			want: true,
		},
		{
			name: "Horizontal Win Last",
			args: args{
				board: `32 37 83 85 24
46 59 14 76 66
71 28 94 98 11
16 40 74  6 51
              
`,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winner(tt.args.board); got != tt.want {
				t.Errorf("winner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mark(t *testing.T) {
	type args struct {
		board string
		repl  *regexp.Regexp
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Middle of string",
			args: args{
				board: `24 66 11  9 51
32 37 83 35 85
46 59 14 97 76
71 28 94  8 98
16 40 74 52  6
`,
				repl: regexp.MustCompile("35( |\n)"),
			},
			want: `24 66 11  9 51
32 37 83    85
46 59 14 97 76
71 28 94  8 98
16 40 74 52  6
`,
		},
		{
			name: "End of string",
			args: args{
				board: `24 66 11  9 51
32 37 83 35 85
46 59 14 97 76
71 28 94  8 98
16 40 74 52  6
`,
				repl: regexp.MustCompile("85( |\n)"),
			},
			want: `24 66 11  9 51
32 37 83 35   
46 59 14 97 76
71 28 94  8 98
16 40 74 52  6
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mark(tt.args.board, tt.args.repl); got != tt.want {
				t.Errorf("mark() = %v, want %v", got, tt.want)
			}
		})
	}
}
