#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=long
#SBATCH --nodelist=xgpg5
#SBATCH --ntasks=1 --cpus-per-task=5

srun ../client