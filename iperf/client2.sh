#!/bin/sh
#SBATCH --time=1
#SBATCH --nodelist=xcne2
screen ./iperf -s -p 5050
sleep 5
screen ./iperf -c 192.168.51.84 -p 5050 -i 2