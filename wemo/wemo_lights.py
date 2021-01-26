#!/usr/bin/env python3

import argparse
import pywemo

parser = argparse.ArgumentParser()
parser.add_argument("addr")
parser.add_argument('-p', '--port', dest='port', type=int, default=49153)
group = parser.add_mutually_exclusive_group()
group.add_argument("--on", action="store_true")
group.add_argument("--off", action="store_true")

parser.parse_args()

args = parser.parse_args()
if not args.addr:
    exit()

host = args.addr
port = args.port

url = pywemo.setup_url_for_address(host, port)
d = pywemo.discovery.device_from_description(url, None)

if args.on:
    d.on()
elif args.off:
    d.off()
else:
    print(d.get_state())

