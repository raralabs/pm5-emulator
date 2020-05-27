package csafe

import "testing"

func Test_padStart(t *testing.T) {
	type args struct {
		str       string
		pad       string
		maxlength int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test1", args{str: "1", pad: "0", maxlength: 4}, "0001", false},
		{"Test2", args{str: "1111", pad: "0", maxlength: 4}, "1111", true},
		{"Test3", args{str: "pal", pad: "Ne", maxlength: 6}, "Nepal", false},
		{"Test4", args{str: "", pad: "He", maxlength: 7}, "HeHeHe", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := padStart(tt.args.str, tt.args.pad, tt.args.maxlength)
			if (err != nil) != tt.wantErr {
				t.Errorf("padStart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("padStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
