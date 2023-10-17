package render

import (
	"fmt"

	"github.com/openconfig/ygot/ygot"
	"ygottesting.com/pkg/ocntp"
)

func CreateNtpJson(server *string) string {
	NtpServer := ocntp.System_Ntp_Server{
		Address: ygot.String(*server),
	}
	NtpMap := make(map[string]*ocntp.System_Ntp_Server)
	NtpMap[*server] = &NtpServer
	NtpSys := ocntp.System_Ntp{
		Server: NtpMap,
	}
	Sys := &ocntp.System{Ntp: &NtpSys}
	json, err := ygot.EmitJSON(Sys, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Value error: %v", err))
	}
	fmt.Println("JSON Testing output")
	fmt.Println(json)

	return json
}