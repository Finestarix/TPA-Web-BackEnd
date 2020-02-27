package train

import (
	models "../../models/train"
	"errors"
	"github.com/graphql-go/graphql"
)

func AllTrainClass(p graphql.ResolveParams) (i interface{}, err error) {
	trainClass := models.GetAllTrainClass()

	if len(trainClass) == 0 {
		return nil, errors.New("error: no data found")
	}

	return trainClass, nil
}

func InsertTrainClass(p graphql.ResolveParams) (i interface{}, err error) {
	name := p.Args["name"].(string)
	trainId := p.Args["trainId"].(int)

	newTrainClass := models.InsertTrainClass(name, trainId)

	return newTrainClass, nil
}
