package protocol

//Protocol
type Protocol interface {
	ReadPayload(payload []byte) error //for reading

	WritePayload(payload []byte) ([]byte, error) //for creating a payload

	//ByteStuffing(cmd csafe.Command) (csafe.Command,error)

	//ByteUnStuffing(cmd csafe.Command) (csafe.Command,error)

	//TODO: add more related methods
}
