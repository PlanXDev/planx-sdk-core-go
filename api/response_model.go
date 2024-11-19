package api

type ResponseBalancesBase struct {
	Symbol    string `json:"symbol"`    //The token name of the account
	Available string `json:"available"` //The available balance of the account
	Frozen    string `json:"frozen"`    //The frozen balance of the account
}

type ResponseBalancesPool struct {
	PoolID       string `json:"poolId"`       //The unique identifier of the capital pool
	SourceSymbol string `json:"sourceSymbol"` //Token name on the left
	TargetSymbol string `json:"targetSymbol"` //Token name on the right
	SourceAmount string `json:"sourceAmount"` //The amount of tokens on the left
	TargetAmount string `json:"targetAmount"` //The amount of tokens on the right
}

type ResponseBalancesMarket struct {
	SourceSymbol string `json:"sourceSymbol"` //Token name on the left
	TargetSymbol string `json:"targetSymbol"` //Token name on the right
	SourceAmount string `json:"sourceAmount"` //The amount of tokens on the left
	TargetAmount string `json:"targetAmount"` //The amount of tokens on the right
}

type ResponseNewOrder struct {
	UserName        string `json:"userName"`        //Account ID
	OrderID         string `json:"orderId"`         //Order ID
	ExternalOrderID string `json:"externalOrderId"` //The operation identifier passed in through parameters
	Symbol          string `json:"symbol"`          //Target Token Name
	OrderType       string `json:"orderType"`       //Order type, only returns model.OrderTypeInvest
	OriginalAmount  string `json:"originalAmount"`  //Target token actual usage amount
	Amount          string `json:"amount"`          // Target token usage amount
	OrderStatus     string `json:"orderStatus"`     //Order status, only returns model.OrderStatusSuccess
	TimeCreated     int64  `json:"timeCreated"`     //The 10-digit second-level timestamp of this operation
}

type ResponseGiftSourceInfo struct {
	PoolID        string `json:"poolId"`        //The unique identifier of the capital pool
	UserName      string `json:"userName"`      //Account ID
	GiftSourceId  string `json:"giftSourceId"`  //The unique identifier of the gift source
	QrCode        string `json:"qrcode"`        //The unique CODE of the gift
	GiftLink      string `json:"giftLink"`      //GIFT pickup address
	GiftClaimLink string `json:"giftClaimLink"` //GIFT pickup address
	GiftName      string `json:"giftName"`      //Gift name
	GiftType      string `json:"giftType"`      //Gift type model.GiftTypeAirdrop,model.GiftTypeRedemption
	Quantity      int64  `json:"quantity"`      //Generate quantity
	PriceAmount   string `json:"priceAmount"`   //Number of tokens for a single Gift
	SourceSymbol  string `json:"sourceSymbol"`  //Token name on the left
	SourceAmount  string `json:"sourceAmount"`  //The amount of tokens on the left
	TargetSymbol  string `json:"targetSymbol"`  //Token name on the right
	TargetAmount  string `json:"targetAmount"`  //The amount of tokens on the right
	ServiceFee    string `json:"serviceFee"`    //Service fee
}

type ResponseGift struct {
	GiftID        string `json:"giftId"`        //The unique identifier of the gift
	PoolID        string `json:"poolId"`        //The unique identifier of the capital pool
	UserName      string `json:"userName"`      //Account ID
	SourceSymbol  string `json:"sourceSymbol"`  // Token name on the left
	SourceAmount  string `json:"sourceAmount"`  // The amount of tokens on the left
	TargetSymbol  string `json:"targetSymbol"`  // Token name on the right
	TargetAmount  string `json:"targetAmount"`  // The amount of tokens on the right
	ClaimSymbol   string `json:"claimSymbol"`   //Claim token name
	ClaimAmount   string `json:"claimAmount"`   //Number of tokens received
	ClaimUserName string `json:"claimUserName"` //Receive user ID
	ClaimTime     int64  `json:"claimTime"`     //Receive time
	QRCode        string `json:"qrcode"`        //The unique CODE of the gift
	GiftLink      string `json:"giftLink"`      //GIFT pickup address
	GiftStatus    string `json:"giftStatus"`    //Gift status,model.GiftStatusWaiting,model.GiftStatusIssued,model.GiftStatusClaimed,model.GiftStatusCancel,model.GiftStatusExpired
	ExpiresTime   int64  `json:"expiresTime"`   //The 10-digit second-level timestamp of this operation
}

