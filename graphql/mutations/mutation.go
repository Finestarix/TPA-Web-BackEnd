package mutations

import (
	"../../resolvers"
	"../types"
	"github.com/graphql-go/graphql"
)

func Root() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"InsertNewSubscription": &graphql.Field{
				Type: types.GetSubscriptionType(),
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertSubscription,
			},

			"InsertNewAdmin": &graphql.Field{
				Type: types.GetAdminType(),
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertAdmin,
			},

			"InsertNewUser": &graphql.Field{
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"firstname": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastname": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"phonecode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertUser,
			},
			"UpdateUserProfile": &graphql.Field{
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"firstname": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastname": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"zipcode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.UpdateUserProfile,
			},
			"DeleteUser": &graphql.Field{
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.DeleteUser,
			},

			"InsertNewHotel": &graphql.Field{
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"rating": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"information": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertHotel,
			},
			"UpdateHotel": &graphql.Field{
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"rating": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"information": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.UpdateHotel,
			},
			"DeleteHotel": &graphql.Field{
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.DeleteHotel,
			},

			"InsertNewLocation": &graphql.Field{
				Type: types.GetLocationType(),
				Args: graphql.FieldConfigArgument{
					"city": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"province": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"region": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertLocation,
			},

			"InsertHotelImage": &graphql.Field{
				Type: types.GetHotelImageType(),
				Args: graphql.FieldConfigArgument{
					"hotelid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"source": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertHotelImage,
			},

			"InsertHotelFacility": &graphql.Field{
				Type: types.GetHotelFacilityType(),
				Args: graphql.FieldConfigArgument{
					"hotelid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertHotelFacility,
			},

			"InsertHotelType": &graphql.Field{
				Type: types.GetHotelTypeType(),
				Args: graphql.FieldConfigArgument{
					"hotelid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertHotelType,
			},

			"InsertNewCarModel": &graphql.Field{
				Type: types.GetCarModelType(),
				Args: graphql.FieldConfigArgument{
					"brand": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"model": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"passenger": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"baggage": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertCarModel,
			},

			"InsertNewCar": &graphql.Field{
				Type: types.GetCarType(),
				Args: graphql.FieldConfigArgument{
					"location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"model": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.InsertCar,
			},

			"InsertNewTrain": &graphql.Field{
				Type: types.GetTrainType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"code": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"arrivalTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"transit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"departureTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"seat": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.InsertTrain,
			},
			"UpdateTrain": &graphql.Field{
				Type: types.GetTrainType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"arrivalTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"departureTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"seat": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.UpdateTrain,
			},
			"DeleteTrain": &graphql.Field{
				Type: types.GetTrainType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.DeleteTrain,
			},

			"InsertNewTrainClass": &graphql.Field{
				Type: types.GetTrainClassType(),
				Args: graphql.FieldConfigArgument{
					"trainId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertTrainClass,
			},

			"InsertNewTrainStation": &graphql.Field{
				Type: types.GetTrainStationType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"code": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.InsertTrainStation,
			},
		},
	})
}
