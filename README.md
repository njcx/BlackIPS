# blackips

开源威胁情报，包含3个组件，300万+恶意IP。


#### --blackapi

Go语言开发，威胁情报查询API

#### --blackip

Go语言开发，威胁情报查询前端


#### --iplocation

PHP语言开发，IP地理位置查询API，包含Nginx配置


#### --blackwtredis

读取威胁情报，写入Redis Shell的脚本，以及Crontab


####  使用说明

本项目威胁情报会定期自动更新，威胁情报存储到Redis，请自行安装Redis并设置密码

1，修改biptoredis.sh 里面的Redis 配置，把biptoredis.sh 放到/opt 目录下，并把Crontab记录添加上

2, 修改 blackapi 的Redis配置，编译启动

3，搭建nginx +php5.x +php-fpm环境，启动iplocation 服务

4, 修改blackip里面的api配置，编译启动即可

5， 第一次启动，先手动执行一下Crontab里面的命令，完成第一次添加，后续即自动更新



![Image](https://raw.githubusercontent.com/njcx/blackips/master/img.jpg)