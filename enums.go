package bittrex

type OrderType string

const (
	MARKET OrderType = "MARKET"
	LIMIT OrderType = "LIMIT"
	CEILING_LIMIT OrderType = "CEILING_LIMIT"
	CEILING_MARKET OrderType = "CEILING_MARKET"
)

type TimeInForce string

const (
	GOOD_TIL_CANCELLED TimeInForce = "GOOD_TIL_CANCELLED"
	IMMEDIATE_OR_CANCEL TimeInForce = "IMMEDIATE_OR_CANCEL"
	FILL_OR_KILL TimeInForce = "FILL_OR_KILL"
	POST_ONLY_GOOD_TIL_CANCELLED TimeInForce = "POST_ONLY_GOOD_TIL_CANCELLED"
	BUY_NOW TimeInForce = "BUY_NOW"
)

type OrderDirection string

const (
	BUY OrderDirection = "BUY"
	SELL OrderDirection = "SELL"
)

type WithdrawalStatus string

const (
	REQUESTED WithdrawalStatus = "REQUESTED"
	AUTHORIZED WithdrawalStatus = "AUTHORIZED"
	PENDING WithdrawalStatus = "PENDING"
	ERROR_INVALID_ADDRESS WithdrawalStatus = "ERROR_INVALID_ADDRESS"
	ALL WithdrawalStatus = ""
)