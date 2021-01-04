#!/bin/bash

mkdir /tmp/SevenShifts
cp ~/Documents/Projects/go/bin/linux_amd64/SevenShifts /tmp/SevenShifts
cp Dockerfile /tmp/SevenShifts
docker build -t SevenShifts /tmp/SevenShifts
rm -rf /tmp/SevenShifts
