#!/bin/sh
#SBATCH --time=10
#SBATCH --partition=medium
#SBATCH --nodelist=xcne5
#SBATCH --ntasks=1 --cpus-per-task=10

srun ./client