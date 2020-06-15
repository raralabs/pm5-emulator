package csafe

import "errors"

// Calculates the checksum of a buffer
func calculateChecksum(buffer []byte) byte {
	checksum := byte(0)
	for i := 0; i < len(buffer); i++ {
		checksum ^= buffer[i]
	}
	return checksum
}

// padStart adds the 'pad' string infront of 'str'
func padStart(str, pad string, maxlength int) (string, error) {
	strlen := len(str)
	padlen := len(pad)

	if strlen+padlen > maxlength {
		return str, errors.New("Could not pad")
	}

	for len(str)+padlen <= maxlength {
		str = pad + str
	}

	return str, nil
}
