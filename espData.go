package enocean

import (
	"encoding/json"
	"log"
)

type ESPdata struct {
	RORG           byte
	FUNC           byte
	TYPE           byte
	ManufacturerId int
	TeachIn        bool
	OriginatorID   []byte
	DestinationID  []byte
	DataPayload    []byte
	RSSI           byte
}

func ToJSON(e ESPData) string {
	return json.Marshal(e)
}

func NewESPData(src []byte) (error, int, ESPData) {
	e = ESPData{}

	// Check Header length : shall be > 6
	if len(src) <= 6 {
		consumedBytes = 0
		return errors.New("too short data length %v", len(src)), consumedBytes, nil
	}

	// Check sync byte
	if src[0] != SyncByte {
		consumedBytes = 1
		return errors.New("Sync Byte does not match. Please shift one byte"), consumedBytes, nil
	}

	// Check data length
	dataLength = (int(src[1]) << 8) + int(src[2])
	optionalDataLength = int(src[3])
	totalLength = ESPHeaderLength + dataLength + optionalDataLength + ESPCRCLength
	if totalLength > len(src) {
		consumedBytes = 0
		return errors.New("too short data length %v than total length: %v", len(src), totalLength), consumedBytes, nil
	}

	// Check packet type is ERP telegram
	if src[4] != 0x0A {
		return errors.New("Unknown packet type %v than 0x0A (Enocean Radio Telegram)"), int(src[4]), totalLength, nil
	}

	// Check Address Controll
	addressControl := (int(src[6]) & 0xE0) >> 5
	extendedHeaderExists := (int(src[6]) & 0x10) == 0x10
	if extendedHeaderExists == true {
		return errors.New("Unknown packet structure which have Extended Header"), totalLength, nil
	}

	payloadPosition = 10
	switch addressControl {
	case 0:
		for i := 7; i < 10; i++ {
			e.OriginatorId = append(e.OriginatorId, src[i])
		}
		payloadPosition = 10
	case 1:
		for i := 7; i < 11; i++ {
			e.OriginatorId = append(e.OriginatorId, src[i])
		}
		payloadPosition = 11
	case 2:
		for i := 7; i < 11; i++ {
			e.OriginatorId = append(e.OriginatorId, src[i])
		}
		for i := 11; i < 15; i++ {
			e.DestinationId = append(e.DestinationId, src[i])
		}
		payloadPosition = 15
	case 3:
		for i := 7; i < 13; i++ {
			e.OriginatorId = append(e.OriginatorId, src[i])
		}
		payloadPosition = 13
	}

	telegramType := src[6] & 0x07
	switch telegramType {
	case 0x0:
		e.RORG = 0xf6
		payloadLength = 1
	case 0x1:
		e.RORG = 0xd5
		payloadLength = 1
	case 0x2:
		e.RORG = 0xa5
		payloadLength = 4
	}

	for i := payloadPosition; i < (payloadPosition + payloadLength); i++ {
		e.PayloadData = append(e.PayloadData, src[i])
	}

	e.TeachIn = int(src[payloadPosition+payloadLength-1])&0x08 == 0x08

	if e.TeachIn && (e.RORG == 0xa5) {
		e.FUNC = int(e.PayloadData[0] && 0xFC) >> 2
		e.TYPE = int(e.PayloadData[0])<<5 + int(e.PayloadData[1])>>3
		e.ManufacturerId = int(e.PayloadData[1] && 0x07)<<8 + int(e.PayloadData[2])
	}

	return nil, totalLength, e
}
