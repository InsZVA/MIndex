# MIndex
## 用法
MIndex致力于将MySQL储存的文章分词做成倒排索引，并存入monkeyDB中，以便搜索引擎使用

> $ ./main -h
Usage of ./main:
>  -database string
    	要倒排的MySQL数据库 (default "CloudKt")

>  -monkey string
    	monkeyDB地址 (default "127.0.0.1")

>  -monkeyP string
    	monkeyDB端口号 (default "1517")

>  -monkeyp string
    	monkeyDB登陆口令 (default "monkey")

>  -mysql string
    	MySQL服务器地址及端口号 (default "127.0.0.1:3306")
>
  -mysqlp string
    	MySQL密码 (default "root")

>  -mysqlu string
    	MySQL用户名 (default "root")

> -table string
    	要倒排的MySQL数据表 (default "courses")

MIndex会将指定的数据表做成倒排索引作为表名相同的数据库在monkeyDB中