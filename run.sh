#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=long
#SBATCH --nodes=25
#SBATCH --ntasks=25 --cpus-per-task=1
#SBATCH --ntasks-per-node=1

srun -n 25 ./client
# srun -N 2 -n 2 ./client
