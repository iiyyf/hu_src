//
//
//    hu.go这个文件用于检测hu，作为public的权限提供给外部的函数条用判胡
//    cards[]代表index->num的键值对,num之和为sum，sum满足sum%3 == 2
//    guiIndex: 代表癞子的索引,CheckHu(cards,MaxCard)计算不带癞子的判胡

package hulib

var GlobalTableMgr *TableMgr

func init() {
	GlobalTableMgr = &TableMgr{}
	GlobalTableMgr.Init()
	GlobalTableMgr.Gen()
}

func CheckHu(cards []int32, guiIndex int32) bool {

	sum := int32(0)
	for _, c := range cards {
		sum = sum + c
	}

	if sum%3 != 2 {
		return false
	}

	if guiIndex >= MaxCard || guiIndex < 0 {
		guiIndex = MaxCard
	}

	tmpCards := [MaxCard]int{}
	for i := 0; i < len(cards) && i < MaxCard; i++ {
		if cards[i] > 0 {
			tmpCards[i] = int(cards[i])
		}
	}
	return getHuInfo(GlobalTableMgr, tmpCards[:], MaxCard, MaxCard, int(guiIndex))
}

//
//func CheckHu(cards []int, guiIndex int) bool {
//
//	sum := 0
//	for _, c := range cards {
//		sum = sum + c
//	}
//
//	if sum%3 != 2 {
//		return false
//	}
//
//	if guiIndex >= MaxCard || guiIndex < 0 {
//		guiIndex = MaxCard
//	}
//
//	tmpCards := [MaxCard]int{}
//	for i := 0; i < len(cards) && i < MaxCard; i++ {
//		if cards[i] > 0 {
//			tmpCards[i] = cards[i]
//		}
//	}
//	return getHuInfo(GlobalTableMgr, tmpCards[:], MaxCard, MaxCard, guiIndex)
//}
