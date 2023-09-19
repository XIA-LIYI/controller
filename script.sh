#!/bin/bash

xcne=(xcne2.sh xcne3.sh)
# xcne=(xcne1.sh xcne2.sh xcne3.sh)
# xcne=(xcne1.sh xcne2.sh xcne3.sh)
xgph=(xgph5.sh xgph6.sh xgph7.sh)

for i in "${xcne[@]}"  
do
    srun --partition=long --nodelist=$i --ntasks=1 --cpus-per-task=10 ./client
done

# for i in "${xgph[@]}"  
# do
#     echo $i
# done

