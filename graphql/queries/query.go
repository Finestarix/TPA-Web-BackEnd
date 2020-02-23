package queries

import (
	"../../resolvers"
	"../types"
	"github.com/graphql-go/graphql"
)

func Root() *graphql.Object {

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"SendSubscriptionToAll": {
				Type:        types.GetSubscriptionType(),
				Resolve:     resolvers.SendSubscriptionToAll,
				Description: "Send Subscription",
			},

			"AllAdmin": {
				Type:        graphql.NewList(types.GetAdminType()),
				Resolve:     resolvers.AllAdmin,
				Description: "Get All Admins",
			},
			"AdminLogin": {
				Type: types.GetJWTType(),
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.LoginAdmin,
				Description: "Admin Login (Email, Password)",
			},

			"AllUser": {
				Type:        graphql.NewList(types.GetUserType()),
				Resolve:     resolvers.AllUsers,
				Description: "Get All Users",
			},
			"UserByID": {
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetUserByID,
				Description: "Get User By ID",
			},
			"UserByEmailAndPhone": {
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"emailphone": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetUserByPhoneEmail,
				Description: "Get User By Email or Phone",
			},
			"UserLogin": {
				Type: types.GetJWTType(),
				Args: graphql.FieldConfigArgument{
					"emailphone": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.UserLogin,
				Description: "User Login (Email, Phone, Password)",
			},

			"AllPhoneCode": {
				Type:        graphql.NewList(types.GetPhoneCodeType()),
				Resolve:     resolvers.AllPhoneCode,
				Description: "Get All Phone Code",
			},
			"GetPhoneCode": {
				Type: types.GetPhoneCodeType(),
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetPhoneCode,
				Description: "Get Phone Code",
			},

			"AllHotel": {
				Type:        graphql.NewList(types.GetHotelType()),
				Resolve:     resolvers.AllHotel,
				Description: "Get All Hotel",
			},
			"NearestHotel": {
				Type: graphql.NewList(types.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve:     resolvers.GetNearestHotel,
				Description: "Get Nearest Hotel",
			},
			"GetHotelByLocation": {
				Type: graphql.NewList(types.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"province": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetHotelByProvince,
				Description: "Get Hotel By Province",
			},
			"GetHotelByLatLong": {
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve:     resolvers.GetHotelByLatLong,
				Description: "Get Hotel By Latitude and Longitude",
			},
			"GetHotelByID": {
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetHotelByID,
				Description: "Get Hotel By ID",
			},

			"AllLocation": {
				Type:        graphql.NewList(types.GetLocationType()),
				Resolve:     resolvers.AllLocation,
				Description: "Get All Location",
			},
			"GetCityByProvince": {
				Type: graphql.NewList(types.GetLocationType()),
				Args: graphql.FieldConfigArgument{
					"province": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetCityByProvince,
				Description: "Get Location with Province",
			},

			"AllHotelImage": {
				Type:        graphql.NewList(types.GetHotelImageType()),
				Resolve:     resolvers.AllHotelImage,
				Description: "Get All Hotel Image",
			},

			"AllHotelFacility": {
				Type:        graphql.NewList(types.GetHotelFacilityType()),
				Resolve:     resolvers.AllHotelFacility,
				Description: "Get All Hotel Facility",
			},

			"AllHotelType": {
				Type:        graphql.NewList(types.GetHotelTypeType()),
				Resolve:     resolvers.AllHotelType,
				Description: "Get All Hotel Type",
			},

			"AllCarModel": {
				Type:        graphql.NewList(types.GetCarModelType()),
				Resolve:     resolvers.AllCarModel,
				Description: "Get All Car Model",
			},

			"AllCar": {
				Type:        graphql.NewList(types.GetCarType()),
				Resolve:     resolvers.AllCar,
				Description: "Get All Car",
			},
			"GetCarByCity": {
				Type: graphql.NewList(types.GetCarType()),
				Args: graphql.FieldConfigArgument{
					"city": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolvers.GetCarByCity,
				Description: "Get Car with City",
			},
		},
	})

}
