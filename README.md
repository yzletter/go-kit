# go-kit

## GO泛型工具库

用于重构在业务代码中的冗余、模板代码。该工具库提供了具备缩容的高性能辅助方法和数据结构。

注意：当前版本并发不安全

## 数据结构

- Slice：
  - 辅助方法
      - Insert：在指定位置插入一个元素
      - Delete：删除指定位置的元素
      - Find：查找一个元素
      - Reverse：返回翻转后的切片
      - ReverseItSelf：翻转自身
      - UpperBound：返回有序 slice 中大于 x 的最小值的下标
      - LowerBound：返回有序 slice 中大于等于 x 的最小值的下标
      - Unique：Slice 去重
- Set：
  - 集合(NewSet)
    - Insert：插入元素到集合
    - Size：返回集合中元素个数
    - Empty：判断是否为空集
    - Clear：清空集合
    - Erase：删除集合中某个元素
    - Exist：查询元素是否存在
    - Keys：返回集合所有元素
- Queue：
  - 普通队列（NewQueue）
    - Pop：队头出队
    - Front：取队头元素
    - Push：新元素入队
    - Size：返回队列中元素个数
    - Back：取队尾元素
    - Empty：判断队空
    - Clear：清空队列
  - 堆（NewPriorityQueue）：大根堆 / 小根堆取决于传入的 compareFunc
    - Pop：弹出堆顶
    - Top：取堆顶
    - Size：堆中元素个数
    - Push：新元素入堆
    - Empty：判断堆空
    - Clear：清空堆
  - 普通双端队列（Deque）
    - PopFront：弹出队头
    - PopBack：弹出队尾
    - PushFront：插入队头
    - PushBack：插入队尾
    - Front：取队头
    - Back：取队尾
    - Size：返回队列中元素个数
    - Empty：判断队空
    - Clear：清空队列
- Stack：
  - 普通栈（NewStack）
    - Push：新元素入栈
    - Top：取栈顶
    - Pop：弹出栈顶
    - Size：返回栈中元素个数
    - Empty：判断栈空
    - Clear：清空栈