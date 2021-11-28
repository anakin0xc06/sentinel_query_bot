package types

type Price struct {
	Symbol              string
	Current             float32
	ATH                 float64
	ATL                 float64
	High24h             float64
	Low24h              float64
	ChangePercentage24h float64
}
