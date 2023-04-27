package format

import "testing"

func TestReplaceD(t *testing.T) {

	type args struct {
		form      string
		replaceTo string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one replace",
			args: args{form: "solve {d} / {f}", replaceTo: "atcoder"},
			want: "solve atcoder / {f}",
		},
		{
			name: "none replace",
			args: args{form: "solve {f}", replaceTo: "atcoder"},
			want: "solve {f}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceD(tt.args.form, tt.args.replaceTo); got != tt.want {
				t.Errorf("ReplaceD() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestReplaceF(t *testing.T) {

	type args struct {
		form      string
		replaceTo string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one replace",
			args: args{form: "solve atcoder / {f}", replaceTo: "A.txt"},
			want: "solve atcoder / A.txt",
		},
		{
			name: "none replace",
			args: args{form: "solve {d}", replaceTo: "atcoder"},
			want: "solve {d}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceF(tt.args.form, tt.args.replaceTo); got != tt.want {
				t.Errorf("ReplaceF() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestReplaceN(t *testing.T) {

	type args struct {
		form      string
		replaceTo string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one replace",
			args: args{form: "solve atcoder / {n}", replaceTo: "A.txt"},
			want: "solve atcoder / A",
		},
		{
			name: "none replace",
			args: args{form: "solve {d}", replaceTo: "atcoder"},
			want: "solve {d}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceN(tt.args.form, tt.args.replaceTo); got != tt.want {
				t.Errorf("ReplaceN() = %v, want %v", got, tt.want)
			}
		})
	}
}
