## This is an example using the same methodology as the binary / go code but with gnmic for ntp servers. 

## Configure a NTP server first with the CLI

```
➜  ygotntpexample git:(main) ✗ docker exec -it clab-initial-lab-ceos1 Cli
ceos1>en
ceos1#conf t
ceos1(config)#ntp server 8.8.8.8
ceos1(config)#
```

### Get all NTP servers using gnmic 
ygotntpexample git:(main) ✗ gnmic -a 172.20.20.2:6030 -u admin -p admin get  --path "/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server" --insecure

```
[
  {
    "source": "172.20.20.2:6030",
    "timestamp": 1697633505054684964,
    "time": "2023-10-18T08:51:45.054684964-04:00",
    "updates": [
      {
        "Path": "system/ntp/servers/server[address=8.8.8.8]",
        "values": {
          "system/ntp/servers/server": {
            "openconfig-system:address": "8.8.8.8",
            "openconfig-system:config": {
              "address": "8.8.8.8"
            },
            "openconfig-system:state": {
              "address": "8.8.8.8"
            }
          }
        }
      }
    ]
  }
]
```


### Get a specific NTP server in this case 8.8.8.8
gnmic -a 172.20.20.2:6030 -u admin -p admin get  --path "/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=8.8.8.8]" --insecure

```
[
  {
    "source": "172.20.20.2:6030",
    "timestamp": 1697633593956661221,
    "time": "2023-10-18T08:53:13.956661221-04:00",
    "updates": [
      {
        "Path": "system/ntp/servers/server[address=8.8.8.8]",
        "values": {
          "system/ntp/servers/server": {
            "openconfig-system:address": "8.8.8.8",
            "openconfig-system:config": {
              "address": "8.8.8.8"
            },
            "openconfig-system:state": {
              "address": "8.8.8.8"
            }
          }
        }
      }
    ]
  }
]
```

### Add a brand new ntp server 1.1.1.1 for example.
gnmic -a 172.20.20.2:6030 -u admin -p admin set --update-path "/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=1.1.1.1]/openconfig-system:config/address" --update-value 1.1.1.1 --insecure

```
{
  "source": "172.20.20.2:6030",
  "timestamp": 1697634018187624032,
  "time": "2023-10-18T09:00:18.187624032-04:00",
  "results": [
    {
      "operation": "UPDATE",
      "path": "openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=1.1.1.1]/openconfig-system:config/address"
    }
  ]
}
```
gnmic -a 172.20.20.2:6030 -u admin -p admin get  --path "/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=1.1.1.1]" --insecure

[
  {
    "source": "172.20.20.2:6030",
    "timestamp": 1697634075966459376,
    "time": "2023-10-18T09:01:15.966459376-04:00",
    "updates": [
      {
        "Path": "system/ntp/servers/server[address=1.1.1.1]",
        "values": {
          "system/ntp/servers/server": {
            "openconfig-system:address": "1.1.1.1",
            "openconfig-system:config": {
              "address": "1.1.1.1"
            },
            "openconfig-system:state": {
              "address": "1.1.1.1"
            }
          }
        }
      }
    ]
  }
]

### Delete 1.1.1.1 NTP server
gnmic -a 172.20.20.2:6030 -u admin -p admin set --delete "/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=1.1.1.1]" --insecure

'''
{
  "source": "172.20.20.2:6030",
  "timestamp": 1697634629411128597,
  "time": "2023-10-18T09:10:29.411128597-04:00",
  "results": [
    {
      "operation": "DELETE",
      "path": "openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=1.1.1.1]"
    }
  ]
}
'''
'''
gnmic -a 172.20.20.2:6030 -u admin -p admin get  --path "/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=1.1.1.1]" --insecure  

[
  {
    "source": "172.20.20.2:6030",
    "timestamp": 1697634655913466867,
    "time": "2023-10-18T09:10:55.913466867-04:00"
  }
]
'''