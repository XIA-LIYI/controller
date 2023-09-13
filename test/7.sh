#!/bin/sh
#SBATCH --time=10
#SBATCH --nodelist=xcnf7
#SBATCH --ntasks=1 --cpus-per-task=20

srun ./test_7