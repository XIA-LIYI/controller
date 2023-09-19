#!/bin/bash

xcne=(xcne2.sh xcne3.sh)
# xcne=(xcne1.sh xcne2.sh xcne3.sh)
# xcne=(xcne1.sh xcne2.sh xcne3.sh)
xgph=(xgph5 xgph6 xgph7)

for i in "${xcne[@]}"  
do
   sbatch ./xcne/$i
done

# for i in "${xgph[@]}"  
# do
#     echo $i
# done

