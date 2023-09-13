#!/bin/sh
#SBATCH --time=10
#SBATCH --nodelist=xcnf6
#SBATCH --ports=40000-41000

srun ./client_xcnf6