package structs

type OcpxReplacementPkgBid struct {
	OldPkg string  `json:"old_pkg"`
	Bid    float64 `json:"bid"`
}
