package main

import (
	"flag"

	"github.com/openconfig/ygot/ygot"
	"ygottesting.com/pkg/connectgnmi"
	"ygottesting.com/pkg/render"
)

func main() {
	Target := flag.String("target", "172.20.20.2", "gnmi target")
	Port := flag.String("port", "6030", "gNMI port default is 6030")
	Username := flag.String("username", "admin", "admin")
	Password := flag.String("password", "admin", "admin")
	NtpServerAddress := flag.String("ntpserveraddress", "", "Address in which you want to render a NtpServerAddress")
	SetNtpAddress := flag.String("setntpaddress", "", "Address in which you want to set as the ntp server.")
	DeleteNtpAddress := flag.String("deletentpaddress", "", "ntp server you want to delete")
	Subscribe := flag.Bool("subscribe", false, "Subscribe method to the ntp servers")
	flag.Parse()
	if *NtpServerAddress != "" {
		render.CreateNtpJson(ygot.String(*NtpServerAddress))
	}
	if *SetNtpAddress != "" {
		connectgnmi.Set(*Target, *Port, *Username, *Password, *SetNtpAddress)
	}
	if *DeleteNtpAddress != "" {
		connectgnmi.Delete(*Target, *Port, *Username, *Password, *DeleteNtpAddress)
	}
	if *Subscribe == true {
		connectgnmi.Subscribe(*Target, *Port, *Username, *Password)
	} else {
		connectgnmi.Get(*Target, *Port, *Username, *Password)
	}
}
