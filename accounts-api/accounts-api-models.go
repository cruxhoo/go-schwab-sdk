package accounts_api

import "time"

type ErrorResponse struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}
type AccountNumber struct {
	AccountNumber string `json:"accountNumber"`
	HashValue     string `json:"hashValue"`
}
type Instrument struct {
	AssetType     string `json:"assetType"`
	Cusip         string `json:"cusip"`
	Symbol        string `json:"symbol"`
	UniformSymbol string `json:"uniformSymbol"`
	Description   string `json:"description"`
	Type          string `json:"type"`
}

type Position struct {
	ShortQuantity                  int         `json:"shortQuantity"`
	AveragePrice                   float64     `json:"averagePrice"`
	CurrentDayProfitLoss           int         `json:"currentDayProfitLoss"`
	CurrentDayProfitLossPercentage int         `json:"currentDayProfitLossPercentage"`
	LongQuantity                   int         `json:"longQuantity"`
	SettledLongQuantity            int         `json:"settledLongQuantity"`
	SettledShortQuantity           int         `json:"settledShortQuantity"`
	Instrument                     *Instrument `json:"instrument"`
	MarketValue                    float64     `json:"marketValue"`
	MaintenanceRequirement         float64     `json:"maintenanceRequirement"`
	AverageLongPrice               float64     `json:"averageLongPrice"`
	TaxLotAverageLongPrice         float64     `json:"taxLotAverageLongPrice"`
	LongOpenProfitLoss             float64     `json:"longOpenProfitLoss"`
	PreviousSessionLongQuantity    int         `json:"previousSessionLongQuantity"`
	CurrentDayCost                 int         `json:"currentDayCost"`
}
type InitialBalances struct {
	AccruedInterest                  int     `json:"accruedInterest"`
	AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade"`
	BondValue                        float64 `json:"bondValue"`
	BuyingPower                      float64 `json:"buyingPower"`
	CashBalance                      float64 `json:"cashBalance"`
	CashAvailableForTrading          int     `json:"cashAvailableForTrading"`
	CashReceipts                     int     `json:"cashReceipts"`
	DayTradingBuyingPower            int     `json:"dayTradingBuyingPower"`
	DayTradingBuyingPowerCall        int     `json:"dayTradingBuyingPowerCall"`
	DayTradingEquityCall             int     `json:"dayTradingEquityCall"`
	Equity                           float64 `json:"equity"`
	EquityPercentage                 int     `json:"equityPercentage"`
	LiquidationValue                 float64 `json:"liquidationValue"`
	LongMarginValue                  float64 `json:"longMarginValue"`
	LongOptionMarketValue            int     `json:"longOptionMarketValue"`
	LongStockValue                   float64 `json:"longStockValue"`
	MaintenanceCall                  int     `json:"maintenanceCall"`
	MaintenanceRequirement           int     `json:"maintenanceRequirement"`
	Margin                           int     `json:"margin"`
	MarginEquity                     float64 `json:"marginEquity"`
	MoneyMarketFund                  int     `json:"moneyMarketFund"`
	MutualFundValue                  float64 `json:"mutualFundValue"`
	RegTCall                         int     `json:"regTCall"`
	ShortMarginValue                 int     `json:"shortMarginValue"`
	ShortOptionMarketValue           int     `json:"shortOptionMarketValue"`
	ShortStockValue                  int     `json:"shortStockValue"`
	TotalCash                        int     `json:"totalCash"`
	IsInCall                         bool    `json:"isInCall"`
	PendingDeposits                  int     `json:"pendingDeposits"`
	MarginBalance                    int     `json:"marginBalance"`
	ShortBalance                     int     `json:"shortBalance"`
	AccountValue                     float64 `json:"accountValue"`
}

