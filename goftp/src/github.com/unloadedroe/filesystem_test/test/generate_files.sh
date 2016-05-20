#!/bin/bash

CHUNK_SIZE=1024

#for i in {1..1024}
#do
	# for linux
	# cat /dev/urandom | tr -cd 'a-f0-9' | head -c 32

	# for OSX
	cat /dev/urandom | env LC_CTYPE=C tr -cd 'a-f0-9' | head -c $(($CHUNK_SIZE * $CHUNK_SIZE)) > "file_$CHUNK_SIZE.test"
#done

# to test file size : wc -c < file_$i.test 
