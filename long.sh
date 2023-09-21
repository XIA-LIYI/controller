#!/bin/sh
#SBATCH --time=20
#SBATCH --partition=long
#SBATCH --nodes=17
#SBATCH --ntasks=17 --cpus-per-task=1
#SBATCH --ntasks-per-node=1
#SBATCH --nodelist=xcne4,xcne5,xcne6,xcne7,xgpd9,xgpd7,xgpd9,xgpd6,xgpe3,xgpe4,xgpe5,xgpe6,xgpe7,xgph7,xgph17,xgph18,xgph19
srun -n 17 ./client
# srun -N 2 -n 2 ./client
