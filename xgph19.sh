#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=long
#SBATCH --nodelist=xgph19
#SBATCH --ntasks=1 --cpus-per-task=2100

srun ./client