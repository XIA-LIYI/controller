#!/bin/sh
#SBATCH --time=2
#SBATCH --nodelist=xcnf8

srun ./iperf -s -p 50120