#!/bin/sh
#SBATCH --time=2
#SBATCH --nodelist=xcnf7

srun ./iperf -s -p 15000