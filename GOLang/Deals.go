package main

import (
	"io"
	"net/http"
)

func deals(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"application/json",
	)
	io.WriteString(
		res,
		`{
			"Deals": [
				{
					"DealId": "3802e795-4db8-4a58-ae30-d548b8ad95e9",
					"DealType": "Auto",
					"BuyCurrency": "NZD",
					"BuyAmount": 1335,
					"SellCurrency": "AUD",
					"SellAmount": 1180.4,
					"Rate": 0.884196,
					"TransferMethod": "Internet Banking",
					"Mode": "Web",
					"Status": "Paid",
					"Reference": "3802E795",
					"EnteredDate": 1061846131000,
					"MaturityDate": 1061906400000,
					"MaturityDateString": "2003-8-27",
					"Verified": "H",
					"IsWholesaleDealReleased": false,
					"Converted": 1061922402000,
					"Fee": 15,
					"FeeCurrency": "NZD",
					"BPayNumber": 125100669,
					"IsPartialDeal": false,
					"IsPartOfRegularPayment": false,
					"IsDirectDebit": false,
					"IsConvertable": false,
					"IsFirstDealOfUser": false,
					"Beneficiaries": [
						"70f940da-4c38-453e-9cc1-668e4bfdb1af"
					],
					"StateInfo": {
						"DealState": "Remitted",
						"IsSettled": true,
						"IsAwaitingFunds": false,
						"AllowsConversion": false,
						"DealStateDisplay": "Funds Paid"
					},
					"MarketConventionDisplay": {
						"MarketConventionDisplay": false,
						"BuyCurrency": "NZD",
						"SellCurrency": "AUD",
						"Rate": 0.884196
					},
					"IsRdaEligible": false
				},
				{
					"DealId": "3dcc32a1-787c-11d7-9766-00902794e16e",
					"DealType": "Auto",
					"BuyCurrency": "NZD",
					"BuyAmount": 635,
					"SellCurrency": "GBP",
					"SellAmount": 220.69,
					"Rate": 0.347548,
					"TransferMethod": "Internet Banking",
					"Mode": "Web",
					"Status": "Paid",
					"Reference": "3DCC32A1",
					"EnteredDate": 1051479831000,
					"MaturityDate": 1051538400000,
					"MaturityDateString": "2003-4-29",
					"Verified": "N",
					"IsWholesaleDealReleased": false,
					"Converted": 1051554121000,
					"Fee": 15,
					"FeeCurrency": "NZD",
					"BPayNumber": 125100669,
					"IsPartialDeal": false,
					"IsPartOfRegularPayment": false,
					"IsDirectDebit": false,
					"IsConvertable": false,
					"IsFirstDealOfUser": false,
					"Beneficiaries": [
						"3dcc32ab-787c-11d7-9766-00902794e16e"
					],
					"StateInfo": {
						"DealState": "Remitted",
						"IsSettled": true,
						"IsAwaitingFunds": false,
						"AllowsConversion": false,
						"DealStateDisplay": "Funds Paid"
					},
					"MarketConventionDisplay": {
						"MarketConventionDisplay": false,
						"BuyCurrency": "NZD",
						"SellCurrency": "GBP",
						"Rate": 0.347548
					},
					"IsRdaEligible": false
				},
				{
					"DealId": "c88a7607-69a3-11d7-9766-00902794e16e",
					"DealType": "Auto",
					"BuyCurrency": "NZD",
					"BuyAmount": 360,
					"SellCurrency": "GBP",
					"SellAmount": 133.2,
					"Rate": 0.350536,
					"TransferMethod": "Internet Banking",
					"Mode": "Web",
					"Status": "Paid",
					"Reference": "C88A7607",
					"EnteredDate": 1049857447000,
					"MaturityDate": 1049896800000,
					"MaturityDateString": "2003-4-10",
					"Verified": "N",
					"IsWholesaleDealReleased": false,
					"Converted": 1049898400000,
					"Fee": 20,
					"FeeCurrency": "NZD",
					"BPayNumber": 125100669,
					"IsPartialDeal": false,
					"IsPartOfRegularPayment": false,
					"IsDirectDebit": false,
					"IsConvertable": false,
					"IsFirstDealOfUser": true,
					"Beneficiaries": [
						"c88a760a-69a3-11d7-9766-00902794e16e"
					],
					"StateInfo": {
						"DealState": "Remitted",
						"IsSettled": true,
						"IsAwaitingFunds": false,
						"AllowsConversion": false,
						"DealStateDisplay": "Funds Paid"
					},
					"MarketConventionDisplay": {
						"MarketConventionDisplay": false,
						"BuyCurrency": "NZD",
						"SellCurrency": "GBP",
						"Rate": 0.350536
					},
					"IsRdaEligible": false
				}
			],
			"Paging": {
				"Page": 1,
				"PageSize": 3,
				"PageCount": 1,
				"TotalCount": 3
			}
		}`,
	)
}
