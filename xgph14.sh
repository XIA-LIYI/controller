#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=long
#SBATCH --nodelist=xgph14
#SBATCH --ntasks=1 --cpus-per-task=10

srun ./client