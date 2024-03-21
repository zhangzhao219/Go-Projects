package structs

type CtrParamStruct struct {
	AvgPCtr, Ks, Bs []float64
}

type CtrCalibSmoothStruct struct {
	AidSmoothCtr map[string]float64
}
