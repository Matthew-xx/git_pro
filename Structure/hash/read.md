哈希表可以实现快速查找，但不适用频繁插入删除

1、将数据分成100个文件，开启100个线程搜索xxxx

    :每个线程的搜索结果都放在一个线程安全的（优先）队列
    
    ：为用户名为yy的优先级设置2，用户名包含yy的优先级设置为1，密码包含yy的设置优先级为0
    
