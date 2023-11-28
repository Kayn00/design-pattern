来源 https://github.com/silsuer/golang-design-patterns/blob/master/builder-pattern/README.md?plain=1
### 概述

> wiki: **建造者模式(Builder Pattern)**：将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

直白一点的说，就是将我们在开发过程中遇到的大型对象，拆分成多个小对象，然后将多个小对象组装成大对象，并且对外部隐藏建造过程.

### 结构

建造者模式由一下4个部分组成

- `Builder`: 抽象建造者

- `ConcreteBuilder`: 具体建造者

- `Director`: 指挥者

- `Production`: 产品(大型产品以及拆分成的小型产品)

### 类图 && 时序图

![](https://design-patterns.readthedocs.io/zh_CN/latest/_images/Builder.jpg)
![](https://design-patterns.readthedocs.io/zh_CN/latest/_images/seq_Builder.jpg)
-----------
(*图源网络*)

从上面两张图可以看出建造者模式的使用流程:

  1. 创建大型产品建造者
  2. 创建指挥者
  3. 将建造者传入指挥者对象中
  4. 由指挥者指挥建造者创建对象，并返回
