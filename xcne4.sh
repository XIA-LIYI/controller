#!/bin/sh
#SBATCH --time=10
#SBATCH --nodelist=xcne3
#SBATCH --ntasks=1 --cpus-per-task=20

srun ./client