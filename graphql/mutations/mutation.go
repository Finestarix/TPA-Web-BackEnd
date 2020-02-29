package mutations

import (
	rBlog "../../resolvers/blog"
	rCar "../../resolvers/car"
	rCore "../../resolvers/core"
	rFlight "../../resolvers/flight"
	rHotel "../../resolvers/hotel"
	rTrain "../../resolvers/train"
	rUser "../../resolvers/user"
	tBlog "../types/blog"
	tCar "../types/car"
	tCore "../types/core"
	tFlight "../types/flight"
	tHotel "../types/hotel"
	tTrain "../types/train"
	tUser "../types/user"
	"github.com/graphql-go/graphql"
)

func Root() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"InsertNewSubscription": &graphql.Field{
				Type: tCore.GetSubscriptionType(),
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rCore.InsertSubscription,
			},
			"InsertNewLocation": &graphql.Field{
				Type: tCore.GetLocationType(),
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
				Resolve: rCore.InsertLocation,
			},

			"InsertNewAdmin": &graphql.Field{
				Type: tUser.GetAdminType(),
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rUser.InsertAdmin,
			},
			"InsertNewUser": &graphql.Field{
				Type: tUser.GetUserType(),
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
				Resolve: rUser.InsertUser,
			},
			"UpdateUserProfile": &graphql.Field{
				Type: tUser.GetUserType(),
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
				Resolve: rUser.UpdateUserProfile,
			},
			"DeleteUser": &graphql.Field{
				Type: tUser.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: rUser.DeleteUser,
			},

			"InsertNewHotel": &graphql.Field{
				Type: tHotel.GetHotelType(),
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
				Resolve: rHotel.InsertHotel,
			},
			"UpdateHotel": &graphql.Field{
				Type: tHotel.GetHotelType(),
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
				Resolve: rHotel.UpdateHotel,
			},
			"DeleteHotel": &graphql.Field{
				Type: tHotel.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rHotel.DeleteHotel,
			},
			"InsertHotelImage": &graphql.Field{
				Type: tHotel.GetHotelImageType(),
				Args: graphql.FieldConfigArgument{
					"hotelid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"source": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rHotel.InsertHotelImage,
			},
			"InsertHotelFacility": &graphql.Field{
				Type: tHotel.GetHotelFacilityType(),
				Args: graphql.FieldConfigArgument{
					"hotelid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rHotel.InsertHotelFacility,
			},
			"InsertHotelType": &graphql.Field{
				Type: tHotel.GetHotelTypeType(),
				Args: graphql.FieldConfigArgument{
					"hotelid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rHotel.InsertHotelType,
			},

			"InsertNewCarModel": &graphql.Field{
				Type: tCar.GetCarModelType(),
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
				Resolve: rCar.InsertCarModel,
			},
			"InsertNewCar": &graphql.Field{
				Type: tCar.GetCarType(),
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
				Resolve: rCar.InsertCar,
			},

			"InsertNewTrain": &graphql.Field{
				Type: tTrain.GetTrainType(),
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
				Resolve: rTrain.InsertTrain,
			},
			"UpdateTrain": &graphql.Field{
				Type: tTrain.GetTrainType(),
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
				Resolve: rTrain.UpdateTrain,
			},
			"DeleteTrain": &graphql.Field{
				Type: tTrain.GetTrainType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rTrain.DeleteTrain,
			},
			"InsertNewTrainClass": &graphql.Field{
				Type: tTrain.GetTrainClassType(),
				Args: graphql.FieldConfigArgument{
					"trainId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rTrain.InsertTrainClass,
			},
			"InsertNewTrainStation": &graphql.Field{
				Type: tTrain.GetTrainStationType(),
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
				Resolve: rTrain.InsertTrainStation,
			},

			"InsertNewFlight": &graphql.Field{
				Type: tFlight.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"company": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"fromAirport": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"toAirport": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"transitAirport": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"arrivalTime": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"departureTime": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"model": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rFlight.InsertFlight,
			},
			"UpdateFlight": &graphql.Field{
				Type: tFlight.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"fromAirport": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"toAirport": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"transitAirport": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"arrivalTime": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"departureTime": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"model": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rFlight.UpdateFlight,
			},
			"DeleteFlight": &graphql.Field{
				Type: tFlight.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: rFlight.DeleteFlight,
			},
			"InsertNewFlightAirport": &graphql.Field{
				Type: tFlight.GetFlightAirportType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"code": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rFlight.InsertFlightAirport,
			},
			"InsertNewFlightCompany": &graphql.Field{
				Type: tFlight.GetFlightCompanyType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rFlight.InsertFlightCompany,
			},
			"InsertNewFlightFacility": &graphql.Field{
				Type: tFlight.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"flightID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rFlight.InsertFlightFacility,
			},

			"InsertNewBlog": &graphql.Field{
				Type: tBlog.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"userID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rBlog.InsertBlog,
			},
			"UpdateBlog": &graphql.Field{
				Type: tBlog.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: rBlog.UpdateBlog,
			},
			"DeleteBlog": &graphql.Field{
				Type: tBlog.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: rBlog.DeleteBlog,
			},
		},
	})
}
