package csafe

import "encoding/binary"

// Encoder can encode a payload according to the csafe format.
type Encoder struct {
}

// Performs the byte stuffing operation if a payload contains byte greater than or equal to 0xF3
func (cp *Encoder) byteStuffing(payload []byte) []byte {
	var buffer []byte
	for i := 0; i < len(payload); i++ {
		curByte := payload[i]
		if (curByte & 0b11111100) == 0b11110000 {
			buffer = append(buffer, 0xF3, curByte&0b00000011)
		} else {
			buffer = append(buffer, curByte)
		}
	}
	return buffer
}

// Converts a number to an array of provided number of bytes
func (cp *Encoder) getBytesArray(val uint64, numBytes int) []byte {

	//* Original Implementation as in Rowforge code

	// bin, _ := padStart(strconv.FormatInt(val, 2), "0", numBytes*8)
	// ret := make([]byte, numBytes)

	// // get array from msb to lsb
	// for i := 0; i < numBytes; i++ {
	// 	low := i * 8
	// 	high := low + 8
	// 	v, _ := strconv.ParseInt(bin[low:high], 2, 8)
	// 	ret[i] = byte(v)
	// }
	// return ret

	//? Probably better implementation. Need to benchmark

	ret := make([]byte, 8)
	binary.BigEndian.PutUint64(ret, val)

	if numBytes <= 8 {
		ret = ret[8-numBytes:]
	} else {
		// The maximum bytes used by uint64 is 8, so add zeros if asked for longer byte array
		zeros := make([]byte, numBytes-8)
		ret = append(zeros, ret...)
	}

	return ret
}

// getType returns the type (0x00 for "T" and 0x80 for others)
func (cp *Encoder) getType(tpe string) byte {
	tp := byte(0x80)
	if tpe == "T" {
		tp = 0x00
	}
	return tp
}

// Creates a payload for the provided command and the data, and returns it
func (cp *Encoder) Encode(p Packet) []byte {
	var buffer []byte // The Payload

	if len(p.Data) > 90 {
		panic("Can only send max 90 data at a time")
	}

	buffer = append(buffer, p.Cmds...) // Commands

	if !p.JustCmd {
		buffer = append(buffer, byte(len(p.Data))) // Data Byte Count

		if len(p.Data) > 0 {
			buffer = append(buffer, p.Data...) // data bytes
		}
	}

	buffer = append(buffer, calculateChecksum(buffer)) // Insert checksum
	buffer = cp.byteStuffing(buffer)                   // Stuff bytes properly

	buffer = append([]byte{FRAME_START_BYTE}, buffer...) // Frame Start Flag
	buffer = append(buffer, FRAME_END_BYTE)              // Stop Frame Flag

	return buffer
}

// EncodeResponse encodes a response packet
func (cp *Encoder) EncodeResponse(rp ResponsePacket) []byte {
	cmds := append([]byte{rp.Status}, rp.CommandResponseData...)
	cmds = append(cmds, rp.Identifier)

	pck := Packet{
		Data:    rp.Data,
		Cmds:    cmds,
		JustCmd: rp.JustCmd,
	}

	return cp.Encode(pck)
}
