package event

import (
	models "../../models/event"
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

func AllEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	event := models.GetAllEntertainment()
	if len(event) == 0 {
		return nil, errors.New("error: no data found")
	}

	return event, nil
}

func GetEntertainmentByCategory(p graphql.ResolveParams) (i interface{}, err error) {
	category := p.Args["category"].(string)

	event := models.GetEntertainmentByCategory(category)
	if len(event) == 0 {
		return nil, errors.New("error: no data found")
	}

	return event, nil
}

func GetFilterEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	lowPrice := p.Args["lowPrice"].(int)
	highPrice := p.Args["highPrice"].(int)
	location := p.Args["location"].(string)
	startDate := p.Args["startDate"].(string)
	endDate	 := p.Args["endDate"].(string)
	isActivity := p.Args["isActivity"].(string)
	isEvent := p.Args["isEvent"].(string)
	isAttraction := p.Args["isAttraction"].(string)

	startDateConv, _ := time.Parse(time.RFC3339, startDate)
	endDateConv, _ := time.Parse(time.RFC3339, endDate)

	event := models.GetFilterEntertainment(lowPrice, highPrice,
		startDateConv, endDateConv, isActivity, isAttraction, isEvent, location)

	return event, nil
}

func InsertEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	title := p.Args["title"].(string)
	price := p.Args["price"].(int)
	location := p.Args["location"].(string)
	latitude := p.Args["latitude"].(float64)
	longitude := p.Args["longitude"].(float64)
	date := p.Args["date"].(string)
	category := p.Args["category"].(string)
	image := p.Args["image"].(string)
	description := p.Args["description"].(string)
	termCondition := p.Args["termCondition"].(string)

	dateConv, _ := time.Parse(time.RFC3339, date)

	newEvent := models.InsertEntertainment(title, price, location,
		latitude, longitude, dateConv, category, image, description, termCondition)

	return newEvent, nil
}

func UpdateEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)
	title := p.Args["title"].(string)
	price := p.Args["price"].(int)
	location := p.Args["location"].(string)
	latitude := p.Args["latitude"].(float64)
	longitude := p.Args["longitude"].(float64)
	date := p.Args["date"].(string)
	category := p.Args["category"].(string)
	image := p.Args["image"].(string)
	description := p.Args["description"].(string)
	termCondition := p.Args["termCondition"].(string)

	dateConv, _ := time.Parse(time.RFC3339, date)

	event := models.UpdateEntertainment(id, title, price, location,
		latitude, longitude, dateConv, category, image, description, termCondition)

	return event, nil
}

func DeleteEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	event := models.DeleteEntertainment(id)

	return event, nil
}
