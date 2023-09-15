#!/bin/sh
#SBATCH --time=1
#SBATCH --nodelist=xcne2
bash -c screen ./iperf -s -p 5050
sleep 5
bash -c screen ./iperf -c 192.168.51.84 -p 5050 -i 2