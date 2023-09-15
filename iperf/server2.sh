#!/bin/sh
#SBATCH --time=2
#SBATCH --nodelist=xcne2

srun ./iperf -s -p 5050