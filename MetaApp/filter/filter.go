package main

import "github.com/RoaringBitmap/roaring"

func Filter1Single(item *AdsItem) bool {
	return item.A == 1
}

func Filter2Single(item *AdsItem) bool {
	return item.A == 2
}

func Filter3Single(item *AdsItem) bool {
	return item.A == 3
}

func Filter4Single(item *AdsItem) bool {
	return item.A == 4
}

func Filter5Single(item *AdsItem) bool {
	return item.B == 5
}

func Filter6Single(item *AdsItem) bool {
	return item.B == 6
}

func Filter7Single(item *AdsItem) bool {
	return item.B == 7
}

func Filter1Multi(items []*AdsItem) []*AdsItem {
	candidateAd := make([]*AdsItem, 0, len(items))
	for _, adItem := range items {
		if adItem.A != 1 {
			candidateAd = append(candidateAd, adItem)
		}
	}
	return candidateAd
}

func Filter2Multi(items []*AdsItem) []*AdsItem {
	candidateAd := make([]*AdsItem, 0, len(items))
	for _, adItem := range items {
		if adItem.A != 2 {
			candidateAd = append(candidateAd, adItem)
		}
	}
	return candidateAd
}

func Filter3Multi(items []*AdsItem) []*AdsItem {
	candidateAd := make([]*AdsItem, 0, len(items))
	for _, adItem := range items {
		if adItem.A != 3 {
			candidateAd = append(candidateAd, adItem)
		}
	}
	return candidateAd
}

func Filter4Multi(items []*AdsItem) []*AdsItem {
	candidateAd := make([]*AdsItem, 0, len(items))
	for _, adItem := range items {
		if adItem.A != 4 {
			candidateAd = append(candidateAd, adItem)
		}
	}
	return candidateAd
}

func Filter5Multi(items []*AdsItem) []*AdsItem {
	candidateAd := make([]*AdsItem, 0, len(items))
	for _, adItem := range items {
		if adItem.B != 5 {
			candidateAd = append(candidateAd, adItem)
		}
	}
	return candidateAd
}

func Filter6Multi(items []*AdsItem) []*AdsItem {
	candidateAd := make([]*AdsItem, 0, len(items))
	for _, adItem := range items {
		if adItem.B != 6 {
			candidateAd = append(candidateAd, adItem)
		}
	}
	return candidateAd
}

func Filter7Multi(items []*AdsItem) []*AdsItem {
	candidateAd := make([]*AdsItem, 0, len(items))
	for _, adItem := range items {
		if adItem.B != 7 {
			candidateAd = append(candidateAd, adItem)
		}
	}
	return candidateAd
}

func Filter1BitMap(items []*AdsItem) *roaring.Bitmap {
	tempBitMap := roaring.BitmapOf()
	for _, adItem := range items {
		if adItem.A != 1 {
			tempBitMap.Add(adItem.id)
		}
	}
	return tempBitMap
}

func Filter2BitMap(items []*AdsItem) *roaring.Bitmap {
	tempBitMap := roaring.BitmapOf()
	for _, adItem := range items {
		if adItem.A != 2 {
			tempBitMap.Add(adItem.id)
		}
	}
	return tempBitMap
}

func Filter3BitMap(items []*AdsItem) *roaring.Bitmap {
	tempBitMap := roaring.BitmapOf()
	for _, adItem := range items {
		if adItem.A != 3 {
			tempBitMap.Add(adItem.id)
		}
	}
	return tempBitMap
}

func Filter4BitMap(items []*AdsItem) *roaring.Bitmap {
	tempBitMap := roaring.BitmapOf()
	for _, adItem := range items {
		if adItem.A != 4 {
			tempBitMap.Add(adItem.id)
		}
	}
	return tempBitMap
}

func Filter5BitMap(items []*AdsItem) *roaring.Bitmap {
	tempBitMap := roaring.BitmapOf()
	for _, adItem := range items {
		if adItem.B != 5 {
			tempBitMap.Add(adItem.id)
		}
	}
	return tempBitMap
}

func Filter6BitMap(items []*AdsItem) *roaring.Bitmap {
	tempBitMap := roaring.BitmapOf()
	for _, adItem := range items {
		if adItem.B != 6 {
			tempBitMap.Add(adItem.id)
		}
	}
	return tempBitMap
}

func Filter7BitMap(items []*AdsItem) *roaring.Bitmap {
	tempBitMap := roaring.BitmapOf()
	for _, adItem := range items {
		if adItem.B != 7 {
			tempBitMap.Add(adItem.id)
		}
	}
	return tempBitMap
}
