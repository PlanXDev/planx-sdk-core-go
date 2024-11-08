package api

const (
	OrderTypeInvest = "invest"
)

const (
	OrderStatusNew     = "new"
	OrderStatusSuccess = "success"
	OrderStatusFail    = "fail"
)

const (
	PoolStatusActive   = "active"
	PoolStatusInActive = "inactive"
	PoolStatusDestroy  = "destroyed"
)

const (
	PoolCapacityActionIncrease = "increase"
	PoolCapacityActionDecrease = "decrease"
)

const (
	GiftTypeAirdrop    = "airdrop"
	GiftTypeRedemption = "redemption"
)

const (
	GiftStatusWaiting = "waiting"
	GiftStatusIssued  = "issued"
	GiftStatusClaimed = "claimed"
	GiftStatusCancel  = "cancel"
	GiftStatusExpired = "expired"
)

const (
	TradeSideBuy  = "buy"
	TradeSideSell = "sell"
)

const (
	TradeStatusDeal    = "deal"
	TradeStatusPending = "pending"
	TradeStatusClosed  = "closed"
)
