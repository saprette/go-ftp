#!/bin/bash

CHUNCK_SIZE=1024

for i in {1..5}
do
	# for linux
	# cat /dev/urandom | tr -cd 'a-f0-9' | head -c 32

	# for OSX
	cat /dev/urandom | env LC_CTYPE=C tr -cd 'a-f0-9' | head -c $(($i * $CHUNCK_SIZE)) > "file_$i.test"
done

# to test file size : wc -c < file_$i.test 