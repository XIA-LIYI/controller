#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=long
#SBATCH --nodes=6
#SBATCH --nodelist=xcne2,xcne3,xcne4,xcne5,xcne6,xcne7
#SBATCH --ntasks=6 --cpus-per-task=1

srun -n 6 ./client
# srun -N 2 -n 2 ./client