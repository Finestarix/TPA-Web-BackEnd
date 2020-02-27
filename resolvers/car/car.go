package car

import (
	models "../../models/car"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllCar(p graphql.ResolveParams) (i interface{}, err error) {
	cars := models.GetAllCar()
	if len(cars) == 0 {
		return nil, errors.New("error: no data found")
	}

	return cars, nil
}

func GetCarByCity(p graphql.ResolveParams) (i interface{}, err error) {
	city := p.Args["city"].(string)

	cars := models.GetCarByLocation(city)

	if len(cars) == 0 {
		return nil, errors.New("error: no data found")
	}

	return cars, nil
}

func InsertCar(p graphql.ResolveParams) (i interface{}, err error) {
	locationID := p.Args["location"].(string)
	carModelID := p.Args["model"].(string)
	price := p.Args["price"].(int)

	newCars := models.InsertCar(carModelID, locationID, price)

	return newCars, nil
}
