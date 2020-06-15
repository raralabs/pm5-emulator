package csafe

import (
	"errors"
)

// Decoder can decode the raw data - considering it as a csafe-encoded packet.
type Decoder struct {
}

// Decode decodes the raw csafe-encoded data.
func (d *Decoder) Decode(raw []byte) (*Packet, error) {

	if len(raw) < 4 {
		return nil, errors.New("raw data length less than minimum length")
	}

	// Remove frame start and end bytes
	body, err := d.stripHeadTail(raw)
	if err != nil {
		return nil, err
	}

	// Perform reverse byte-stuffing
	pck, err := d.unstuff(body)
	if err != nil {
		return nil, err
	}

	// Check the checksum
	dta := pck[0 : len(pck)-1]
	checksum := calculateChecksum(dta)
	if checksum != pck[len(pck)-1] {
		return nil, errors.New("checksum mismatched")
	}

	// Extract Command
	cmd := pck[0]

	if len(pck) == 2 {
		// Command only
		p := &Packet{
			Data:    nil,
			Cmds:    []byte{cmd},
			JustCmd: true,
		}

		return p, nil
	}

	// Extract the data length
	dataLen := pck[1]

	p := &Packet{
		Data:    make([]byte, dataLen),
		Cmds:    []byte{cmd},
		JustCmd: false,
	}

	// Extract data
	for i := 0; i < int(dataLen); i++ {
		p.Data[i] = pck[2+i]
	}

	return p, nil
}

// stripHeadTail removes the framing head and tail bytes.
func (d *Decoder) stripHeadTail(raw []byte) ([]byte, error) {
	if raw[0] != FRAME_START_BYTE || raw[len(raw)-1] != FRAME_END_BYTE {
		return raw, errors.New("not a packet")
	}

	return raw[1 : len(raw)-1], nil
}

// unstuff performs the reverse operation of csafe byte-stuffing.
func (d *Decoder) unstuff(raw []byte) ([]byte, error) {
	var buffer []byte

	for i := 0; i < len(raw); i++ {
		curByte := raw[i]
		if curByte == FRAME_STUFF_BYTE {
			if (i == len(raw)-1) || (0b11111100&raw[i+1]) != 0 {
				return raw, errors.New("unspecified Format")
			}
			buffer = append(buffer, 0xF0|raw[i+1])
			i++
		} else {
			buffer = append(buffer, curByte)
		}
	}

	return buffer, nil
}
