#!/bin/bash

tr 'XYZ' 'ABC' < input_2.txt | awk '{print $0" "$2}' | sed -E "s/A$/1\+/g; s/B$/2\+/g; s/C$/3\+/g; s/([A-Z]\ )\1/3\+/g; s/^(C\ A\ |B\ C\ |A\ B\ )/6\+/g; s/([A-Z]\ ){2}/0\+/g; s/.$//" | tr '\n' '+' | sed 's/.$/\n/' | bc

cat input_2.txt | sed -E "s/([A-C]\ )X/\10\+/g; s/([A-C]\ )Y/\13\+/g; s/([A-C]\ )Z/\16\+/g; s/(C\ 6)/7/g; s/(B\ 6)/9/g; s/(A\ 6)/8/g; s/B\ 0/1/g; s/A\ 0/3/g; s/C\ 0/2/g" | tr 'ABC' '123' | tr ' ' '+' | tr '\n' ' ' | sed 's/..$/\n/' | bc
