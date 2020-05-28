package csafe

import (
	"errors"
)

type Decoder struct {
}

func (d *Decoder) Decode(raw []byte) (*Packet, error) {

	if len(raw) < 5 {
		return nil, errors.New("Raw data length less than minimum length")
	}

	// Remove frame start and end bytes
	body, err := d.stripHeadTail(raw)
	if err != nil {
		return nil, err
	}

	// Perfom reverse byte-stuffing
	pck, err := d.unstuff(body)
	if err != nil {
		return nil, err
	}

	if len(pck) < 3 {
		return nil, errors.New("Useful data info length less than minimum length")
	}

	// Check the checksum
	dta := pck[0 : len(pck)-1]
	checksum := calculateChecksum(dta)
	if checksum != pck[len(pck)-1] {
		return nil, errors.New("Checksum Mismatched")
	}

	// Extract command and the data length
	cmd := pck[0]
	dataLen := pck[1]

	p := &Packet{
		data: make([]byte, dataLen),
		cmd:  cmd,
	}

	// Extract data
	for i := 0; i < int(dataLen); i++ {
		p.data[i] = pck[2+i]
	}

	return p, nil
}

func (d *Decoder) stripHeadTail(raw []byte) ([]byte, error) {
	if raw[0] != FRAME_START_BYTE || raw[len(raw)-1] != FRAME_END_BYTE {
		return raw, errors.New("Not a Packet")
	}

	return raw[1 : len(raw)-1], nil
}

func (d *Decoder) unstuff(raw []byte) ([]byte, error) {
	var buffer []byte

	for i := 0; i < len(raw); i++ {
		curByte := raw[i]
		if curByte == FRAME_STUFF_BYTE {
			if i == len(raw)-1 {
				return raw, errors.New("Unspecified Format")
			}
			buffer = append(buffer, 0xF0|raw[i+1])
			i++
		} else {
			buffer = append(buffer, curByte)
		}
	}

	return buffer, nil
}
