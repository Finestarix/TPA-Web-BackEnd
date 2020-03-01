package event

import (
	models "../../models/event"
	"errors"
	"github.com/graphql-go/graphql"
	"strconv"
	"time"
)

func AllEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	event := models.GetAllEntertainment()
	if len(event) == 0 {
		return nil, errors.New("error: no data found")
	}

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

	dateConv, _ := time.Parse(time.RFC3339, date)

	newEvent := models.InsertEntertainment(title, price, location, latitude, longitude, dateConv, category, image)

	return newEvent, nil
}

func UpdateEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(string)
	title := p.Args["title"].(string)
	price := p.Args["price"].(int)
	location := p.Args["location"].(string)
	latitude := p.Args["latitude"].(float64)
	longitude := p.Args["longitude"].(float64)
	date := p.Args["date"].(string)
	category := p.Args["category"].(string)
	image := p.Args["image"].(string)

	intConv, _ :=  strconv.Atoi(id)
	dateConv, _ := time.Parse(time.RFC3339, date)

	event := models.UpdateEntertainent(intConv, title, price, location, latitude, longitude, dateConv, category, image)

	return event, nil
}

func DeleteEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	event := models.DeleteEntertainment(id)

	return event, nil
}
