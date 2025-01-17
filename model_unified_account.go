/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type UnifiedAccount struct {
	// User ID
	UserId int64 `json:"user_id,omitempty"`
	// Time of the most recent refresh
	RefreshTime int64 `json:"refresh_time,omitempty"`
	// Whether account is locked
	Locked   bool                      `json:"locked,omitempty"`
	Balances map[string]UnifiedBalance `json:"balances,omitempty"`
	// The total asset value in USDT. Sum of `(available + freeze) * price`
	Total string `json:"total,omitempty"`
	// The total borrowed amount in USDT equivalent. Sum of `borrowed * price`
	Borrowed string `json:"borrowed,omitempty"`
	// Total initial margin
	TotalInitialMargin string `json:"total_initial_margin,omitempty"`
	// Total margin balance
	TotalMarginBalance string `json:"total_margin_balance,omitempty"`
	// Total maintenance margin
	TotalMaintenanceMargin string `json:"total_maintenance_margin,omitempty"`
	// Total initial margin rate
	TotalInitialMarginRate string `json:"total_initial_margin_rate,omitempty"`
	// Total maintenance margin rate
	TotalMaintenanceMarginRate string `json:"total_maintenance_margin_rate,omitempty"`
	// Total available margin
	TotalAvailableMargin string `json:"total_available_margin,omitempty"`
	// Total amount of the portfolio margin account
	UnifiedAccountTotal string `json:"unified_account_total,omitempty"`
	// Total liabilities of the portfolio margin account
	UnifiedAccountTotalLiab string `json:"unified_account_total_liab,omitempty"`
	// Total equity of the portfolio margin account
	UnifiedAccountTotalEquity string `json:"unified_account_total_equity,omitempty"`
	// Leverage
	Leverage string `json:"leverage,omitempty"`
}
