基于Golang的分布式爬虫
目前有三个分支：
##master:
一个完备的单机爬虫, have been  concurrent branch to master.

###concurrent:
并发版

###distribute:
基于jsonrpc的分布式版，分为三个模块

* engine
  负责整个爬虫的调度
  
* worker
  多个goroutine，负责并发爬取网页
 
* persist
  负责讲解析后的数据持久化保存
