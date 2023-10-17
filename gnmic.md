gnmic -a 172.20.20.3:6030 -u admin -p admin set  --update-path "/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=1.1.1.1]/openconfig-system:config/address" --update-value 1.1.1.1 --insecure

{
  "source": "172.20.20.3:6030",
  "timestamp": 1697553326315925704,
  "time": "2023-10-17T10:35:26.315925704-04:00",
  "results": [
    {
      "operation": "UPDATE",
      "path": "openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=1.1.1.1]/openconfig-system:config/address"
    }
  ]
}
