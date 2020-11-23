package order

const (
	// TypeLimit represents that the order type is a limit.
	TypeLimit = "LIMIT"
	// TypeMarket represents that the order type is a market.
	TypeMarket = "MARKET"
	// TypeStop represents that the order type is a stop.
	TypeStop = "STOP"
	// TypeStopLimit represents that the order type is a stop limit.
	TypeStopLimit = "STOP_LIMIT"
	// TypeTrail represents that the order type is a trail.
	TypeTrail = "TRAIL"

	// SideBuy represents that the side is a buy.
	SideBuy = "BUY"
	// SideSell represents that the side is a sell.
	SideSell = "SELL"

	// TimeInForceGTC represents that the time in force is good till canceled.
	TimeInForceGTC = "GTC"
	// TimeInForceIOC represents that the time in force is immediate or cancel.
	TimeInForceIOC = "IOC"
	// TimeInForceFOK represents that the time in force is fill or kill.
	TimeInForceFOK = "FOK"

	// MethodSimple represents a special order that issues a single order.
	MethodSimple = "SIMPLE"
	// MethodIFD represents an IFD order.
	MethodIFD = "IFD"
	// MethodOCO represents an OCO order.
	MethodOCO = "OCO"
	// MethodIFDOCO represents an IFD-OCO order.
	MethodIFDOCO = "IFDOCO"
)
