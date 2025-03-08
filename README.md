# `mpesa_go_sdk`: The Best M-Pesa SDK for Golang

![coverage](https://img.shields.io/badge/coverage-100.0%25-brightgreen)
[![Go Reference](https://pkg.go.dev/badge/github.com/coleYab/mpesasdk.svg)](https://pkg.go.dev/github.com/coleYab/mpesasdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/coleYab/mpesasdk)](https://goreportcard.com/report/github.com/coleYab/mpesasdk)

`mpesasdk` is a Go SDK for interacting with Safaricom's M-Pesa API. This SDK simplifies integration with M-Pesa's services, enabling operations such as B2C payments, C2B URL registration, USSD Push payments, transaction status queries, account balance checks, and transaction reversals.

## Features

- **B2C Payments**: Transfer funds from a business account to a customer account.
- **C2B URL Registration**: Register URLs for payment notifications.
- **USSD Push**: Initiate USSD-based payment requests.
- **Transaction Status**: Query the status of transactions.
- **Account Balance**: Retrieve M-Pesa account balances.
- **Transaction Reversal**: Reverse a completed M-Pesa transaction.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Examples](#examples)
  - [Register C2B URL](#register-c2b-url)
  - [Simulate C2B Payment](#simulate-c2b-payment)
  - [Make B2C Payment](#make-b2c-payment)
  - [Transaction Status Query](#transaction-status-query)
  - [Account Balance Query](#account-balance-query)
  - [Transaction Reversal](#transaction-reversal)
  - [USSD Push Payment](#stk-push-payment)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install the SDK, use the following command:

```bash
go get github.com/coleYab/mpesasdk
```

## Quick Start

### Initialize the MpesaClient

```go
package main

import (
	"fmt"
	"log"

	"github.com/coleYab/mpesasdk"
	"github.com/coleYab/mpesasdk/config"
)

func main() {
	// Load configuration from environment variables
	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create a new Mpesa client
	app := mpesasdk.New(cfg)
	fmt.Println("Application is created with: ", *app)
}
```

## Examples

### Register C2B URL

```go
// Register C2B URL Example
res, err := app.RegisterNewURL(
    c2b.RegisterC2BURLRequest{
        ShortCode:       "802000",
        ResponseType:    "Completed",
        CommandID:       "RegisterURL",
        ConfirmationURL: "https://www.myservice:8080/confirmation",
        ValidationURL:   "https://www.myservice:8080/validation",
    })
if err != nil {
    log.Printf("failed to register C2B URL: %v\n", err)
}
fmt.Println("C2B URL Registration Response: ", res)
```

### Simulate C2B Payment

```go
// Simulate C2B Payment Example
res, err := app.SimulateCustomerInitiatedPayment(c2b.SimulateCustomerInititatedPayment{
    CommandID:     "CustomerPayBillOnline",
    Amount:        110,
    Msisdn:        "251945628580",
    BillRefNumber: "091091",
    ShortCode:     "443443",
})
if err != nil {
    log.Printf("failed to simulate C2B payment: %v\n", err)
}
fmt.Println("C2B Payment Simulation Response: ", res)
```

### Make B2C Payment

```go
id := uuid.New().String()

// B2C Payment Example
response, err := app.MakeB2CPaymentRequest(b2c.B2CRequest{
    InitiatorName:            "apiuser",
    SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
    OriginatorConversationID: id,
    CommandID:                types.BusinessPaymentCommand,
    Occasion:                 "Occasion",
    Amount:                   1030,
    PartyA:                   600000,
    PartyB:                   251700404709,
    Remarks:                  "Salary Payment",
    QueueTimeOutURL:          "https://yourdomain.com/timeout",
    ResultURL:                "https://yourdomain.com/result",
})
if err != nil {
    log.Fatalf("failed to make b2c payment: %v", err)
}
fmt.Println("B2C Payment Response: ", response)
```

### Transaction Status Query

```go
// Transaction Status Query Example
res, err := app.MakeTransactionStatusQuery(transaction.TransactionStatusRequest{
    Initiator:                "apitest",
    SecurityCredential:       "lMhf0UqE4ydeEDwpUskmPgkNDZnA6NLi7z3T1TQuWCkH3/ScW8pRRnobq/AcwFvbC961+zDMgOEYGm8Oivb7L/7Y9ED3lhR7pJvnH8B1wYis5ifdeeWI6XE2NSq8X1Tc7QB9Dg8SlPEud3tgloB2DlT+JIv3ebIl/J/8ihGVrq499bt1pz/EA2nzkCtGeHRNbEDxkqkEnbioV0OM//0bv4K++XyV6jUFlIIgkDkmcK6aOU8mPBHs2um9aP+Y+nTJaa6uHDudRFg0+3G6gt1zRCPs8AYbts2IebseBGfZKv5K6Lqk9/W8657gEkrDZE8Mi78MVianqHdY/8d6D9KKhw==",
    CommandID:                "TransactionStatusQuery",
    TransactionID:            "0",
    OriginatorConversationID: "AG-20190826-0000777ab7d848b9e721",
    PartyA:                   "1020",
    IdentifierType:           "4",
    ResultURL:                "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
    QueueTimeOutURL:          "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
    Remarks:                  "Trans Status",
    Occasion:                 "Query trans status",
})
if err != nil {
    log.Printf("failed to make transaction status query: %v\n", err)
}
fmt.Println("Transaction Status Query Response: ", res)
```

### Account Balance Query

```go
// Account Balance Query Example
id = uuid.New().String()
resp, err := app.MakeAccountBalanceQuery(account.AccountBalanceRequest{
    Initiator:                "apiuser",
    IdentifierType:           types.ShortCodeIdentifierType,
    SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
    OriginatorConversationID: id,
    CommandID:                types.BusinessPaymentCommand,
    PartyA:                   600000,
    Remarks:                  "Salary Payment",
    QueueTimeOutURL:          "https://yourdomain.com/timeout",
    ResultURL:                "https://yourdomain.com/result",
})
if err != nil {
    log.Fatalf("failed to make account balance query: %v", err)
}
fmt.Println("Account Balance Response: ", resp)
```

### Transaction Reversal

```go
// Transaction Reversal Example
id = uuid.New().String()
res, err := app.MakeTransactionReversalRequest(transaction.TransactionReversalRequest{
    Initiator:                "appuser",
    TransactionID:            "LKXXXX1234",
    SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
    PartyA:                   "600000",
    CommandID:                types.TransactionReversalCommand,
    IdentifierType:           types.ShortCodeIdentifierType,
    OriginatorConversationID: id,
    Amount:                   1000,
    Remarks:                  "Reversing transaction",
    ResultURL:                "https://yourdomain.com/result",
    QueueTimeOutURL:          "https://yourdomain.com/timeout",
})
if err != nil {
    log.Fatalf("failed to make transaction reversal request: %v", err)
}
fmt.Println("Transaction Reversal Response: ", res)
```

### USSD Push Payment

```go
// USSD Payment Example
id = uuid.New().String()
ressss, err := app.USSDPaymentRequest(c2b.USSDPaymentRequest{
    MerchantRequestID: id,
    BusinessShortCode: "1020",
    Password:          "M2VkZGU2YWY1Y2RhMzIyOWRjMmFkMTRiMjdjOWIwOWUxZDFlZDZiNGQ0OGYyMDRiNjg0ZDZhNWM2NTQyNTk2ZA==",
    Timestamp:         "20240918055823",
    TransactionType:   "CustomerPayBillOnline",
    Amount:            20,
    PartyA:            "251710404709",
    PartyB:            "1020",
    PhoneNumber:       "251700404709",
    CallBackURL:       "https://www.myservice:8080/result",
    AccountReference:  "Partner Unique ID",
    TransactionDesc:   "Payment Reason",
    ReferenceData: []c2b.ReferenceDataRequest{
        {
            Key:   "ThirdPartyReference",
            Value: "Ref-12345",
        },
    },
})
if err != nil {
    log.Printf("failed to make USSD payment request: %v", err)
} else {
    fmt.Println("USSD Payment Response: ", ressss)
}
```

## Contributing

1. Fork the repository.
2. Create a new branch.
3. Commit your changes.
4. Push to the branch.
5. Open a pull request.

## License

This project is licensed under the MIT License.

