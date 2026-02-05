package eps

import (
	"log"

	utility "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Utility"
)

type EntryPointStruct struct {
	EPS_2 *EntryPointStruct_2_1
	EPS_3 *EntryPointStruct_3_0
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
		eps3 := &EntryPointStruct_3_0{}
		err := utility.ReadIntoStruct(eps_bytes, eps3)
		if err != nil {
			log.Printf("Failed to load 3.0 entry point: %+v\n", err)
			return eps_struct
		}
		eps_struct.EPS_3 = eps3
	} else {
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
