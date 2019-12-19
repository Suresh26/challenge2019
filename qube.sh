#!/bin/bash

if [ -e main ]
then
    ./main ${1} ${2} {3}
else
    ./setup
    ./main ${1}
fi