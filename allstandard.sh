#!/bin/sh
#SBATCH --time=20
#SBATCH --partition=standard
#SBATCH --nodes=20
#SBATCH --ntasks=20 --cpus-per-task=1
#SBATCH --ntasks-per-node=1
#SBATCH --nodelist=xgph0,xgph1,xgph2,xgph3,xgph4,xgpg0,xgpg1,xgpg2,xgpe1,xgpe2
srun -n 20 ./client
# srun -N 2 -n 2 ./client
