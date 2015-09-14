package enocean

import (
	"testing"
)

func TestManufacturerIdRange(t *testing.T) {
	for _, id := range []int{
		int(MANUFACTURER_RESERVED),
		int(PEHA),
		int(THERMOKON),
		int(SERVODAN),
		int(ECHOFLEX_SOLUTIONS),
		int(OMNIO_AG),
		int(HARDMEIER_ELECTRONICS),
		int(REGULVAR_INC),
		int(AD_HOC_ELECTRONICS),
		int(DISTECH_CONTROLS),
		int(KIEBACK_AND_PETER),
		int(ENOCEAN_GMBH),
		int(PROBARE),
		int(ELTAKO),
		int(LEVITON),
		int(HONEYWELL),
		int(SPARTAN_PERIPHERAL_DEVICES),
		int(SIEMENS),
		int(T_MAC),
		int(RELIABLE_CONTROLS_CORPORATION),
		int(ELSNER_ELEKTRONIK_GMBH),
		int(DIEHL_CONTROLS),
		int(BSC_COMPUTER),
		int(S_AND_S_REGELTECHNIK_GMBH),
		int(MASCO_CORPORATION),
		int(INTESIS_SOFTWARE_SL),
		int(VIESSMANN),
		int(LUTUO_TECHNOLOGY),
		int(SCHNEIDER_ELECTRIC),
		int(SAUTER),
		int(BOOT_UP),
		int(OSRAM_SYLVANIA),
		int(UNOTECH),
		int(DELTA_CONTROLS_INC),
		int(UNITRONIC_AG),
		int(NANOSENSE),
		int(THE_S4_GROUP),
		int(MSR_SOLUTIONS),
		int(GE),
		int(MAICO),
		int(RUSKIN_COMPANY),
		int(MAGNUM_ENERGY_SOLUTIONS),
		int(KMC_CONTROLS),
		int(ECOLOGIX_CONTROLS),
		int(TRIO_2_SYS),
		int(AFRISO_EURO_INDEX),
		int(NEC_ACCESSTECHNICA_LTD),
		int(ITEC_CORPORATION),
		int(SIMICX_CO_LTD),
		int(EUROTRONIC_TECHNOLOGY_GMBH),
		int(ART_JAPAN_CO_LTD),
		int(TIANSU_AUTOMATION_CONTROL_SYSTE_CO_LTD),
		int(GRUPPO_GIORDANO_IDEA_SPA),
		int(ALPHAEOS_AG),
		int(TAG_TECHNOLOGIES),
		int(CLOUD_BUILDINGS_LTD),
		int(GIGA_CONCEPT),
		int(SENSORTEC),
		int(JAEGER_DIREKT),
		int(AIR_SYSTEM_COMPONENTS_INC),
		int(MULTI_USER_MANUFACTURER),
	} {
		err, val := GetManufacturerName(id)
		if err != nil {
			t.Errorf("ID: %v is errored to get its name", id)
		}
		if "" == val {
			t.Errorf("ID: %v does not have its name", id)
		}
	}
}
