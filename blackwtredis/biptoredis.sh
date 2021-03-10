dir=/opt/blocklist-ipsets
 
for file in $dir/*; do

if [[ $file == *".ipset"* ]]
then
    echo $file | xargs cat| grep -v ^# | while read i;do      
    redis-cli -h 0.0.0.0 -p 6381 -a nBbIBrmoWHaLRJzu -c  set $i  $(basename ${file} .ipset); done
fi

if [[ $file == *".netset"* ]]
then
    echo $file | xargs cat| grep -v ^# | while read i;do  
    redis-cli -h 0.0.0.0 -p 6381 -a nBbIBrmoWHaLRJzu -c  set $i  $(basename ${file} .netset); done
fi


done


