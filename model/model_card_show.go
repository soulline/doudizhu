package model

import (
	"doudizhu/enum"
	"time"
)

type CardShow struct {
	Id             int //出牌Id
	ShowTime       time.Time
	ShowValue      []string            //牌面数组
	CardMap        map[int]int         //牌面计算结果
	MaxCount       int                 //同值牌出现的最大次数
	MaxValues      []int               //同值牌出现的最大次数列表
	CompareValue   int                 //用于比较大小的值
	CardTypeStatus enum.CardTypeStatus //牌面类型
}
