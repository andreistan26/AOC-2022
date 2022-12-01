#!/bin/bash

file='../res/input_1.txt'

#first sol
awk '{ RS="\n\n"; ORS="|"; }{ print }' $file | tr '\n' '+' | tr '|' '\n' | sed 's/+$//g' | bc | sort -n | tail -1 | tr '\n' '+' | sed 's/+$/\n/g' | bc

#second sol
awk '{ RS="\n\n"; ORS="|"; }{ print }' $file | tr '\n' '+' | tr '|' '\n' | sed 's/+$//g' | bc | sort -n | tail -3 | tr '\n' '+' | sed 's/+$/\n/g' | bc

#nth many

#n=10
#awk '{ RS="\n\n"; ORS="|"; }{ print }' $file | tr '\n' '+' | tr '|' '\n' | sed 's/+$//g' | bc | sort -n | tail -$n | tr '\n' '+' | sed 's/+$/\n/g' | bc
