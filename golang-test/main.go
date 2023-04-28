package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gosnmp/gosnmp"
	snmp "github.com/gosnmp/gosnmp"
)

func main() {

	// get Target and Port from environment
	envTarget := "192.168.1.59"
	envPort := "161"
	oid := "1.3.6.1.4"
	if len(envTarget) <= 0 {
		log.Fatalf("environment variable not set: GOSNMP_TARGET")
	}
	if len(envPort) <= 0 {
		log.Fatalf("environment variable not set: GOSNMP_PORT")
	}
	port, _ := strconv.ParseUint(envPort, 10, 16)

	// Build our own GoSNMP struct, rather than using g.Default.
	// Do verbose logging of packets.
	params := &snmp.GoSNMP{
		Target:    envTarget,
		Port:      uint16(port),
		Community: "public",
		Version:   snmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		Logger:    snmp.NewLogger(log.New(os.Stdout, "", 0)),
	}
	err := params.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer params.Conn.Close()

	// Function handles for collecting metrics on query latencies.
	var sent time.Time
	params.OnSent = func(x *snmp.GoSNMP) {
		sent = time.Now()
	}
	params.OnRecv = func(x *snmp.GoSNMP) {
		log.Println("Query latency in seconds:", time.Since(sent).Seconds())
	}

	err = params.BulkWalk(oid, printValue)
	if err != nil {
		fmt.Printf("Walk Error: %v\n", err)
		os.Exit(1)
	}
}

func printValue(pdu gosnmp.SnmpPDU) error {
	fmt.Printf("%s = ", pdu.Name)

	switch pdu.Type {
	case gosnmp.OctetString:
		b := pdu.Value.([]byte)
		fmt.Printf("STRING: %s\n", string(b))
	default:
		fmt.Printf("TYPE %d: %d\n", pdu.Type, gosnmp.ToBigInt(pdu.Value))
	}
	return nil
}
