/* ASN.1 DaytimeClient
 */

package main

import (
	// "bytes"
	"encoding/asn1"
	"fmt"
	// "io"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	defer conn.Close()

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	var newtime time.Time
	_, err1 := asn1.Unmarshal(result, &newtime)
	checkError(err1)

	fmt.Println("Before marshal/unmarshal:", result)
	fmt.Println("After marshal/unmarshal:", newtime.String())
	// example output:
	// Before marshal/unmarshal: [23 17 49 55 49 49 51 48 49 56 50 57 50 54 45 48 56 48 48]
	// After marshal/unmarshal: 2017-11-30 18:29:26 -0800 PST

	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
