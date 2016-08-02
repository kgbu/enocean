package enocean

// type of EnOcean Radio Protocol2 Raw Data packet in EnOcean Serial Protocol3

type ERP2Payload struct {
	PayloadType byte
	Identifier  []byte
	Data        []byte
	CRC8        byte
}

type ERP2inESP3 struct {
	ERP2Header       byte
	ERP2ExtendHeader []byte
	Payload          []byte
}

type ERP2 struct {
	Length byte
	Body   ERP2inESP3
}
