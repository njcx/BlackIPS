#!/bin/sh

dir=/opt/blocklist-ipsets
redis=127.0.0.1
redis_port=6381
redis_pw=nBbIBrmoWHaLRJzu
 
for file in $dir/*; do

if [[ $file == *".ipset"* ]]
then
    echo $file | xargs cat| grep -v ^# | while read i;do      
    redis-cli -h $redis -p $redis_port -a $redis_pw -c  set $i  $(basename ${file} .ipset); done
fi

if [[ $file == *".netset"* ]]
then
    echo $file | xargs cat| grep -v ^# | while read i;do  
    redis-cli -h $redis -p $redis_port -a $redis_pw -c  set $i  $(basename ${file} .netset); done
fi


done


