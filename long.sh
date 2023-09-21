#!/bin/sh
#SBATCH --time=20
#SBATCH --partition=long
#SBATCH --nodes=16
#SBATCH --ntasks=16 --cpus-per-task=1
#SBATCH --ntasks-per-node=1
#SBATCH --nodelist=xcne4,xcne5,xcne6,xcne7,xgpd6,xgpd7,xgpd9,xgpe3,xgpe4,xgpe5,xgpe6,xgpe7,xgpe8,xgpe9,xgph7,xgph17
srun -n 16 ./client
# srun -N 2 -n 2 ./client
