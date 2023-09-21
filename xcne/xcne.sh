#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=long
#SBATCH --nodes=2
#SBATCH --nodelist=xcne4,xcne5
#SBATCH --ntasks=2 --cpus-per-task=1

srun -N 1 -n 1 ./client
srun -N 1 -n 2 ./client