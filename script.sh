#!/bin/bash

xcne=(2 3 4 5 6 7)
xgph=(5 6 7 8 9 10 11 12 13 14 15 16 17 18 19)
xgpg=(3 4 5 6 7 8 9)

for i in "${xcne[@]}"  
do
   sbatch ./xcne/xcne$i.sh
done

for i in "${xgpg[@]}"  
do
   sbatch ./xgpg/xgpg$i.sh
done

for i in "${xgph[@]}"  
do
   sbatch ./xgph/xgph$i.sh
done

# for i in "${xgph[@]}"  
# do
#     echo $i
# done

