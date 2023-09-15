#!/bin/sh
#SBATCH --time=1
#SBATCH --nodelist=xcne2
srun ./iperf -c 192.168.51.84 -p 5050 -i 2