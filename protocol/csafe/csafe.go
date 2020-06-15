package csafe

type Command []byte

type CSAFE struct{}

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
