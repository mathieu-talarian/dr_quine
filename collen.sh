#!/usr/bin/env bash
gcc Collen.c -o Collen
./Collen > tmp
diff tmp Collen.c
