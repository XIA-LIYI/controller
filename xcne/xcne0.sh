#!/bin/sh
#SBATCH --time=10
#SBATCH --nodelist=xcne0
#SBATCH --ntasks=1 --cpus-per-task=10

srun ../client