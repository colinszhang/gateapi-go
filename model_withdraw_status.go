/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type WithdrawStatus struct {
	// Currency
	Currency string `json:"currency,omitempty"`
	// Currency name
	Name string `json:"name,omitempty"`
	// Currency Chinese name
	NameCn string `json:"name_cn,omitempty"`
	// Deposits fee
	Deposit string `json:"deposit,omitempty"`
	// Withdrawal fee rate percentage
	WithdrawPercent string `json:"withdraw_percent,omitempty"`
	// Fixed withdrawal fee
	WithdrawFix string `json:"withdraw_fix,omitempty"`
	// Daily allowed withdrawal amount
	WithdrawDayLimit string `json:"withdraw_day_limit,omitempty"`
	// Minimum withdrawal amount
	WithdrawAmountMini string `json:"withdraw_amount_mini,omitempty"`
	// Daily withdrawal amount left
	WithdrawDayLimitRemain string `json:"withdraw_day_limit_remain,omitempty"`
	// Maximum amount for each withdrawal
	WithdrawEachtimeLimit string `json:"withdraw_eachtime_limit,omitempty"`
}
