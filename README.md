# go-kit

## GO泛型工具库

用于重构在业务代码中的冗余、模板代码。该工具库提供了具备缩容的高性能辅助方法和数据结构。

注意: 当前版本并发不安全

## 内容概览

- Slice: （slicex）
  - 辅助方法
      - Insert: 在指定位置插入一个元素
      - Delete: 删除指定位置的元素
      - Find: 查找一个元素
      - Reverse: 返回翻转后的切片
      - ReverseItSelf: 翻转自身
      - UpperBound: 返回有序 slice 中大于 x 的最小值的下标
      - LowerBound: 返回有序 slice 中大于等于 x 的最小值的下标
      - Unique: Slice 去重
- Set: （setx）
  - 集合(NewSet)
    - Insert: 插入元素到集合
    - Size: 返回集合中元素个数
    - Empty: 判断是否为空集
    - Clear: 清空集合
    - Erase: 删除集合中某个元素
    - Exist: 查询元素是否存在
    - Keys: 返回集合所有元素
- Queue: （queuex）
  - 普通队列（NewQueue）
    - Pop: 队头出队
    - Front: 取队头元素
    - Push: 新元素入队
    - Size: 返回队列中元素个数
    - Back: 取队尾元素
    - Empty: 判断队空
    - Clear: 清空队列
  - 堆（NewPriorityQueue）: 大根堆 / 小根堆取决于传入的 compareFunc
    - Pop: 弹出堆顶
    - Top: 取堆顶
    - Size: 堆中元素个数
    - Push: 新元素入堆
    - Empty: 判断堆空
    - Clear: 清空堆
  - 普通双端队列（Deque）
    - PopFront: 弹出队头
    - PopBack: 弹出队尾
    - PushFront: 插入队头
    - PushBack: 插入队尾
    - Front: 取队头
    - Back: 取队尾
    - Size: 返回队列中元素个数
    - Empty: 判断队空
    - Clear: 清空队列
- Stack: (stackx)
  - 普通栈（NewStack）
    - Push: 新元素入栈
    - Top: 取栈顶
    - Pop: 弹出栈顶
    - Size: 返回栈中元素个数
    - Empty: 判断栈空
    - Clear: 清空栈
- 数学: (mathx)
  - GCD: 求 a, b 最大公约数
  - LCM: 求 a, b 最小公倍数
  - IsPrime: 判断素数
  - QMI: 求 a ^ k % p
- Map: (mapx)
  - 辅助方法
    - Keys: 返回 map 中的所有键
    - Values: 返回 map 中的所有值
    - KeysValues: 返回 map 的键和值
    - ToMap: K、V 切片转 map
- 链表: (listx)
  - 双向链表（带头结点）
    - Len: 返回链表长度
    - Values: 返回链表所有节点值构成的切片
    - InsertToHead: 头插法
    - InsertToTail: 尾插法
    - Insert: 插入到指定位置
    - Delete: 删除链表指定位置
    - Sort: 链表排序
    - Reverse: 返回翻转后的链表
    - ReverseItSelf: 翻转链表自身
    - UniqueByOrder: 有序链表去重
    - Unique: 链表去重
  - LRU 缓存机制
    - Get: 返回 key 所对应的值
    - Put: 修改 key 所对应的值为 value，不存在则进行插入
- 树: (treex)
  - 线段树: 
    - 求区间和（未实现）
    - 求区间最大（未实现）
    - 求区间最小（未实现）
    - 求区间异或（未实现）
  - 红黑树: （未实现）
- KIT: (kitx)
  - ToPtr: 转指针
