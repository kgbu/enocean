package enocean

import (
	"fmt"
)

type Payload interface {
	Id() []byte
	Teaching() bool
	Data() []byte
}

type ESP interface {
	Serialize() []byte
	Validate() bool
	Id() []byte
	Payload() Payload
}

const SyncByte byte = 0x55

type ESP3Header struct {
	DataLength         [2]byte
	OptionalDataLength byte
	Type               byte
}

type ESP3Data struct {
	Data         []byte
	OptionalData []byte
}

type ESP3 struct {
	Sync         byte
	Header       ESP3Header
	HeaderCRC8   byte
	Data         ESP3Data
	DataCRC8     byte
}

func (h ESP3.Header) Serialize() []byte {
	ret := []byte {
		h.DataLength[0],
		h.DataLength[1],
		h.OptionalDataLength, 
		h.Type,
	}
	return ret
}

func (d ESP3Data) Serialize() []byte {
	ret := d.Data
	for _, v := range d.OptionalData {
		append(ret, v)
	}
	return ret
}

func (e ESP3) Serialize() []byte {
}

func (esp ESP3) Validate() bool {
	assert.Equal(esp.Sync, SyncByte)
	assert.Equal(CRC8(esp.Header.Serialize()), esp.HeaderCRC8)
	assert.Equal(CRC8(esp.Data.Serialize()), esp.HeaderCRC8)
}

func NewESP3(src byte[]) ESP3 {
	e := ESP3 {}
	return e
}
