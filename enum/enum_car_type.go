package enum

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
