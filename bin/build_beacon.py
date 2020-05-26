#!/usr/bin/env python3

import argparse
import subprocess

def main(argv):
    parser = argparse.ArgumentParser()
    parser.add_argument("-p", "--dst-port", help="The port to call back on", type=int) # TODO: Add support for a list of ports to try calling back to
    parser.add_argument("-h", "--dst-host", help="The host to call back to (hostname or IP)", type=string) # TODO: Add support for a list of hostnames to try to call back to
    parser.add_argument("-t", "--client-token", help="Unique identifier for this beacon", type=string)
    parser.add_argument("-c", "--callback-interval", help="How often the beacon should call back (in seconds); defaults to 300 seconds (five minutes)", type=string)
    parser.add_argument("-b", "--beacon-type", help="What type of beacon you want", type=string)
    parser.add_argument("-o", "--output-file", help="Where to save the compiled beacon file to", type=string)
    # TODO: Add target-os argument / add support for cross compiling across multiple platforms
    args = parser.parse_args()

    # TODO: Create build command and then execute it
    subprocess.run([])
