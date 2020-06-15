package csafe

type Command []byte

type CSAFE struct{}

//Packet defines request frame format
type Packet struct {
	Data    []byte
	Cmds    []byte
	JustCmd bool
}

//ResponsePacket defines response frame format
type ResponsePacket struct {
	Status              byte
	CommandResponseData []byte
	Identifier          byte
	Data                []byte
	JustCmd             bool
}
