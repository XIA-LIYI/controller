#!/bin/sh
#SBATCH --time=10
#SBATCH --nodelist=xcne1
#SBATCH --ntasks=1 --cpus-per-task=10

srun ../client