type CurrentBalances struct {
	AccruedInterest                  int     `json:"accruedInterest"`
	CashBalance                      float64 `json:"cashBalance"`
	CashReceipts                     int     `json:"cashReceipts"`
	LongOptionMarketValue            int     `json:"longOptionMarketValue"`
	LiquidationValue                 float64 `json:"liquidationValue"`
	LongMarketValue                  float64 `json:"longMarketValue"`
	MoneyMarketFund                  int     `json:"moneyMarketFund"`
	Savings                          int     `json:"savings"`
	ShortMarketValue                 int     `json:"shortMarketValue"`
	PendingDeposits                  int     `json:"pendingDeposits"`
	MutualFundValue                  int     `json:"mutualFundValue"`
	BondValue                        int     `json:"bondValue"`
	ShortOptionMarketValue           int     `json:"shortOptionMarketValue"`
	AvailableFunds                   float64 `json:"availableFunds"`
	AvailableFundsNonMarginableTrade int     `json:"availableFundsNonMarginableTrade"`
	BuyingPower                      float64 `json:"buyingPower"`
	BuyingPowerNonMarginableTrade    float64 `json:"buyingPowerNonMarginableTrade"`
	DayTradingBuyingPower            int     `json:"dayTradingBuyingPower"`
	Equity                           float64 `json:"equity"`
	EquityPercentage                 int     `json:"equityPercentage"`
	LongMarginValue                  float64 `json:"longMarginValue"`
	MaintenanceCall                  int     `json:"maintenanceCall"`
	MaintenanceRequirement           float64 `json:"maintenanceRequirement"`
	MarginBalance                    int     `json:"marginBalance"`
	RegTCall                         int     `json:"regTCall"`
	ShortBalance                     int     `json:"shortBalance"`
	ShortMarginValue                 int     `json:"shortMarginValue"`
	Sma                              float64 `json:"sma"`
	IntradayBuyingPowerAmount        float64 `json:"intradayBuyingPowerAmount"`
}

type ProjectedBalances struct {
	AvailableFunds                   float64 `json:"availableFunds"`
	AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade"`
	BuyingPower                      float64 `json:"buyingPower"`
	DayTradingBuyingPower            int     `json:"dayTradingBuyingPower"`
	DayTradingBuyingPowerCall        int     `json:"dayTradingBuyingPowerCall"`
	MaintenanceCall                  int     `json:"maintenanceCall"`
	RegTCall                         int     `json:"regTCall"`
	IsInCall                         bool    `json:"isInCall"`
	StockBuyingPower                 float64 `json:"stockBuyingPower"`
	IntradayBuyingPowerAmount        float64 `json:"intradayBuyingPowerAmount"`
}

type SecuritiesAccount struct {
	Type                    string            `json:"type"`
	AccountNumber           string            `json:"accountNumber"`
	RoundTrips              int               `json:"roundTrips"`
	IsDayTrader             bool              `json:"isDayTrader"`
	IsClosingOnlyRestricted bool              `json:"isClosingOnlyRestricted"`
	IsPortfolioMargin       bool              `json:"isPortfolioMargin"`
	PfcbFlag                bool              `json:"pfcbFlag"`
	Positions               []Position        `json:"positions"`
	IsIntradayMargin        bool              `json:"isIntradayMargin"`
	InitialBalances         InitialBalances   `json:"initialBalances"`
	CurrentBalances         CurrentBalances   `json:"currentBalances"`
	ProjectedBalances       ProjectedBalances `json:"projectedBalances"`
}

type AggregatedBalance struct {
	CurrentLiquidationValue          float64 `json:"currentLiquidationValue"`
	LiquidationValue                 float64 `json:"liquidationValue"`
	IntradayBuyingPowerAmount        float64 `json:"intradayBuyingPowerAmount"`
	CurrentIntradayBuyingPowerAmount float64 `json:"currentIntradayBuyingPowerAmount"`
}

type Account struct {
	SecuritiesAccount SecuritiesAccount `json:"securitiesAccount"`
	AggregatedBalance AggregatedBalance `json:"aggregatedBalance"`
}

type OrderLeg struct {
	OrderLegType   string     `json:"orderLegType"`
	LegId          int        `json:"legId"`
	Instrument     Instrument `json:"instrument"`
	Instruction    string     `json:"instruction"`
	PositionEffect string     `json:"positionEffect"`
	Quantity       int        `json:"quantity"`
}

type ExecutionLeg struct {
	LegId             int     `json:"legId"`
	Quantity          int     `json:"quantity"`
	MismarkedQuantity int     `json:"mismarkedQuantity"`
	Price             float64 `json:"price"`
	Time              string  `json:"time"`
	InstrumentId      int     `json:"instrumentId"`
}

