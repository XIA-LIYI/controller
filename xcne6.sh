#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=medium
#SBATCH --nodelist=xcne6
#SBATCH --ntasks=1 --cpus-per-task=20

srun ./client