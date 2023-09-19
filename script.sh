#!/bin/bash

xcne=(3 4 5)
xgph=(xgph7.sh)

for i in "${xcne[@]}"  
do
   sbatch ./xcne/xcne$i.sh
done

for i in "${xgpg[@]}"  
do
   sbatch ./xcne/xcne$i.sh
done

for i in "${xgph[@]}"  
do
   sbatch ./xcne/xcne$i.sh
done

# for i in "${xgph[@]}"  
# do
#     echo $i
# done

