package game
const(
	ROYAL_FLUSH    int= 10  // 皇家同花顺：同一花色的最大顺子。（最大牌：A-K-Q-J-10）
	STRAIGHT_FLUSH int= 9  // 同花顺：同一花色的顺子。（最大牌：K-Q-J-10-9 最小牌：A-2-3-4-5）
	FOUR           int= 8 // 四条：四同张加单张。（最大牌：A-A-A-A-K 最小牌：2-2-2-2-3）
	FULL_HOUSE     int= 7  // 葫芦（豪斯）：三同张加对子。（最大牌：A-A-A-K-K 最小牌：2-2-2-3-3）
	FLUSH          int= 6  // 同花：同一花色。（最大牌：A-K-Q-J-9 最小牌：2-3-4-5-7）
	STRAIGHT       int= 5 // 顺子：花色不一样的顺子。（最大牌：A-K-Q-J-10 最小牌：A-2-3-4-5）
	THREE          int= 4  // 三条：三同张加两单张。（最大牌：A-A-A-K-Q 最小牌：2-2-2-3-4）
	TWO_PAIR       int= 3  // 两对：（最大牌：A-A-K-K-Q 最小牌：2-2-3-3-4）
	ONE_PAIR       int= 2  // 一对：（最大牌：A-A-K-Q-J 最小牌：2-2-3-4-5）
	HIGH_CARD      int= 1 // 高牌：（最大牌：A-K-Q-J-9 最小牌：2-3-4-5-7）
)

func GetResult(result int) string{
	switch result{
	case ROYAL_FLUSH:
		return "ROYAL_FLUSH"
	case STRAIGHT_FLUSH:
		return "STRAIGHT_FLUSH"
	case FOUR:
		return "FOUR"
	case FULL_HOUSE:
		return "FULL_HOUSE"
	case FLUSH:
		return "FLUSH"
	case STRAIGHT:
		return "STRAIGHT"
	case THREE:
		return "THREE"
	case TWO_PAIR:
		return "TWO_PAIR"
	case ONE_PAIR:
		return "ONE_PAIR"
	case HIGH_CARD:
		return "HIGH_CARD"
	default:
		return "HIGH_CARD"
	}
}