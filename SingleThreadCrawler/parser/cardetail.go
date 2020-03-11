package parser

import (
	"GoLangIntro/SingleThreadCrawler/engine"
	"GoLangIntro/SingleThreadCrawler/model"
	"fmt"
	"regexp"
	"strconv"
)

var urlRegexp = `http://newcar.xcar.com.cn/(m\d+)/`
var regexpPtrId = regexp.MustCompile(urlRegexp)
var nameRegexp = "<title>【(.*)】报价_图片_参数-爱卡汽车</title>"
var regexpPtrName = regexp.MustCompile(nameRegexp)
var imageRegexp = `<img class="color_car_img_new" src="([^"]+)"`
var regexpPtrImage = regexp.MustCompile(imageRegexp)
var sizeRegexp = `<li.*车身尺寸.*<em>(\d+[^\d]\d+[^\d]\d+mm)`
var regexpPtrSize = regexp.MustCompile(sizeRegexp)
var fuelRegexp = `<span.*工信部油耗.*<em>(\d+\.\d+)L/100km</em>`
var regexpPtrFuel = regexp.MustCompile(fuelRegexp)
var transmissionRegexp = `<li>.*变\s*速\s*箱.*<em>(.+)</em></li>`
var regexpPtrTransmission = regexp.MustCompile(transmissionRegexp)
var displacementRegexp = `<li.*排.*量.*(\d+\.\d+)L`
var regexpPtrDisplacement = regexp.MustCompile(displacementRegexp)
var maxSpeedRegexp = ` <td.*最高车速\(km/h\).*\s*<td[^>]*>(\d+)</td>`
var regexpPtrMaxSpeed = regexp.MustCompile(maxSpeedRegexp)
var accelerationRegexp = `<td.*0-100加速时间\(s\).*\s*<td[^>]*>(\d+\.\d+)</td>`
var regexpPtrAcceleration = regexp.MustCompile(accelerationRegexp)
var priceRegexp = `<a href="/%s/baojia/".*>(\d+.\d+)</a>`

// This function goes over a car detail page and return all Car information in
// the ParseResult.Item object
func ParseCarDetail(contents []byte, url string) engine.ParseResult {
	id := extractString([]byte(url), regexpPtrId)

	car := model.Car{
		Name:         extractString(contents, regexpPtrName),
		ImageURL:     "http:" + extractString(contents, regexpPtrImage),
		Size:         extractString(contents, regexpPtrSize),
		Fuel:         extractFloat(contents, regexpPtrFuel),
		Transmission: extractString(contents, regexpPtrTransmission),
		Displacement: extractFloat(contents, regexpPtrDisplacement),
		MaxSpeed:     extractFloat(contents, regexpPtrMaxSpeed),
		Acceleration: extractFloat(contents, regexpPtrAcceleration),
	}

	regexpPtrPrice, err := regexp.Compile(fmt.Sprintf(priceRegexp, id))
	if err == nil {
		car.Price = extractFloat(contents, regexpPtrPrice)
	}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Payload: car,
				Type:    "xcar",
				Id:      id,
			},
		},
	}

	newCarModelResult := ParseCarModel(contents, url)
	result.Requests = newCarModelResult.Requests

	return result
}

// Extract string from content byte slice based on regexp
func extractString(contents []byte, regexpPtr *regexp.Regexp) string {
	match := regexpPtr.FindSubmatch(contents)

	if len(match) > 1 {
		return string(match[1])
	} else {
		return ""
	}
}

// Extract float64 number from content byte slice based on regexp
func extractFloat(contents []byte, regexpPtr *regexp.Regexp) float64 {
	numStr := extractString(contents, regexpPtr)
	f, err := strconv.ParseFloat(numStr, 64)

	if err != nil {
		return 0.0
	}
	return f
}
