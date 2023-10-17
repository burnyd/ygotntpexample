## General Demo configuring devices with ygot and gnmic.

Will add more later on. 

## Copy the yang files from the arista public yang file to yang for 4.30.2F

```
./generator -output_file=pkg/ntp/ntp.go -package_name=ocntp -path=yang/ -generate_fakeroot -fakeroot_name=device -exclude_modules=ietf-interfaces -compress_paths=true  yang/openconfig-system.yang
```

## Create json for server 8.8.8.8
go run main.go -ntpserveraddress 8.8.8.8

```
➜  ygotntpexample go run main.go -ntpserveraddress 8.8.8.8

JSON Testing output
{
  "openconfig-system:ntp": {
    "servers": {
      "server": [
        {
          "address": "8.8.8.8",
          "config": {
            "address": "8.8.8.8"
          }
        }
      ]
    }
  }
}```

## Get the ntp server running on device 172.20.20.3
go run main.go -target 172.20.20.3

```
➜  ygotntpexample go run main.go -target 172.20.20.3
NTP Server configured as: 8.8.8.8
```

## Add a ntp server to the target

```
➜  ygotntpexample go run main.go -setntpaddress 1.2.3.4 -target 172.20.20.2

2023/10/17 10:49:54
 trying to update with ntp server address: 1.2.3.4
2023/10/17 10:49:54
 Trying Path: [/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=1.2.3.4]/openconfig-system:config/address]
2023/10/17 10:49:54 ntp server configured: 1.2.3.4
```

## Subscribe to NTP servers

```
ygotntpexample go run main.go -target 172.20.20.2 ~subscribe~=true

[2023-10-17T14:49:54.141624502Z] (172.20.20.2:6030) Update /system/ntp/servers/server[address=1.2.3.4]/address = 1.2.3.4
[2023-10-17T14:49:54.141624502Z] (172.20.20.2:6030) Update /system/ntp/servers/server[address=1.2.3.4]/config/address = 1.2.3.4
[2023-10-17T14:49:54.141624502Z] (172.20.20.2:6030) Update /system/ntp/servers/server[address=1.2.3.4]/state/address = 1.2.3.4
```