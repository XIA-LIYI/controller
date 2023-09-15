#!/bin/sh
#SBATCH --time=2
#SBATCH --nodelist=xcne1

srun ./iperf -s -p 5050