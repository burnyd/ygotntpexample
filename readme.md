## Requirements
```
ceoslab 4.29.2F - Used for testing.
containerlab - 0.44.0.
```

## General Demo configuring devices with ygot and gnmic.

```
The gnmic.md file is a example of everything leveraging gnmic instead of the binary built within bin/
```

### Spin up containerlab
<br>
➜  ygotntpexample git:(main) ✗ sudo containerlab -t clab.yml deploy
<br>
<br>

## Copy the yang files from the arista public yang file to yang for 4.30.2F

<br>

[Arista Yang]("https://github.com/aristanetworks/yang")

## Compile the ygot binary
<br>
This is necessary because it will take our yang files from the 4.30.3M directory and make them a go package essentially. 

Copy this ygot [repo]("https://github.com/openconfig/ygot") else where

Within the repo go to the generator directory and issue a 

```
go build
```

I moved the generator to the /bin directory

```
cd bin 

./generator -output_file=pkg/ntp/ntp.go -package_name=ocntp -path=yang/ -generate_fakeroot -fakeroot_name=device -exclude_modules=ietf-interfaces -compress_paths=true  yang/openconfig-system.yang

```

This will then create the pkg/ocntp directory
<br>

## Create json for server 8.8.8.8
./ntpbuild -ntpserveraddress 8.8.8.8

```
➜  ygotntpexample ./ntpbuild -ntpserveraddress 8.8.8.8

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
./ntpbuild -target 172.20.20.3

```
➜  ygotntpexample ./ntpbuild -target 172.20.20.3
NTP Server configured as: 8.8.8.8
```

## Add a ntp server to the target

```
➜  ygotntpexample ./ntpbuild -setntpaddress 1.2.3.4 -target 172.20.20.2

2023/10/17 10:49:54
 trying to update with ntp server address: 1.2.3.4
2023/10/17 10:49:54
 Trying Path: [/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=1.2.3.4]/openconfig-system:config/address]
2023/10/17 10:49:54 ntp server configured: 1.2.3.4
```

## Subscribe to NTP servers

```
ygotntpexample ./ntpbuild -target 172.20.20.2 ~subscribe~=true

[2023-10-17T14:49:54.141624502Z] (172.20.20.2:6030) Update /system/ntp/servers/server[address=1.2.3.4]/address = 1.2.3.4
[2023-10-17T14:49:54.141624502Z] (172.20.20.2:6030) Update /system/ntp/servers/server[address=1.2.3.4]/config/address = 1.2.3.4
[2023-10-17T14:49:54.141624502Z] (172.20.20.2:6030) Update /system/ntp/servers/server[address=1.2.3.4]/state/address = 1.2.3.4
```