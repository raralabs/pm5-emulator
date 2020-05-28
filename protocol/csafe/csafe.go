package csafe

type Command []byte

type Packet struct {
	Data []byte
	Cmd  byte
}

//CSAFE
type CSAFE struct {
}

//ReadPayload reads frame contents
func (c *CSAFE) ReadPayload(cmd Command) error {
	//implements byte unstuffing
	//checksum
	//frame content decode
	return nil
}

func (c *CSAFE) WritePayload(cmd Command) (Command, error) {
	return []byte{}, nil
}
