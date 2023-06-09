/*
 * Portfolio - API
 *
 * This is a sample server for Kotak Trade API - Portfolio
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package portfolio

// Derivative Profit &amp; Loss Details
type PnlDerivative struct {
	// Name of the Contract
	InstrumentType string `json:"instrumentType,omitempty"`
	// Expiry Date of the Contract
	ExpiryDate string `json:"expiryDate,omitempty"`
	// Type of the Contract
	DerivativeType string `json:"derivativeType,omitempty"`
	// Strike Price of the Contract
	StrikePrice string `json:"strikePrice,omitempty"`
	// Profit and Loss of the Contract
	Pnl float32 `json:"pnl,omitempty"`
	// Sub Type of Record
	SubType string `json:"subType,omitempty"`
	// Status of Instrument Possible values ACTV, SUSP
	InstrumentStatus string `json:"instrumentStatus,omitempty"`
	// Market Lot of Security.
	MarketLot float64 `json:"marketLot,omitempty"`
	// Sector of the instrument
	Sector string `json:"sector,omitempty"`
	// Unique token for instrument
	InstrumentToken int32 `json:"instrumentToken,omitempty"`
	// squareUp of the instrument
	SquareUp string `json:"squareUp,omitempty"`
	// serviceTax for instrument
	ServiceTax int32 `json:"serviceTax,omitempty"`
	// netBrokrage for instrument
	NetBrokrage int32 `json:"netBrokrage,omitempty"`
	// otherCharges for instrument
	OtherCharges int32 `json:"otherCharges,omitempty"`
	// serviceTaxOnOthers for instrument
	ServiceTaxOnOthers int32 `json:"serviceTaxOnOthers,omitempty"`
	// stampDuty for instrument
	StampDuty int32 `json:"stampDuty,omitempty"`
	// stt for instrument
	Stt int32 `json:"stt,omitempty"`
	// ssymbol for stock
	Ssymbol string `json:"ssymbol,omitempty"`
}
