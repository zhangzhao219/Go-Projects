package main

import (
	"errors"
	"fmt"
	"sync"
)

type TemplateData interface {
	getTemplateData(data []string) (map[string]interface{}, error)
}

var doTep sync.Once
var Template = make(map[string]TemplateData)

// FactoryTemplate ...
// tagID: 模板ID
// data: 对应模板参数keyword
func FactoryTemplate(tagID string, data []string) (map[string]interface{}, error) {
	doTep.Do(func() {
		Template["1"] = new(PayOk)
		Template["2"] = new(AutoCancelOrder)
		Template["3"] = new(OrderDelivery)
	})

	if _, ok := Template[tagID]; !ok {
		return nil, errors.New("tagID invalid")
	}

	return Template[tagID].getTemplateData(data)
}

func main() {
	fmt.Println(FactoryTemplate("1", []string{"aaaaa", "bbbbb", "ccccc"}))
}

// PayOk 1支付成功
type PayOk struct{}

func (p PayOk) getTemplateData(data []string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"first": map[string]interface{}{
			"value": "尊敬的用户，您的订单已支付成功，我们会尽快为您发货。",
			"color": "erfg",
		},
		"keyword2": map[string]interface{}{
			"value": data[1],
			"color": "erfg",
		},
		"keyword1": map[string]interface{}{
			"value": data[0],
			"color": "erfg",
		},

		"remark": map[string]interface{}{
			"value": "请耐心等待收货，收到货后记得回来确认哦。",
			"color": "erfg",
		},
	}, nil
}

// AutoCancelOrder 2自动取消订单
type AutoCancelOrder struct{}

func (a AutoCancelOrder) getTemplateData(data []string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"first": map[string]interface{}{
			"value": "尊敬的用户，您的订单未在指定时间内支付，已自动取消。",
			"color": "erfg",
		},
		"keyword2": map[string]interface{}{
			"value": data[1],
			"color": "erfg",
		},
		"keyword1": map[string]interface{}{
			"value": data[0],
			"color": "erfg",
		},

		"remark": map[string]interface{}{
			"value": "感谢您的购买。",
			"color": "erfg",
		},
	}, nil
}

// OrderDelivery 3订单发货
type OrderDelivery struct{}

func (o OrderDelivery) getTemplateData(data []string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"first": map[string]interface{}{
			"value": "尊敬的用户，您的订单已发货，正快马加鞭向您飞奔而去。",
			"color": "erfg",
		},
		"keyword4": map[string]interface{}{
			"value": data[3],
			"color": "erfg",
		},
		"keyword1": map[string]interface{}{
			"value": data[0],
			"color": "erfg",
		},
		"keyword2": map[string]interface{}{
			"value": data[1],
			"color": "erfg",
		},
		"keyword3": map[string]interface{}{
			"value": data[2],
			"color": "erfg",
		},

		"remark": map[string]interface{}{
			"value": "请前往订单详情，查看物流详情。",
			"color": "erfg",
		},
	}, nil
}

// OrderAutoTake 4订单自动收货
type OrderAutoTake struct{}

func (o OrderAutoTake) getTemplateData(data []string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"first": map[string]interface{}{
			"value": "尊敬的客户，您的订单已自动收货：",
			"color": "erfg",
		},
		"keyword3": map[string]interface{}{
			"value": data[2],
			"color": "erfg",
		},
		"keyword1": map[string]interface{}{
			"value": data[0],
			"color": "erfg",
		},
		"keyword2": map[string]interface{}{
			"value": data[1],
			"color": "erfg",
		},

		"remark": map[string]interface{}{
			"value": "祝您购物愉快。",
			"color": "erfg",
		},
	}, nil
}
