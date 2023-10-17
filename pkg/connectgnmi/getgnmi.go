package connectgnmi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aristanetworks/goarista/gnmi"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

type NtpResponse struct {
	OpenconfigSystemAddress string `json:"openconfig-system:address"`
	OpenconfigSystemConfig  struct {
		Address string `json:"address"`
	} `json:"openconfig-system:config"`
	OpenconfigSystemState struct {
		Address string `json:"address"`
	} `json:"openconfig-system:state"`
}

func GetReq(ctx context.Context, client pb.GNMIClient,
	req *pb.GetRequest) ([]byte, error) {
	resp, err := client.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	var reqreturn []string
	for _, notif := range resp.Notification {
		for _, update := range notif.Update {
			reqreturn = append(reqreturn, (gnmi.StrUpdateVal(update)))
		}
	}
	byteSlice := []byte(strings.Join(reqreturn, " "))
	return byteSlice, nil
}

func Get(Target, Port, Username, Password string) {
	var cfg = &gnmi.Config{
		Addr:     Target + ":" + Port,
		Username: Username,
		Password: Password,
	}
	paths := []string{"/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server"}
	var origin = "openconfig"
	ctx := gnmi.NewContext(context.Background(), cfg)
	client, err := gnmi.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	req, err := gnmi.NewGetRequest(gnmi.SplitPaths(paths), origin)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.Addr != "" {
		if req.Prefix == nil {
			req.Prefix = &pb.Path{}
		}
		req.Prefix.Target = cfg.Addr
	}

	Get, err := GetReq(ctx, client, req)
	if err != nil {
		log.Fatal(err)
	}
	Unmarshall(Get)
}
func Unmarshall(data []byte) {
	loadd := &NtpResponse{}
	_ = json.Unmarshal(data, &loadd)
	if len(loadd.OpenconfigSystemState.Address) == 0 {
		fmt.Println("Zero NTP servers exists")
	} else {
		fmt.Println("NTP Server configured as:", loadd.OpenconfigSystemState.Address)
	}
}
