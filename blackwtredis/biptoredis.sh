#!/bin/sh

dir=/data0/BlackIP/blackwtredis/blocklist-ipsets
whitelist=/data0/BlackIP/blackwtredis/whitelist.txt

redis=127.0.0.1
redis_port=6381
redis_pw=123456
mapfile  whitelistarr < $whitelist


for file in $dir/*; do

if [[ $file == *".ipset"* ]]
then
    echo $file | xargs cat| grep -v ^# | while read i;do

        if [[ ! "${whitelistarr[@]}"  =~ "$i" ]]; then
          redis-cli -h $redis -p $redis_port -a $redis_pw -c  set $i  $(basename ${file} .ipset)
        fi

      done

fi

if [[ $file == *".netset"* ]]
then
    echo $file | xargs cat| grep -v ^# | while read i;do

    if [[ ! "${whitelistarr[@]}"  =~ "$i" ]]; then
       redis-cli -h $redis -p $redis_port -a $redis_pw -c  set $i  $(basename ${file} .netset)
    fi

   done
fi


done