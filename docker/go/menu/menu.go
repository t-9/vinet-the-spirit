package menu

const (
	// Exit represents exit program.
	Exit int64 = iota
	// ShowMarkets represents the market display.
	ShowMarkets
	// ShowBoard represents showing the board.
	ShowBoard
	// RegisterAccessKey represents registering an Access Key.
	RegisterAccessKey
	// ShowBalance represents showing abalances.
	ShowBalance
	// ShowCoinIn represents showing coinins.
	ShowCoinIn
	// ShowDeposit represents showing deposits.
	ShowDeposit
	// ShowAddress represents showing addresses.
	ShowAddress
	// SendChildOrder represents sending a child order.
	SendChildOrder
	// SendParentOrder represents sending a parent order.
	SendParentOrder
	// ShowParentOrder represents showing parent orders.
	ShowParentOrder
	// CancelParentOrder represents canceling a parent order.
	CancelParentOrder
)
