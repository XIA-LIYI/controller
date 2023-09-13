#!/bin/sh
#SBATCH --time=1
#SBATCH --nodelist=xcnf6
srun ./iperf -c 192.168.48.135 -p 15000 -i 2