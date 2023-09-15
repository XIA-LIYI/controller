#!/bin/sh
#SBATCH --time=1
#SBATCH --nodelist=xcne1
srun ./iperf -c 192.168.51.85 -p 5050 -i 2