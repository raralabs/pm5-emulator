package csafe

import "errors"

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
