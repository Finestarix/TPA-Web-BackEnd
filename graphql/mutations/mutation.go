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
						Type: graphql.NewNonNull(graphql.Float),
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve: resolvers.InsertHotel,
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
		},
	})
}