type OrderActivity struct {
	ActivityType           string         `json:"activityType"`
	ActivityId             int64          `json:"activityId"`
	ExecutionType          string         `json:"executionType"`
	Quantity               int            `json:"quantity"`
	OrderRemainingQuantity int            `json:"orderRemainingQuantity"`
	ExecutionLegs          []ExecutionLeg `json:"executionLegs"`
}

type Order struct {
	Session                  string          `json:"session"`
	Duration                 string          `json:"duration"`
	OrderType                string          `json:"orderType"`
	ComplexOrderStrategyType string          `json:"complexOrderStrategyType"`
	Quantity                 int             `json:"quantity"`
	FilledQuantity           int             `json:"filledQuantity"`
	RemainingQuantity        int             `json:"remainingQuantity"`
	RequestedDestination     string          `json:"requestedDestination"`
	DestinationLinkName      string          `json:"destinationLinkName"`
	Price                    float64         `json:"price"`
	OrderLegCollection       []OrderLeg      `json:"orderLegCollection"`
	OrderStrategyType        string          `json:"orderStrategyType"`
	OrderId                  int64           `json:"orderId"`
	Cancelable               bool            `json:"cancelable"`
	Editable                 bool            `json:"editable"`
	Status                   string          `json:"status"`
	EnteredTime              string          `json:"enteredTime"`
	CloseTime                string          `json:"closeTime"`
	AccountNumber            int             `json:"accountNumber"`
	OrderActivityCollection  []OrderActivity `json:"orderActivityCollection"`
}

type TransferItemInstrument struct {
	AssetType     string  `json:"assetType"`
	Status        string  `json:"status"`
	Symbol        string  `json:"symbol"`
	Description   string  `json:"description,omitempty"`
	InstrumentId  int     `json:"instrumentId"`
	ClosingPrice  float64 `json:"closingPrice"`
	UniformSymbol string  `json:"uniformSymbol,omitempty"`
	Type          string  `json:"type,omitempty"`
}

type TransferItem struct {
	Instrument     TransferItemInstrument `json:"instrument"`
	Amount         float64                `json:"amount"`
	Cost           float64                `json:"cost"`
	FeeType        string                 `json:"feeType,omitempty"`
	Price          float64                `json:"price,omitempty"`
	PositionEffect string                 `json:"positionEffect,omitempty"`
}

type Transaction struct {
	ActivityId    int64          `json:"activityId"`
	Time          string         `json:"time"`
	AccountNumber string         `json:"accountNumber"`
	Type          string         `json:"type"`
	Status        string         `json:"status"`
	SubAccount    string         `json:"subAccount"`
	TradeDate     string         `json:"tradeDate"`
	PositionId    int64          `json:"positionId"`
	OrderId       int64          `json:"orderId"`
	NetAmount     float64        `json:"netAmount"`
	TransferItems []TransferItem `json:"transferItems"`
}

type UserPreferenceAccount struct {
	AccountNumber      string `json:"accountNumber"`
	PrimaryAccount     bool   `json:"primaryAccount"`
	Type               string `json:"type"`
	NickName           string `json:"nickName"`
	DisplayAcctId      string `json:"displayAcctId"`
	AutoPositionEffect bool   `json:"autoPositionEffect"`
	AccountColor       string `json:"accountColor"`
	LotSelectionMethod string `json:"lotSelectionMethod"`
	HasFuturesAccount  bool   `json:"hasFuturesAccount"`
	HasForexAccount    bool   `json:"hasForexAccount"`
}

type StreamerInfo struct {
	StreamerSocketUrl      string `json:"streamerSocketUrl"`
	SchwabClientCustomerId string `json:"schwabClientCustomerId"`
	SchwabClientCorrelId   string `json:"schwabClientCorrelId"`
	SchwabClientChannel    string `json:"schwabClientChannel"`
	SchwabClientFunctionId string `json:"schwabClientFunctionId"`
}

type Offer struct {
	Level2Permissions bool   `json:"level2Permissions"`
	MktDataPermission string `json:"mktDataPermission"`
}

