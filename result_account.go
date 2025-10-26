package scrapfly

// --- Account Data Structures ---
type Account struct {
	AccountID        string `json:"account_id"`
	Currency         string `json:"currency"`
	Timezone         string `json:"timezone"`
	Suspended        bool   `json:"suspended"`
	SuspensionReason string `json:"suspension_reason"`
}

type Project struct {
	AllowExtraUsage    bool     `json:"allow_extra_usage"`
	AllowedNetworks    []string `json:"allowed_networks"`
	BudgetLimit        *float32 `json:"budget_limit"`
	BudgetSpent        *float32 `json:"budget_spent"`
	ConcurrencyLimit   *int     `json:"concurrency_limit"`
	Name               string   `json:"name"`
	QuotaReached       bool     `json:"quota_reached"`
	ScrapeRequestCount int      `json:"scrape_request_count"`
	ScrapeRequestLimit *int     `json:"scrape_request_limit"`
	Tags               []string `json:"tags"`
}

type Subscription struct {
	Billing            SubscriptionBilling `json:"billing"`
	ExtraScrapeAllowed bool                `json:"extra_scrape_allowed"`
	MaxConcurrency     int                 `json:"max_concurrency"`
	Period             SubscriptionPeriod  `json:"period"`
	PlanName           string              `json:"plan_name"`
	Usage              SubscriptionUsage   `json:"usage"`
}

type SubscriptionUsage struct {
	Spider   SpiderUsage   `json:"spider"`
	Schedule ScheduleUsage `json:"schedule"`
	Scrape   ScrapeUsage   `json:"scrape"`
}

type SpiderUsage struct {
	Current int `json:"current"`
	Limit   int `json:"limit"`
}

type ScheduleUsage struct {
	Current int `json:"current"`
	Limit   int `json:"limit"`
}

type ScrapeUsage struct {
	ConcurrentLimit     int `json:"concurrent_limit"`
	ConcurrentRemaining int `json:"concurrent_remaining"`
	ConcurrentUsage     int `json:"concurrent_usage"`
	Current             int `json:"current"`
	Extra               int `json:"extra"`
	Limit               int `json:"limit"`
	Remaining           int `json:"remaining"`
}

type SubscriptionPeriod struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type SubscriptionIntegerPrice struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

type SubscriptionFloatPrice struct {
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
}

type SubscriptionBilling struct {
	CurrentExtraScrapeRequestPrice SubscriptionIntegerPrice `json:"current_extra_scrape_request_price"`
	ExtraScrapeRequestPricePer10k  SubscriptionFloatPrice   `json:"extra_scrape_request_price_per_10k"`
	OngoingPayment                 SubscriptionFloatPrice   `json:"ongoing_payment"`
	PlanPrice                      SubscriptionIntegerPrice `json:"plan_price"`
}

// AccountData represents detailed account information from Scrapfly.
//
// This includes subscription details, usage statistics, billing information,
// and account limits.
type AccountData struct {
	Account      Account      `json:"account"`
	Project      Project      `json:"project"`
	Subscription Subscription `json:"subscription"`
}
