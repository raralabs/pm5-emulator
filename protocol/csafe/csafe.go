package csafe

// Packet represents the most fundamental unit that the devices can use to
// communicate over gatt services.
type Packet struct {
	Data    []byte // Actual data contents
	Cmds    []byte // Commands held by the packet (Can have multiple commands)
	JustCmd bool   // Represents if a packet contains just commands or both cmds and data
}

// ResponsePacket defines the response packet, that a PM5 device responds with,
// to the client's inquiries.
type ResponsePacket struct {
	Status              byte   // The status of PM5 device
	CommandResponseData []byte // Specific data in response to incoming inquiry
	Identifier          byte   // Identifier of the response packet
	Data                []byte // Additional data to be sent to client
	JustCmd             bool   // Represents if just commands are to be sent
}