type UserPreference struct {
	Accounts     []UserPreferenceAccount `json:"accounts"`
	StreamerInfo []StreamerInfo          `json:"streamerInfo"`
	Offers       []Offer                 `json:"offers"`
}

type OrderPreviewOrderBalance struct {
	OrderValue             int `json:"orderValue"`
	ProjectedAvailableFund int `json:"projectedAvailableFund"`
	ProjectedBuyingPower   int `json:"projectedBuyingPower"`
	ProjectedCommission    int `json:"projectedCommission"`
}

type OrderPreviewOrderLeg struct {
	AskPrice            int    `json:"askPrice"`
	BidPrice            int    `json:"bidPrice"`
	LastPrice           int    `json:"lastPrice"`
	MarkPrice           int    `json:"markPrice"`
	ProjectedCommission int    `json:"projectedCommission"`
	Quantity            int    `json:"quantity"`
	FinalSymbol         string `json:"finalSymbol"`
	LegId               int    `json:"legId"`
	AssetType           string `json:"assetType"`
	Instruction         string `json:"instruction"`
}

type OrderPreviewOrderStrategy struct {
	AccountNumber          string                    `json:"accountNumber"`
	AdvancedOrderType      string                    `json:"advancedOrderType"`
	CloseTime              time.Time                 `json:"closeTime"`
	EnteredTime            time.Time                 `json:"enteredTime"`
	OrderBalance           *OrderPreviewOrderBalance `json:"orderBalance"`
	OrderStrategyType      string                    `json:"orderStrategyType"`
	OrderVersion           int                       `json:"orderVersion"`
	Session                string                    `json:"session"`
	Status                 string                    `json:"status"`
	AllOrNone              bool                      `json:"allOrNone"`
	Discretionary          bool                      `json:"discretionary"`
	Duration               string                    `json:"duration"`
	FilledQuantity         int                       `json:"filledQuantity"`
	OrderType              string                    `json:"orderType"`
	OrderValue             int                       `json:"orderValue"`
	Price                  int                       `json:"price"`
	Quantity               int                       `json:"quantity"`
	RemainingQuantity      int                       `json:"remainingQuantity"`
	SellNonMarginableFirst bool                      `json:"sellNonMarginableFirst"`
	SettlementInstruction  string                    `json:"settlementInstruction"`
	Strategy               string                    `json:"strategy"`
	AmountIndicator        string                    `json:"amountIndicator"`
	OrderLegs              []OrderPreviewOrderLeg    `json:"orderLegs"`
}

type OrderValidationDetail struct {
	ValidationRuleName string `json:"validationRuleName"`
	Message            string `json:"message"`
	ActivityMessage    string `json:"activityMessage"`
	OriginalSeverity   string `json:"originalSeverity"`
	OverrideName       string `json:"overrideName"`
	OverrideSeverity   string `json:"overrideSeverity"`
}

type OrderValidationResult struct {
	Alerts  []OrderValidationDetail `json:"alerts"`
	Accepts []OrderValidationDetail `json:"accepts"`
	Rejects []OrderValidationDetail `json:"rejects"`
	Reviews []OrderValidationDetail `json:"reviews"`
	Warns   []OrderValidationDetail `json:"warns"`
}

type CommissionValue struct {
	Value int    `json:"value"`
	Type  string `json:"type"`
}

type CommissionLeg struct {
	CommissionValues []CommissionValue `json:"commissionValues"`
}

type Commission struct {
	CommissionLegs []CommissionLeg `json:"commissionLegs"`
}

type FeeValue struct {
	Value int    `json:"value"`
	Type  string `json:"type"`
}

type FeeLeg struct {
	FeeValues []FeeValue `json:"feeValues"`
}

type Fee struct {
	FeeLegs []FeeLeg `json:"feeLegs"`
}

type CommissionAndFee struct {
	Commission     Commission `json:"commission"`
	Fee            Fee        `json:"fee"`
	TrueCommission Commission `json:"trueCommission"`
}

type OrderPreview struct {
	OrderId               int                        `json:"orderId"`
	OrderStrategy         *OrderPreviewOrderStrategy `json:"orderStrategy"`
	OrderValidationResult *OrderValidationResult     `json:"orderValidationResult"`
	CommissionAndFee      *CommissionAndFee          `json:"commissionAndFee"`
}
