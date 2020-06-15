package csafe

import (
	"reflect"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {
	tests := []struct {
		name    string
		raw     []byte
		want    *Packet
		wantErr bool
	}{
		// Just Commands tests
		{"Test1", []byte{0xF1, 0x00, 0x00, 0xF2}, &Packet{Data: nil, Cmds: []byte{0x00}, JustCmd: true}, false},
		{"Test2", []byte{0xF1, 0x01, 0x01, 0xF2}, &Packet{Data: nil, Cmds: []byte{0x01}, JustCmd: true}, false},

		// Incorrect frame start or end bytes
		{"Test4", []byte{0xF0, 0x00, 0x00, 0x00, 0xF2}, nil, true},
		{"Test5", []byte{0xF1, 0x00, 0x00, 0x00, 0xF1}, nil, true},

		{"Test With Data", []byte{0xF1, 0x01, 0x01, 0x02, calculateChecksum([]byte{0x01, 0x01, 0x02}), 0xF2},
			&Packet{Data: []byte{0x02}, Cmds: []byte{0x01}, JustCmd: false}, false},
		{"Test ByteStuffing", []byte{0xF1, 0xF3, 0x00, 0x01, 0xF3, 0x01, calculateChecksum([]byte{0xF0, 0x01, 0xF1}), 0xF2},
			&Packet{Data: []byte{0xF1}, Cmds: []byte{0xF0}, JustCmd: false}, false},

		// Incorrect byte stuffing
		{"Test6", []byte{0xF1, 0xF3, 0x00, 0x01, 0xF3, 0xF2},
			nil, true},
		{"Test7", []byte{0xF1, 0xF3, 0x00, 0x01, 0xF3, 0x04, 0xF2},
			nil, true},

		// Incorrect Checksum
		{"Test2", []byte{0xF1, 0x01, 0x02, 0xF2}, nil, true},

		// Less Data than required
		{"Test2", []byte{0xF1, 0x01, 0xF2}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Decoder{}
			got, err := d.Decode(tt.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
