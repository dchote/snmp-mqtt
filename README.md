# snmp-mqtt

A simple go app that reads SNMP values and publishes it to the specified MQTT endpoint at the specified interval.

Please download the precompiled binary from the releases page: https://github.com/dchote/snmp-mqtt/releases

```
Usage: snmp-mqtt [options]

Options:
  --endpoints_map=<endpoints_map>     SNMP Endpoints Map File [default: ./endpoints.json]
  --server=<server>                   MQTT server host/IP [default: 127.0.0.1]
  --port=<port>                       MQTT server port [default: 1883]
  --topic=<topic>                     MQTT topic prefix [default: snmp]
  --clientid=<clientid>               MQTT client identifier [default: snmp]
  --interval=<interval>               Poll interval (seconds) [default: 5]
  -h, --help                          Show this screen.
  -v, --version                       Show version.
```

An example endpoints.json file:
```
{
  "snmpEndpoints": [
    {
      "endpoint": "172.18.0.1",
      "community": "public",
      "oidTopics": [
        {
          "oid": ".1.3.6.1.2.1.31.1.1.1.6.4",
          "topic": "router/bytesIn"
        },
        {
          "oid": ".1.3.6.1.2.1.31.1.1.1.10.4",
          "topic": "router/bytesOut"
        }
      ]
    }
 ]
}
```