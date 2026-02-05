package parsers

// AI generated
type BIOSCharacteristics struct {
	Reserved0             bool
	Reserved1             bool
	Unknown               bool
	NotSupported          bool
	ISASupported          bool
	MCASupported          bool
	EISASupported         bool
	PCISupported          bool
	PCCardSupported       bool
	PlugAndPlay           bool
	APMSupported          bool
	BIOSUpgradeable       bool
	ShadowingAllowed      bool
	VLVESASupported       bool
	ESCDSupported         bool
	BootFromCD            bool
	SelectableBoot        bool
	BIOSROMSocketed       bool
	BootFromPCCard        bool
	EDDSupported          bool
	JapaneseFloppyNEC     bool
	JapaneseFloppyToshiba bool
	Floppy5_25_360KB      bool
	Floppy5_25_1_2MB      bool
	Floppy3_5_720KB       bool
	Floppy3_5_2_88MB      bool
	PrintScreenService    bool
	KeyboardService       bool
	SerialService         bool
	PrinterService        bool
	VideoService          bool
	NECPC98               bool
	VendorReserved        uint16
	SystemReserved        uint16
}

func ParseBIOSCharacteristics(mask uint64) BIOSCharacteristics {
	return BIOSCharacteristics{
		Reserved0:             (mask & (1 << 0)) != 0,
		Reserved1:             (mask & (1 << 1)) != 0,
		Unknown:               (mask & (1 << 2)) != 0,
		NotSupported:          (mask & (1 << 3)) != 0,
		ISASupported:          (mask & (1 << 4)) != 0,
		MCASupported:          (mask & (1 << 5)) != 0,
		EISASupported:         (mask & (1 << 6)) != 0,
		PCISupported:          (mask & (1 << 7)) != 0,
		PCCardSupported:       (mask & (1 << 8)) != 0,
		PlugAndPlay:           (mask & (1 << 9)) != 0,
		APMSupported:          (mask & (1 << 10)) != 0,
		BIOSUpgradeable:       (mask & (1 << 11)) != 0,
		ShadowingAllowed:      (mask & (1 << 12)) != 0,
		VLVESASupported:       (mask & (1 << 13)) != 0,
		ESCDSupported:         (mask & (1 << 14)) != 0,
		BootFromCD:            (mask & (1 << 15)) != 0,
		SelectableBoot:        (mask & (1 << 16)) != 0,
		BIOSROMSocketed:       (mask & (1 << 17)) != 0,
		BootFromPCCard:        (mask & (1 << 18)) != 0,
		EDDSupported:          (mask & (1 << 19)) != 0,
		JapaneseFloppyNEC:     (mask & (1 << 20)) != 0,
		JapaneseFloppyToshiba: (mask & (1 << 21)) != 0,
		Floppy5_25_360KB:      (mask & (1 << 22)) != 0,
		Floppy5_25_1_2MB:      (mask & (1 << 23)) != 0,
		Floppy3_5_720KB:       (mask & (1 << 24)) != 0,
		Floppy3_5_2_88MB:      (mask & (1 << 25)) != 0,
		PrintScreenService:    (mask & (1 << 26)) != 0,
		KeyboardService:       (mask & (1 << 27)) != 0,
		SerialService:         (mask & (1 << 28)) != 0,
		PrinterService:        (mask & (1 << 29)) != 0,
		VideoService:          (mask & (1 << 30)) != 0,
		NECPC98:               (mask & (1 << 31)) != 0,
		VendorReserved:        uint16((mask >> 32) & 0xFFFF),
		SystemReserved:        uint16((mask >> 48) & 0xFFFF),
	}
}

type BIOSCharacteristicsExtended struct {
	//byte 1
	ACPISupported      bool
	USBLegacySupported bool
	AGPSupported       bool
	I2OBootSupported   bool
	LS120SuperDiskBoot bool
	ATAPIZIPBoot       bool
	IEEE1394Boot       bool
	SmartBattery       bool
	//byte 2
	BIOSBootSpecSupported       bool
	FunctionKeyNetworkBoot      bool
	TargetedContentDistribution bool
	UEFISupported               bool
	VirtualMachine              bool
	Reserved5                   bool
	Reserved6                   bool
	Reserved7                   bool
}

// End AI generation

func ParseBiosExtendedCharacteristics(mask uint16) BIOSCharacteristicsExtended {
	byte1 := byte((mask >> 6))
	byte2 := byte(mask & 0xff)

	return BIOSCharacteristicsExtended{
		ACPISupported:      (byte1 & (1 << 0)) != 0,
		USBLegacySupported: (byte1 & (1 << 1)) != 0,
		AGPSupported:       (byte1 & (1 << 2)) != 0,
		I2OBootSupported:   (byte1 & (1 << 3)) != 0,
		LS120SuperDiskBoot: (byte1 & (1 << 4)) != 0,
		ATAPIZIPBoot:       (byte1 & (1 << 5)) != 0,
		IEEE1394Boot:       (byte1 & (1 << 6)) != 0,
		SmartBattery:       (byte1 & (1 << 7)) != 0,

		BIOSBootSpecSupported:       (byte2 & (1 << 0)) != 0,
		FunctionKeyNetworkBoot:      (byte2 & (1 << 1)) != 0,
		TargetedContentDistribution: (byte2 & (1 << 2)) != 0,
		UEFISupported:               (byte2 & (1 << 3)) != 0,
		VirtualMachine:              (byte2 & (1 << 4)) != 0,
		Reserved5:                   (byte2 & (1 << 5)) != 0,
		Reserved6:                   (byte2 & (1 << 6)) != 0,
		Reserved7:                   (byte2 & (1 << 7)) != 0,
	}

}
