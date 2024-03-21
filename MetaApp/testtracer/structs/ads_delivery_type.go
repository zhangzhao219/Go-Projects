package structs

type AdDeliveryType struct {
	AppType      int
	DeliveryType int
	BusinessType int
}

var AdDeliveryTypeArr = [2]AdDeliveryType{
	{10, 4, 0},
	{10, 4, 2},
}

var DeliveryType string = "%d_%d_%d"
