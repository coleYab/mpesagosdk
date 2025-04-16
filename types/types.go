// Package types provides various types and constants used throughout the M-Pesa SDK 
// to interact with the M-Pesa API. These types are used for specifying different command 
// identifiers, transaction types, response types, and identifier types to ensure the correct 
// format and usage of data when making requests or handling responses.
package types

// CommandId represents the unique identifier for a command type to be executed on the M-Pesa API.
//
// The following constants define the available command types:
// 	- CustomerPayBillOnlineCommand: Command for initiating a customer bill payment online.
// 	- AccountBalanceCommand: Command for querying the account balance.
// 	- CustomerBuyGoodsOnlineCommand: Command for initiating a customer purchase online.
// 	- BusinessPaymentCommand: Command for making a business payment.
// 	- SalaryPaymentCommand: Command for making a salary payment.
// 	- PromotionPaymentCommand: Command for making a promotional payment.
// 	- RegisterURLCommand: Command for registering a callback URL for C2B transactions.
// 	- TransactionStatusCommand: Command for querying the status of a transaction.
// 	- TransactionReversalCommand: Command for reversing a transaction.
//	 
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

// IdentifierType represents the type of identifier used to specify a party in a 
// transaction request, such as a customer or business identifier.
//
// The following constants define the available identifier types:
// 	- MsisdnIdentifierType: Represents the MSISDN (Mobile Subscriber Integrated Services Digital Network) number.
// 	- TillNumberIdentifierType: Represents the till number used in payments.
// 	- ShortCodeIdentifierType: Represents the short code used in payments.
type IdentifierType string

const (
	MsisdnIdentifierType     IdentifierType = "1"
	TillNumberIdentifierType IdentifierType = "2"
	ShortCodeIdentifierType  IdentifierType = "4"
)

// ResponseType represents the status of a transaction or request after it has been processed by M-Pesa 
// when registering callback for c2b payments.
//
// The following constants define the available response types:
// 	- CompletedResponse: Indicates that the url has to be requested when the transaction was successfully completed.
// 	- CancelledResponse: Indicates that the url has to be requested when the transaction was cancelled.
type ResponseType string

const (
	CompletedResponse ResponseType = "Completed"
	CancelledResponse ResponseType = "Cancelled"
)

// TransactionType represents the type of transaction being processed or initiated. It helps 
// to specify the purpose of the transaction, such as paying a bill or purchasing goods.
//
// The following constants define the available transaction types:
// 	- CustomerPayBillOnlineTransaction: Transaction type for customer bill payments online.
// 	- CustomerBuyGoodsOnlineTransaction: Transaction type for customer purchases online.
type TransactionType string

const (
	CustomerPayBillOnlineTransaction  TransactionType = "CustomerPayBillOnline"
	CustomerBuyGoodsOnlineTransaction TransactionType = "CustomerBuyGoodsOnline"
)
