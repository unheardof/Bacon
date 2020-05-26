#!/usr/bin/env python3

import argparse
import subprocess

def main(argv):
    parser = argparse.ArgumentParser()
    parser.add_argument("-p", "--dst-port", help="The port to call back on", type=int)
    parser.add_argument("--h", "--dst-host", help="The host to call back to (hostname or IP)", type=string)
    parser.add_argument("-t", "--client-token", help="Unique identifier for this beacon", type=string)
    parser.add_argument("-i", "--callback-interval", help="How often the becon should call back (in seconds); defaults to 300 seconds (five minutes)", type=string)
    args = parser.parse_args()

    subprocess.run([])

    
    # try:
    #     opts, args = getopt.getopt(argv, "".join(short_options), long_options)
    # except getopt.GetoptError:
    #     print_usage_and_exit(2)


    # # TODO: Init variables to None
        
    # for opt, arg in opts:
    #     if opt in ('-h', '--help'):
    #         print_usage_and_exit(0)
    #     elif opt in ('-i', '--input-file'):
    #         if not os.path.isfile(arg):
    #             print("\nInvalid input file '%s'; file does not exist\n" % arg)
    #             sys.exit(1)
                
    #         input_file = arg
    #     elif opt in ('-o', '--output-file'):
    #         output_file = arg
    #     elif opt in ('-p', '--ports'):
    #         ports_filter = parse_filter_string(arg)
    #     elif opt in ('-a', '--ip-addrs'):
    #         ips_filter = parse_filter_string(arg)
    #     elif opt in ('-s', '--operating-system'):
    #         os_filter = parse_filter_string(arg)
    #     elif opt in ('-d', '--device-type'):
    #         device_type_filter = parse_filter_string(arg)
    #     elif opt in ('-c', '--output-csv'):
    #         output_format = 'csv'
    #     elif opt in ('-q', '--query-mode'):
    #         query_mode = True
    #     else:
    #         print("\nUnknown option '%s' encountered" % (opt))
    #         print_usage_and_exit(0)

