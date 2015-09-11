package enocean

import (
	"errors"
)

// constants for Manufacturers ID
const (
	MANUFACTURER_RESERVED                  byte = 0x00
	PEHA                                   byte = 0x01
	THERMOKON                              byte = 0x02
	SERVODAN                               byte = 0x03
	ECHOFLEX_SOLUTIONS                     byte = 0x04
	OMNIO_AG                               byte = 0x05
	HARDMEIER_ELECTRONICS                  byte = 0x06
	REGULVAR_INC                           byte = 0x07
	AD_HOC_ELECTRONICS                     byte = 0x08
	DISTECH_CONTROLS                       byte = 0x09
	KIEBACK_AND_PETER                      byte = 0x0A
	ENOCEAN_GMBH                           byte = 0x0B
	PROBARE                                byte = 0x0C
	ELTAKO                                 byte = 0x0D
	LEVITON                                byte = 0x0E
	HONEYWELL                              byte = 0x0F
	SPARTAN_PERIPHERAL_DEVICES             byte = 0x10
	SIEMENS                                byte = 0x11
	T_MAC                                  byte = 0x12
	RELIABLE_CONTROLS_CORPORATION          byte = 0x13
	ELSNER_ELEKTRONIK_GMBH                 byte = 0x14
	DIEHL_CONTROLS                         byte = 0x15
	BSC_COMPUTER                           byte = 0x16
	S_AND_S_REGELTECHNIK_GMBH              byte = 0x17
	MASCO_CORPORATION                      byte = 0x18
	INTESIS_SOFTWARE_SL                    byte = 0x19
	VIESSMANN                              byte = 0x1A
	LUTUO_TECHNOLOGY                       byte = 0x1B
	SCHNEIDER_ELECTRIC                     byte = 0x1C
	SAUTER                                 byte = 0x1D
	BOOT_UP                                byte = 0x1E
	OSRAM_SYLVANIA                         byte = 0x1F
	UNOTECH                                byte = 0x20
	DELTA_CONTROLS_INC                     byte = 0x21
	UNITRONIC_AG                           byte = 0x22
	NANOSENSE                              byte = 0x23
	THE_S4_GROUP                           byte = 0x24
	MSR_SOLUTIONS                          byte = 0x25
	GE                                     byte = 0x26
	MAICO                                  byte = 0x27
	RUSKIN_COMPANY                         byte = 0x28
	MAGNUM_ENERGY_SOLUTIONS                byte = 0x29
	KMC_CONTROLS                           byte = 0x2A
	ECOLOGIX_CONTROLS                      byte = 0x2B
	TRIO_2_SYS                             byte = 0x2C
	AFRISO_EURO_INDEX                      byte = 0x2D
	NEC_ACCESSTECHNICA_LTD                 byte = 0x30
	ITEC_CORPORATION                       byte = 0x31
	SIMICX_CO_LTD                          byte = 0x32
	EUROTRONIC_TECHNOLOGY_GMBH             byte = 0x34
	ART_JAPAN_CO_LTD                       byte = 0x35
	TIANSU_AUTOMATION_CONTROL_SYSTE_CO_LTD byte = 0x36
	GRUPPO_GIORDANO_IDEA_SPA               byte = 0x38
	ALPHAEOS_AG                            byte = 0x39
	TAG_TECHNOLOGIES                       byte = 0x3A
	CLOUD_BUILDINGS_LTD                    byte = 0x3C
	GIGA_CONCEPT                           byte = 0x3E
	SENSORTEC                              byte = 0x3F
	JAEGER_DIREKT                          byte = 0x40
	AIR_SYSTEM_COMPONENTS_INC              byte = 0x41
	MULTI_USER_MANUFACTURER                int  = 0x7FF
)

var manufacturerNames = []string{
	"MANUFACTURER_RESERVED",
	"PEHA",
	"THERMOKON",
	"SERVODAN",
	"ECHOFLEX_SOLUTIONS",
	"OMNIO_AG",
	"HARDMEIER_ELECTRONICS",
	"REGULVAR_INC",
	"AD_HOC_ELECTRONICS",
	"DISTECH_CONTROLS",
	"KIEBACK_AND_PETER",
	"ENOCEAN_GMBH",
	"PROBARE",
	"ELTAKO",
	"LEVITON",
	"HONEYWELL",
	"SPARTAN_PERIPHERAL_DEVICES",
	"SIEMENS",
	"T_MAC",
	"RELIABLE_CONTROLS_CORPORATION",
	"ELSNER_ELEKTRONIK_GMBH",
	"DIEHL_CONTROLS",
	"BSC_COMPUTER",
	"S_AND_S_REGELTECHNIK_GMBH",
	"MASCO_CORPORATION",
	"INTESIS_SOFTWARE_SL",
	"VIESSMANN",
	"LUTUO_TECHNOLOGY",
	"SCHNEIDER_ELECTRIC",
	"SAUTER",
	"BOOT_UP",
	"OSRAM_SYLVANIA",
	"UNOTECH",
	"DELTA_CONTROLS_INC",
	"UNITRONIC_AG",
	"NANOSENSE",
	"THE_S4_GROUP",
	"MSR_SOLUTIONS",
	"GE",
	"MAICO",
	"RUSKIN_COMPANY",
	"MAGNUM_ENERGY_SOLUTIONS",
	"KMC_CONTROLS",
	"ECOLOGIX_CONTROLS",
	"TRIO_2_SYS",
	"AFRISO_EURO_INDEX",
	"",
	"",
	"NEC_ACCESSTECHNICA_LTD",
	"ITEC_CORPORATION",
	"SIMICX_CO_LTD",
	"",
	"EUROTRONIC_TECHNOLOGY_GMBH",
	"ART_JAPAN_CO_LTD",
	"TIANSU_AUTOMATION_CONTROL_SYSTE_CO_LTD",
	"",
	"GRUPPO_GIORDANO_IDEA_SPA",
	"ALPHAEOS_AG",
	"TAG_TECHNOLOGIES",
	"",
	"CLOUD_BUILDINGS_LTD",
	"",
	"GIGA_CONCEPT",
	"SENSORTEC",
	"JAEGER_DIREKT",
	"AIR_SYSTEM_COMPONENTS_INC",
}

func GetManufacturerName(id int) (error, string) {
	if id == 0x7ff {
		return nil, "AIR_SYSTEM_COMPONENTS_INC"
	}
	if id > len(manufacturerNames) {
		return errors.New("Manufacturer ID is out of range"), ""
	}
	return nil, manufacturerNames[id]
}
