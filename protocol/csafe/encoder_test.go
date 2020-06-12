package csafe

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsafeConsts(t *testing.T) {
	fmt.Println("Hello World!!")
}

func TestEncoder_getBytesArray(t *testing.T) {
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
			cp := &Encoder{}
			if got := cp.getBytesArray(tt.args.val, tt.args.numBytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encoder.getBytesArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoder_Encode(t *testing.T) {
	type args struct {
		p Packet
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{"Test1: Just Commands", args{Packet{
			Cmds:    []byte{0x01},
			Data:    []byte{},
			JustCmd: true,
		}}, []byte{0xF1, 0x01, 0x01, 0xF2}},

		{"Test2: Data and Commands", args{Packet{
			Cmds:    []byte{0x02, 0x03, 0x04, 0x01},
			Data:    []byte{0x11, 0x22},
			JustCmd: false,
		}}, []byte{0xF1, 0x02, 0x03, 0x04, 0x01, 0x02, 0x11, 0x22,
			calculateChecksum([]byte{0x02, 0x03, 0x04, 0x01, 0x02, 0x11, 0x22}), 0xF2}},

		{"Test3: ByteStuffing", args{Packet{
			Cmds:    []byte{0xF1},
			Data:    []byte{},
			JustCmd: true,
		}}, []byte{0xF1, 0xF3, 0x01, 0xF3, 0x01, 0xF2}},
	}
	cp := Encoder{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cp.Encode(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encoder.Encode() = %v, want %v", got, tt.want)
			}
		})
	}

	// Panic Test
	pck := Packet{
		Cmds:    []byte{0x01},
		Data:    make([]byte, 100),
		JustCmd: false,
	}

	f := func() {
		cp.Encode(pck)
	}

	assert.Panics(t, f)
}

func TestEncoder_EncodeResponse(t *testing.T) {
	type args struct {
		rp ResponsePacket
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{"Test1", args{ResponsePacket{
			Status:              0x02,
			CommandResponseData: []byte{0x03, 0x04},
			Identifier:          0x01,
			Data:                []byte{0x11, 0x22},
			JustCmd:             false,
		}}, []byte{0xF1, 0x02, 0x03, 0x04, 0x01, 0x02, 0x11, 0x22,
			calculateChecksum([]byte{0x02, 0x03, 0x04, 0x01, 0x02, 0x11, 0x22}), 0xF2}},
	}
	cp := Encoder{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cp.EncodeResponse(tt.args.rp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encoder.EncodeResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoder_getType(t *testing.T) {
	type args struct {
		tpe string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{"Test1", args{"T"}, 0x00},
		{"Test2", args{"D"}, 0x80},
	}
	cp := Encoder{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cp.getType(tt.args.tpe); got != tt.want {
				t.Errorf("Encoder.getType() = %v, want %v", got, tt.want)
			}
		})
	}
}
