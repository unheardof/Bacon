#!/bin/bash

# Reference: https://stackoverflow.com/questions/192249/how-do-i-parse-command-line-arguments-in-bash
POSITIONAL=()
while [[ $# -gt 0 ]]
do
key="$1"

case $key in
    -p|--dst-port)
    DST_PORT="$2"
    shift # past argument
    shift # past value
    ;;
    -h|--dst-host)
    DST_HOST="$2"
    shift # past argument
    shift # past value
    ;;
    -t|--client-token)
    CLIENT_TOKEN="$2"
    shift # past argument
    shift # past value
    ;;
    -i|--callback-interval)
    CALLBACK_INTERVAL="$2"
    shift # past argument
    shift # past value
    ;;
    *)    # unknown option
    POSITIONAL+=("$1") # save it in an array for later
    shift # past argument
    ;;
esac
done
set -- "${POSITIONAL[@]}" # restore positional parameters

echo $DST_PORT
echo $DST_HOST
echo $CLIENT_TOKEN
echo $CALLBACK_INTERVAL
echo ""

# TODO: Validate command line args
#go build -x -ldflags="-X main.dstPort " simple_http_beacon.go
