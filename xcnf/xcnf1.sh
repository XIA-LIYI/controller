#!/bin/sh
#SBATCH --time=10
#SBATCH --nodelist=xcnf1
#SBATCH --ntasks=1 --cpus-per-task=20

srun ./client_xcnf1