package queries

import (
	rBlog "../../resolvers/blog"
	rCar "../../resolvers/car"
	rCore "../../resolvers/core"
	rEvent "../../resolvers/event"
	rFlight "../../resolvers/flight"
	rHotel "../../resolvers/hotel"
	rTrain "../../resolvers/train"
	rUser "../../resolvers/user"
	tBlog "../types/blog"
	tCar "../types/car"
	tCore "../types/core"
	tEvent "../types/event"
	tFlight "../types/flight"
	tHotel "../types/hotel"
	tTrain "../types/train"
	tUser "../types/user"
	"github.com/graphql-go/graphql"
)

func Root() *graphql.Object {

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"SendSubscriptionToAll": {
				Type:        tCore.GetSubscriptionType(),
				Resolve:     rCore.SendSubscriptionToAll,
				Description: "Send Subscription",
			},
			"AllLocation": {
				Type:        graphql.NewList(tCore.GetLocationType()),
				Resolve:     rCore.AllLocation,
				Description: "Get All Location",
			},
			"GetCityByProvince": {
				Type: graphql.NewList(tCore.GetLocationType()),
				Args: graphql.FieldConfigArgument{
					"province": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rCore.GetCityByProvince,
				Description: "Get Location with Province",
			},

			"AllAdmin": {
				Type:        graphql.NewList(tUser.GetAdminType()),
				Resolve:     rUser.AllAdmin,
				Description: "Get All Admins",
			},
			"AdminLogin": {
				Type: tUser.GetJWTType(),
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rUser.LoginAdmin,
				Description: "Admin Login (Email, Password)",
			},
			"AllUser": {
				Type:        graphql.NewList(tUser.GetUserType()),
				Resolve:     rUser.AllUsers,
				Description: "Get All Users",
			},
			"UserByID": {
				Type: tUser.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rUser.GetUserByID,
				Description: "Get User By ID",
			},
			"UserByEmailAndPhone": {
				Type: tUser.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"emailphone": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rUser.GetUserByPhoneEmail,
				Description: "Get User By Email or Phone",
			},
			"UserLogin": {
				Type: tUser.GetJWTType(),
				Args: graphql.FieldConfigArgument{
					"emailphone": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rUser.UserLogin,
				Description: "User Login (Email, Phone, Password)",
			},
			"AllPhoneCode": {
				Type:        graphql.NewList(tUser.GetPhoneCodeType()),
				Resolve:     rUser.AllPhoneCode,
				Description: "Get All Phone Code",
			},
			"GetPhoneCode": {
				Type: tUser.GetPhoneCodeType(),
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rUser.GetPhoneCode,
				Description: "Get Phone Code",
			},

			"AllHotel": {
				Type:        graphql.NewList(tHotel.GetHotelType()),
				Resolve:     rHotel.AllHotel,
				Description: "Get All Hotel",
			},
			"NearestHotel": {
				Type: graphql.NewList(tHotel.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve:     rHotel.GetNearestHotel,
				Description: "Get Nearest Hotel",
			},
			"GetHotelByLocation": {
				Type: graphql.NewList(tHotel.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"province": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rHotel.GetHotelByProvince,
				Description: "Get Hotel By Province",
			},
			"GetHotelByLatLong": {
				Type: tHotel.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve:     rHotel.GetHotelByLatLong,
				Description: "Get Hotel By Latitude and Longitude",
			},
			"GetHotelByID": {
				Type: tHotel.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rHotel.GetHotelByID,
				Description: "Get Hotel By ID",
			},
			"AllHotelImage": {
				Type:        graphql.NewList(tHotel.GetHotelImageType()),
				Resolve:     rHotel.AllHotelImage,
				Description: "Get All Hotel Image",
			},
			"AllHotelFacility": {
				Type:        graphql.NewList(tHotel.GetHotelFacilityType()),
				Resolve:     rHotel.AllHotelFacility,
				Description: "Get All Hotel Facility",
			},
			"AllHotelType": {
				Type:        graphql.NewList(tHotel.GetHotelTypeType()),
				Resolve:     rHotel.AllHotelType,
				Description: "Get All Hotel Type",
			},

			"AllCarModel": {
				Type:        graphql.NewList(tCar.GetCarModelType()),
				Resolve:     rCar.AllCarModel,
				Description: "Get All Car Model",
			},
			"AllCar": {
				Type:        graphql.NewList(tCar.GetCarType()),
				Resolve:     rCar.AllCar,
				Description: "Get All Car",
			},
			"GetCarByCity": {
				Type: graphql.NewList(tCar.GetCarType()),
				Args: graphql.FieldConfigArgument{
					"city": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rCar.GetCarByCity,
				Description: "Get Car with City",
			},

			"AllTrain": {
				Type:        graphql.NewList(tTrain.GetTrainType()),
				Resolve:     rTrain.AllTrain,
				Description: "Get All Train",
			},
			"GetTrainByLocation": {
				Type: graphql.NewList(tTrain.GetTrainType()),
				Args: graphql.FieldConfigArgument{
					"date": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rTrain.GetTrainByArrivalDestination,
				Description: "Get Train By Arrival and Destination",
			},
			"AllTrainStation": {
				Type:        graphql.NewList(tTrain.GetTrainStationType()),
				Resolve:     rTrain.AllTrainStation,
				Description: "Get All Train Station",
			},

			"AllFlight": {
				Type:        graphql.NewList(tFlight.GetFlightType()),
				Resolve:     rFlight.AllFlight,
				Description: "Get All Flight",
			},
			"FlightByLocation": {
				Type:        graphql.NewList(tFlight.GetFlightType()),
				Args: graphql.FieldConfigArgument{
					"date": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"fromAirport": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"toAirport": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     rFlight.GetFlightByLocation,
				Description: "Get All Flight",
			},
			"AllFlightFacility": {
				Type:        graphql.NewList(tFlight.GetFlightFacilityType()),
				Resolve:     rFlight.AllFlightFacility,
				Description: "Get All Flight Facility",
			},
			"AllFlightCompany": {
				Type:        graphql.NewList(tFlight.GetFlightCompanyType()),
				Resolve:     rFlight.AllFlightCompany,
				Description: "Get All Flight Company",
			},
			"AllFlightAirport": {
				Type:        graphql.NewList(tFlight.GetFlightAirportType()),
				Resolve:     rFlight.AllFlightAirport,
				Description: "Get All Flight Airport",
			},

			"AllBlog": {
				Type:        graphql.NewList(tBlog.GetBlogType()),
				Resolve:     rBlog.AllBlog,
				Description: "Get All Blog",
			},

			"AllEntertainment": {
				Type:        graphql.NewList(tEvent.GetEntertainmentType()),
				Resolve:     rEvent.AllEntertainment,
				Description: "Get All Entertainment",
			},
		},
	})

}
