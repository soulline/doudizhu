package computer

import (
	"doudizhu/enum"
	"doudizhu/model"
	"doudizhu/util"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

/**
构造一副牌
*/
func CreateNew() []string {
	numbers := make([]string, 54)
	start := 0
	for i := 3; i <= 16; i++ {
		fmt.Println(start)
		if i == 16 {
			fmt.Println("end start : " + strconv.Itoa(start))
			numbers[start] = "Q88"
			numbers[start+1] = "K99"
		} else {
			numbers[start] = "A" + strconv.Itoa(i)
			numbers[start+1] = "B" + strconv.Itoa(i)
			numbers[start+2] = "C" + strconv.Itoa(i)
			numbers[start+3] = "D" + strconv.Itoa(i)
			start += 4
		}
	}
	return numbers
}

/**
洗牌
*/
func Shuffle(vals []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

/**
*发牌
*order==0 玩家1次序
*order==1 玩家2次序
*order==2 玩家3次序
*order==3 底牌次序
 */
func Dispacther(order int, vals []string) []string {
	var playCards []string
	if order < 0 || order > 3 {
		return []string{}
	} else {
		size := 17
		if order == 3 {
			size = 3
		}
		for i := 0; i < size; i++ {
			playCards = append(playCards, vals[order*17+i])
		}
	}
	return playCards
}

/**
* 根据牌面数量判断牌面类型
 */
func ParseCardsInSize(plays []string) model.CardShow {
	cardShow := model.CardShow{
		ShowValue: plays,
		ShowTime:  util.GetNowTime(),
	}
	switch len(plays) {
	case 1:
		cardShow.CardTypeStatus = enum.SINGLE
		cardShow.CompareValue = GetCardValue(plays[0])
		cardShow.MaxCount = 1
		cardShow.MaxValues = []int{cardShow.CompareValue}
		fmt.Printf("根%d", GetCardValue(plays[0]))
		break
	case 2:
		if plays[0] == "Q88" && plays[1] == "K99" {
			cardShow.CardTypeStatus = enum.KING_BOMB
			cardShow.CompareValue = GetCardValue(plays[0])
			cardShow.MaxCount = 2
			cardShow.MaxValues = []int{cardShow.CompareValue}
			fmt.Println("王炸")
		} else {
			ParseCardsType(plays, &cardShow)
		}
		break
	}
	if len(plays) > 2 {
		ParseCardsType(plays, &cardShow)
	} else {
		cardShow.CardTypeStatus = enum.ERROR_TYPE
	}
	return cardShow
}

/**
* 获取牌面类型
 */
func ParseCardsType(cards []string, cardShow *model.CardShow) {
	mapCard, maxCount, maxValues := ComputerValueTimes(cards)
	cardShow.MaxCount = maxCount
	cardShow.MaxValues = maxValues
	cardShow.CardMap = mapCard
	cardShow.CompareValue = maxValues[len(maxValues)-1]
	switch maxCount {
	case 4:
		if maxCount == len(cards) {
			cardShow.CardTypeStatus = enum.KING_BOMB
			fmt.Println("炸弹")
		} else if len(cards) == 6 {
			cardShow.CardTypeStatus = enum.FOUR_TWO
			fmt.Println("四带二")
		} else {
			cardShow.CardTypeStatus = enum.ERROR_TYPE
			fmt.Println("不合法出牌")
		}
		break
	case 3:
		alive := len(cards) - len(maxValues)*maxCount
		if len(maxValues) == alive {
			if len(maxValues) == 1 {
				cardShow.CardTypeStatus = enum.THREE_AND_ONE
				fmt.Println("三带一")
			} else if len(maxValues) > 1 {
				if IsContinuity(mapCard, false) {
					cardShow.CardTypeStatus = enum.PLANE
					fmt.Printf("飞机%d", len(maxValues))
				} else {
					cardShow.CardTypeStatus = enum.ERROR_TYPE
					fmt.Println("非法飞机")
				}
			}
		} else if alive == 0 {
			if len(maxValues) > 1 {
				if IsContinuity(mapCard, false) {
					cardShow.CardTypeStatus = enum.PLANE_EMPTY
					fmt.Printf("三不带飞机%d", len(maxValues))
				} else {
					cardShow.CardTypeStatus = enum.ERROR_TYPE
					fmt.Println("非法三不带飞机")
				}

			} else {
				cardShow.CardTypeStatus = enum.THREE
				fmt.Println("三不带")
			}
		} else {
			cardShow.CardTypeStatus = enum.ERROR_TYPE
			fmt.Println("不合法飞机或三带一")
		}
		break
	case 2:
		if len(maxValues) == (len(cards) / 2) {
			if len(maxValues) > 1 {
				if IsContinuity(mapCard, false) && len(maxValues) > 2 {
					cardShow.CardTypeStatus = enum.DOUBLE_ALONE
					fmt.Printf("%d连队", len(maxValues))
				} else {
					cardShow.CardTypeStatus = enum.ERROR_TYPE
					fmt.Println("非法连对")
				}
			} else if len(maxValues) == 1 {
				cardShow.CardTypeStatus = enum.DOUBLE
				fmt.Printf("对%d", GetCardValue(cards[0]))
			}
		} else {
			cardShow.CardTypeStatus = enum.ERROR_TYPE
			fmt.Println("不合法出牌")
		}
		break
	case 1:
		if IsContinuity(mapCard, true) && len(cards) >= 5 {
			cardShow.CardTypeStatus = enum.SINGLE_ALONE
			fmt.Printf("%d顺子", len(mapCard))
		} else {
			cardShow.CardTypeStatus = enum.ERROR_TYPE
			fmt.Println("非法顺子")
		}
		break
	}
}

/**
* 获取顺序的key值数组
 */
func GetOrderKeys(cardMap map[int]int, isSingle bool) []int {
	var keys []int
	for key, value := range cardMap {
		if (!isSingle && value > 1) || isSingle {
			keys = append(keys, key)
		}
	}
	sort.Ints(keys)
	return keys
}

/**
* 计算牌面值是否连续
 */
func IsContinuity(cardMap map[int]int, isSingle bool) bool {
	keys := GetOrderKeys(cardMap, isSingle)
	lastKey := 0
	for i := 0; i < len(keys); i++ {
		if (lastKey > 0 && (keys[i]-lastKey) != 1) || keys[i] == 15 {
			return false
		}
		lastKey = keys[i]
	}
	if lastKey > 0 {
		return true
	} else {
		return false
	}
}

/**
* 计算每张牌面出现的次数
* mapCard 标记结果
* MaxCount 出现最多的次数
* MaxValues 出现次数最多的所有值
 */
func ComputerValueTimes(cards []string) (mapCard map[int]int, MaxCount int, MaxValues []int) {
	newMap := make(map[int]int)
	if len(cards) == 0 {
		return newMap, 0, nil
	}
	for _, value := range cards {
		cardValue := GetCardValue(value)
		if newMap[cardValue] != 0 {
			newMap[cardValue]++
		} else {
			newMap[cardValue] = 1
		}
	}
	var allCount []int //所有的次数
	var maxCount int   //出现最多的次数
	for _, value := range newMap {
		allCount = append(allCount, value)
	}
	maxCount = allCount[0]
	for i := 0; i < len(allCount); i++ {
		if maxCount < allCount[i] {
			maxCount = allCount[i]
		}
	}
	var maxValue []int
	for key, value := range newMap {
		if value == maxCount {
			maxValue = append(maxValue, key)
		}
	}
	sort.Ints(maxValue)
	return newMap, maxCount, maxValue
}

/**
* 获取牌面值
 */
func GetCardValue(card string) int {
	stringValue := util.Substring(card, 1, len(card))
	value, err := strconv.Atoi(stringValue)
	if err == nil {
		return value
	}
	return -1
}
