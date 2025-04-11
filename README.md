# go-kit

## GO泛型工具库

用于重构在业务代码中的冗余、模板代码。该工具库提供了方便的辅助方法和数据结构。

## 数据结构

- Slice：
  - Insert：在指定位置插入一个元素
  - Find：查找一个元素
  - Reverse：返回翻转后的切片
  - ReverseItSelf：翻转自身
  - UpperBound：返回有序 slice 中大于 x 的最小值的下标
  - LowerBound：返回有序 slice 中大于等于 x 的最小值的下标
  - Unique：Slice 去重

- Set：
  - 集合(NewSet)

- Queue：
  - 普通队列（NewQueue）
  - 堆（NewPriorityQueue）：大根堆 / 小根堆取决于传入的 compareFunc
  - 普通双端队列（Dequeue）（未实现）

- Stack：
  - 普通栈（NewStack）

- Tuple:
  - 二元组（Pair)