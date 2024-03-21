package structs

type CpmStatUpdateDate struct {
	UpdateDate string
	NewDate    bool
	Updated    bool
}

type AdCostStatDO struct {
	Cost int
	Show int
	Cpm  int
	Atr  float64
}
