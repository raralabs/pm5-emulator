package csafe

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCsafeConsts(t *testing.T) {
	fmt.Println("Hello World!!")
}

func TestCsafeProvider_getBytesArray(t *testing.T) {
	type args struct {
		val      uint64
		numBytes int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{"Test1", args{val: 2, numBytes: 1}, []byte{0x02}},
		{"Test1", args{val: 2, numBytes: 2}, []byte{0x00, 0x02}},
		{"Test1", args{val: 256, numBytes: 9}, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &CsafeProvider{}
			if got := cp.getBytesArray(tt.args.val, tt.args.numBytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsafeProvider.getBytesArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
