package protocol

import "pm5-emulator/protocol/csafe"

//NewProtocol factory provides new Protocol instance
func NewProtocol() (Protocol, error) {
	return &csafe.CSAFE{}, nil
}
