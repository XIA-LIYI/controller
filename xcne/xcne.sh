#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=long
#SBATCH --nodes=2
#SBATCH --nodelist=xcne0, xcne1, xcne2, xcne3, xcne4,xcne5, xcne6, xcne7
#SBATCH --ntasks= --cpus-per-task=1

srun -n 8 ./client
# srun -N 2 -n 2 ./client