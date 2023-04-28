# SNMP

> Great video covering basics, (How SNMP Works | Network Fundamentals)[https://www.youtube.com/watch?v=vWZefoGNk5g]

## Ports

- Port UDP 161, Manager would poll data from this port
- Port UDP 162, SNMP Trap, Sends Manager Alerts

## System Uptime of a Device

The Manager (NMS) will send a Request to the Agent requesting the the System Uptime – the request is sent as a number with the MIB and the Object of Interest, along with something called the Instance.

OID = 1.3.6.1.2.1.1.3.0

## What is MIB?

MIB stands for Management Information Base

## What is OID?

OID is an Object Identifier

## SNMP Walk

- https://www.comparitech.com/net-admin/snmpwalk-examples-windows-linux/
- https://www.ligowave.com/wiki/faq/ligoptp-how-to-find-oids-and-use-mibs/

### All OIDs

```
$ snmpwalk -v2c -c public 192.168.1.59
```

### PDU OIDs

```
$ snmpwalk -v2c -c public 192.168.1.59 1.3.6.1.4
```
