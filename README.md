#### 牌面协议
A：黑桃 B：红桃 C：梅花 D：方片
|扑克原始值|映射值|
|------|-----|
|3-10|3-10数字|
|J|11|
|Q|12|
|K|13|
|A|14|
|2|15|
|小王|Q88|
|大王|K99|

#### 构造一副牌
```
computer.CreateNew()
```
#### 洗牌
```
initValues := card.CreateNew()
computer.Shuffer(initValues)
```
#### 按序发牌
```
initValues := card.CreateNew()
computer.Shuffle(initValues)
fmt.Println("玩家1：", computer.Dispacther(0, initValues))
fmt.Println("玩家2：", computer.Dispacther(1, initValues))
fmt.Println("玩家3：", computer.Dispacther(2, initValues))
fmt.Println("底牌：", computer.Dispacther(3, initValues))
```

#### 出牌类型枚举
```
type CardTypeStatus int

const (
    _CardTypeStatus = iota
    SINGLE          //单根
    DOUBLE          //对子
    THREE           //三不带
    THREE_AND_ONE   //三带一
    BOMB            //炸弹
    FOUR_TWO        //四带二
    PLANE           //飞机
    PLANE_EMPTY     //三不带飞机
    DOUBLE_ALONE    //连对
    SINGLE_ALONE    //顺子
    KING_BOMB       //王炸
    ERROR_TYPE      //非法类型

)
```
#### 计算模型
```
type CardShow struct {
    ShowValue      []string            //牌面数组
    CardMap        map[int]int         //牌面计算结果
    MaxCount       int                 //同值牌出现的最大次数
    MaxValues      []int               //同值牌出现的次数列表
    CompareValue   int                 //用于比较大小的值
    CardTypeStatus enum.CardTypeStatus //牌面类型
}

```