type ResponseGiftSourceDetail struct {
	PoolID          string `json:"poolId"`          //The unique identifier of the capital pool
	UserName        string `json:"userName"`        //Account ID
	GiftSourceId    string `json:"giftSourceId"`    //The unique identifier of the gift source
	QrCode          string `json:"qrcode"`          //The unique CODE of the gift
	GiftLink        string `json:"giftLink"`        //GIFT pickup address
	GiftClaimLink   string `json:"giftClaimLink"`   //GIFT pickup address
	GiftName        string `json:"giftName"`        //Gift name
	Quantity        int64  `json:"quantity"`        //Generate quantity
	IssuedQuantity  int64  `json:"issuedQuantity"`  //Quantity issued
	ReceiveQuantity int64  `json:"receiveQuantity"` // The number of received
	PriceAmount     string `json:"priceAmount"`     //Number of tokens for a single Gift
	GiftStatus      string `json:"giftStatus"`      //Gift status,model.GiftStatusWaiting,model.GiftStatusIssued,model.GiftStatusClaimed,model.GiftStatusCancel,model.GiftStatusExpired
	ExpiresTime     int64  `json:"expiresTime"`     //The 10-digit second-level timestamp of this operation
}

type ResponseGiftSourceExpire struct {
	PoolID          string `json:"poolId"`          //The unique identifier of the capital pool
	QRCode          string `json:"qrcode"`          //The unique CODE of the gift
	GiftSourceId    string `json:"giftSourceId"`    //The unique identifier of the gift source
	GiftStatus      string `json:"giftStatus"`      //Gift status,model.GiftStatusExpired
	ExpiresTime     int64  `json:"expiresTime"`     //The 10-digit second-level timestamp of this operation
	Quantity        int64  `json:"quantity"`        //Generate quantity
	ReceiveQuantity int64  `json:"receiveQuantity"` //The number of received
	ExpiresQuantity int64  `json:"expiresQuantity"` //The number of expires
}

type ResponseIssuedGift struct {
	GiftID       string `json:"giftId"`       //The unique identifier of the gift
	PoolID       string `json:"poolId"`       //The unique identifier of the capital pool
	UserName     string `json:"userName"`     //Account ID
	QRCode       string `json:"qrcode"`       //The unique CODE of the gift
	GiftLink     string `json:"giftLink"`     //GIFT pickup address
	GiftStatus   string `json:"giftStatus"`   //Gift status,model.GiftStatusIssued
	ExpiresTime  int64  `json:"expiresTime"`  //The 10-digit second-level timestamp of this operation
	SourceSymbol string `json:"sourceSymbol"` //Token name on the left
	SourceAmount string `json:"sourceAmount"` //The amount of tokens on the left
	TargetSymbol string `json:"targetSymbol"` //Token name on the right
	TargetAmount string `json:"targetAmount"` //The amount of tokens on the right
}

type ResponsePool struct {
	PoolID               string `json:"poolId"`               //The unique identifier of the capital pool
	UserName             string `json:"userName"`             //Account ID
	SourceSymbol         string `json:"sourceSymbol"`         //Token name on the left
	TargetSymbol         string `json:"targetSymbol"`         //Token name on the right
	SourceOriginalAmount string `json:"sourceOriginalAmount"` //The original amount of tokens on the left
	TargetOriginalAmount string `json:"targetOriginalAmount"` //The original amount of tokens on the right
	SourceAmount         string `json:"sourceAmount"`         //The amount of tokens on the left
	TargetAmount         string `json:"targetAmount"`         //The amount of tokens on the right
	SourceDealAmount     string `json:"sourceDealAmount"`     //The amount of tokens traded on the left
	TargetDealAmount     string `json:"targetDealAmount"`     //The amount of tokens traded on the right
	PoolStatus           string `json:"poolStatus"`           //Pool status,model.PoolStatusActive,model.PoolStatusInActive,model.PoolStatusDestroy
	CreateTime           int64  `json:"createTime"`           //The 10-digit second-level timestamp of this operation
}

type ResponseTradeToken struct {
	TradeID              string `json:"tradeId"`              //The unique identifier of the trade
	SourceSymbol         string `json:"sourceSymbol"`         //Token name on the left
	TargetSymbol         string `json:"targetSymbol"`         //Token name on the right
	UserName             string `json:"userName"`             //Account ID
	CounterPartyUserName string `json:"counterPartyUserName"` //Counterparty user of trade
	Amount               string `json:"amount"`               //Number of trading tokens
	Price                string `json:"price"`                //Trading token unit price
	Side                 string `json:"side"`                 //Trade side,model.TradeSideBuy,model.TradeSideSell
	TradeStatus          string `json:"tradeStatus"`          //Trade status, model.TradeStatusPending,model.TradeStatusDeal,model.TradeStatusClosed
	VoucherID            string `json:"voucherId"`            //
	CreateTime           int64  `json:"createTime"`           //The 10-digit second-level timestamp of this operation
}
