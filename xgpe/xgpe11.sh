#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=long
#SBATCH --nodelist=xgpe11
#SBATCH --ntasks=1 --cpus-per-task=1

srun ./client