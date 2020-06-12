package csafe

type Command []byte

type Packet struct {
	Data    []byte
	Cmds    []byte
	JustCmd bool
}

type ResponsePacket struct {
	Status              byte
	CommandResponseData []byte
	Identifier          byte
	Data                []byte
	JustCmd             bool
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
