1、内存连续存储（数组

    查找修改的复杂度：O(1)
    
    删除、插入：O(n)
    
    计算机理论上不存在大片连续内存（查找速度快的优势也不再
    
2、链表（小块分布存储

   删除、插入：O(1)
   
   查找、修改：O(n) 
   
3、文件遍历：

   轻量级： 使用数组栈 —— 深度遍历。使用数组队列 —— 广度遍历
   
   重量级： 使用链表栈 —— 深度遍历，链表队列 —— 广度遍历 
    