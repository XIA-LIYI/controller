#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=long
#SBATCH --nodes=2
#SBATCH --nodelist=xcne2,xcne3,xcne4,xcne5,xcne6,xcne7,xcne7
#SBATCH --ntasks=7 --cpus-per-task=1

srun -n 8 ./client
# srun -N 2 -n 2 ./client