package enocean

import (
	"testing"

	"github.com/kgbu/enocean"
)

func TestManufacturerIdRange(t *testing.T) {
	for _, id := range []int{
		int(enocean.MANUFACTURER_RESERVED),
		int(enocean.PEHA),
		int(enocean.THERMOKON),
		int(enocean.SERVODAN),
		int(enocean.ECHOFLEX_SOLUTIONS),
		int(enocean.OMNIO_AG),
		int(enocean.HARDMEIER_ELECTRONICS),
		int(enocean.REGULVAR_INC),
		int(enocean.AD_HOC_ELECTRONICS),
		int(enocean.DISTECH_CONTROLS),
		int(enocean.KIEBACK_AND_PETER),
		int(enocean.ENOCEAN_GMBH),
		int(enocean.PROBARE),
		int(enocean.ELTAKO),
		int(enocean.LEVITON),
		int(enocean.HONEYWELL),
		int(enocean.SPARTAN_PERIPHERAL_DEVICES),
		int(enocean.SIEMENS),
		int(enocean.T_MAC),
		int(enocean.RELIABLE_CONTROLS_CORPORATION),
		int(enocean.ELSNER_ELEKTRONIK_GMBH),
		int(enocean.DIEHL_CONTROLS),
		int(enocean.BSC_COMPUTER),
		int(enocean.S_AND_S_REGELTECHNIK_GMBH),
		int(enocean.MASCO_CORPORATION),
		int(enocean.INTESIS_SOFTWARE_SL),
		int(enocean.VIESSMANN),
		int(enocean.LUTUO_TECHNOLOGY),
		int(enocean.SCHNEIDER_ELECTRIC),
		int(enocean.SAUTER),
		int(enocean.BOOT_UP),
		int(enocean.OSRAM_SYLVANIA),
		int(enocean.UNOTECH),
		int(enocean.DELTA_CONTROLS_INC),
		int(enocean.UNITRONIC_AG),
		int(enocean.NANOSENSE),
		int(enocean.THE_S4_GROUP),
		int(enocean.MSR_SOLUTIONS),
		int(enocean.GE),
		int(enocean.MAICO),
		int(enocean.RUSKIN_COMPANY),
		int(enocean.MAGNUM_ENERGY_SOLUTIONS),
		int(enocean.KMC_CONTROLS),
		int(enocean.ECOLOGIX_CONTROLS),
		int(enocean.TRIO_2_SYS),
		int(enocean.AFRISO_EURO_INDEX),
		int(enocean.NEC_ACCESSTECHNICA_LTD),
		int(enocean.ITEC_CORPORATION),
		int(enocean.SIMICX_CO_LTD),
		int(enocean.EUROTRONIC_TECHNOLOGY_GMBH),
		int(enocean.ART_JAPAN_CO_LTD),
		int(enocean.TIANSU_AUTOMATION_CONTROL_SYSTE_CO_LTD),
		int(enocean.GRUPPO_GIORDANO_IDEA_SPA),
		int(enocean.ALPHAEOS_AG),
		int(enocean.TAG_TECHNOLOGIES),
		int(enocean.CLOUD_BUILDINGS_LTD),
		int(enocean.GIGA_CONCEPT),
		int(enocean.SENSORTEC),
		int(enocean.JAEGER_DIREKT),
		int(enocean.AIR_SYSTEM_COMPONENTS_INC),
		int(enocean.MULTI_USER_MANUFACTURER),
	} {
		err, val := enocean.GetManufacturerName(id)
		if err != nil {
			t.Errorf("ID: %v is errored to get its name", id)
		}
		if "" == val {
			t.Errorf("ID: %v does not have its name", id)
		}
	}
}
