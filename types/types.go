package types

type CommandId string

const (
	CustomerPayBillOnlineCommand  CommandId = "CustomerPayBillOnline"
	AccountBalanceCommand         CommandId = "AccountBalance"
	CustomerBuyGoodsOnlineCommand CommandId = "CustomerBuyGoodsOnline"
	BusinessPaymentCommand        CommandId = "BusinessPayment"
	SalaryPaymentCommand          CommandId = "SalaryPayment"
	PromotionPaymentCommand       CommandId = "PromotionPayment"
	RegisterURLCommand            CommandId = "RegisterURL"
	TransactionStatusCommand      CommandId = "TransactionStatusQuery"
	TransactionReversalCommand    CommandId = "TransactionReversal"
)

type IdentifierType string

const (
	MsisdnIdentifierType     IdentifierType = "1"
	TillNumberIdentifierType IdentifierType = "2"
	ShortCodeIdentifierType  IdentifierType = "4"
)

type ResponseType string

const (
	CompletedResponse ResponseType = "Completed"
	CancelledResponse ResponseType = "Cancelled"
)

type TransactionType string

const (
	CustomerPayBillOnlineTransaction  TransactionType = "CustomerPayBillOnline"
	CustomerBuyGoodsOnlineTransaction TransactionType = "CustomerBuyGoodsOnline"
)
