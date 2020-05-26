package protocol

type Command []byte

//Protocol
type Protocol interface{
	ReadPayload(cmd Command) error  //for reading

	WritePayload(cmd Command) (Command,error) //for creating a payload

	//ByteStuffing(cmd Command) (Command,error)

	//ByteUnStuffing(cmd Command) (Command,error)

	//TODO: add more related methods
}
