package eps

import (
	"fmt"
	"log"
	"strings"

	utility "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Utility"
)

type EntryPointStruct struct {
	Version int
	EPS_2   *EntryPointStruct_2_1
	EPS_3   *EntryPointStruct_3_0
}

/*
SMBios 3.3 has 2 entry points structs,
	2.1 and 3.0
	2.1 is interchangeable with 2.7.1, so im assuming all the 2.x headers are the same
*/

func NewEPS(eps_bytes []byte) *EntryPointStruct {
	eps_struct := &EntryPointStruct{}

	//5F 53 4D 33 5F first 5 bytes of 3.X header
	//5F 53 4D 5F	 first 4 bytes if 2.X header
	is_3_0 := eps_bytes[3] == 0x33
	if is_3_0 {
		eps_struct.Version = 3
		eps3 := &EntryPointStruct_3_0{}
		err := utility.ReadIntoStruct(eps_bytes, eps3)
		if err != nil {
			log.Printf("Failed to load 3.0 entry point: %+v\n", err)
			return eps_struct
		}
		eps_struct.EPS_3 = eps3
	} else {
		eps_struct.Version = 2
		eps2 := &EntryPointStruct_2_1{}
		err := utility.ReadIntoStruct(eps_bytes, eps2)
		if err != nil {
			log.Printf("Failed to load 2.0 entry point: %+v\n", err)
			return eps_struct
		}
		eps_struct.EPS_2 = eps2
	}
	return eps_struct
}
func (eps *EntryPointStruct) VerifyDMITable(table []byte) error {
	var log_str strings.Builder
	log_str.WriteString(fmt.Sprintf("\n--Verifying Table--\n  Table Length: %d\n  ", len(table)))
	switch eps.Version {
	case 2:
		log_str.WriteString(fmt.Sprintf("  EPS Length: %d\n", int(eps.EPS_2.StructTableLen)))
		log.Println(log_str.String())
		if eps.EPS_2.StructTableLen != uint16(len(table)) {
			return fmt.Errorf("EPS table length and actual table length varry")
		}
		return nil
	case 3:
		log_str.WriteString(fmt.Sprintf("  EPS Length: %d\n", int(eps.EPS_3.MaxTableSize)))
		log.Println(log_str.String())
		if eps.EPS_3.MaxTableSize >= uint32(len(table)) {
			return fmt.Errorf("EPS table length and actual table length varry")
		}
		return nil
	default:
		return fmt.Errorf("Unsupported version")
	}
}
