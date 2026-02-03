# naming convention
I am lazy so instead of files being named by their String name, like BioInfomation or ProcessorInformation<br>
We name them by their struct type defined in DSP0134_3.3.0.pdf, so BioInfomation is just 0, and ProcessorInformation is 4<br>
This helps with organization because the numbered files appear in order when viewed through a file browser<br>
#### see DSP0134_3.3.0.pdf page 28 or header.go for struct type numbers


# Structures
Contains the actual structure definitions

# Parsers
Contains go for parsing bytes into the structs defined in Structures<br>
Also contains 
 * dictionaries for converting binary Enums to the values 
 * functions for analyzing bit shift structures like Bios characteristics

Currently supported Types

### checklist
 * [x] 0 TypeBiosInfo
 * [x] 1 TypeSystemInfo
 * [x] 2 TypeBaseboardInfo
 * [x] 3 TypeSystemEnclosure
 * [x] 4 TypeProcessorInfo
 * [ ] 5 TypeMemoryControllerInfo (Obsolete)
 * [ ] 6 TypeMemoryModuleInfo (Obsolete)
 * [ ] 7 TypeCacheInfo
 * [ ] 8 TypePortConnectorInfo
 * [ ] 9 TypeSystemSlots 
 * [ ] 10 TypeObOardDevicesInfo (Obsolete)
 * [ ] 11 TypeOEMString
 * [ ] 12 TypeSystemConfigOptions
 * [ ] 13 TypeBiosLanguageInfo
 * [ ] 14 TypeGroupAssociations
 * [ ] 15 TypeSystemEventLog
 * [ ] 16 TypePhysicalMemoryArray
 * [ ] 17 TypeMemoryDevice
 * [ ] 18 Type32BitMemoryErrorInfo
 * [ ] 19 TypeMemoryArrayMappedAddress
 * [ ] 20 TypeMemoryDeviceMappedAddress
 * [ ] 21 TypeBuiltinPointingDevice
 * [ ] 22 TypePortableBattery
 * [ ] 23 TypeSystemReset
 * [ ] 24 TypeHardwareSecurity
 * [ ] 25 TypeSystemPowerControls
 * [ ] 26 TypeVoltageProbe
 * [ ] 27 TypeCoolingDevice
 * [ ] 28 TypeTemperatureProbe
 * [ ] 29 TypeElectricalCurrentProbe
 * [ ] 30 TypeOutOfBandRemoteAccess
 * [ ] 31 TypeBootIntegrityServicesEntryPoint
 * [ ] 32 TypeSystemBootInfo
 * [ ] 33 Type64bitMemoryErrorInfo
 * [ ] 34 TypeManagmentDevice
 * [ ] 35 TypeManagementDeviceComponent
 * [ ] 36 TypeManagementDeviceThreashold
 * [ ] 37 TypeMemoryChannel
 * [ ] 38 TypeIPMIDeviceInfo
 
I was doing all this manually, i have given up, theres like 100 of these


###### Ive spelled structures so much that its starting to look weird :(