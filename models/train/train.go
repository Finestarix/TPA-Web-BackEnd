package train

import (
	"../../connection"
	"log"
	"time"
)

type Train struct {
	ID            int          `gorm:"PRIMARY_KEY"`
	Name          string       `gorm:"VARCHAR(100); NOT NULL"`
	Code          string       `gorm:"VARCHAR(100); NOT NULL"`
	Class         string       `gorm:"VARCHAR(100); "`
	Arrival       TrainStation `gorm:"FOREIGNKEY:ArrivalID"`
	ArrivalID     int          `gorm:"INTEGER; NOT NULL"`
	ArrivalTime   time.Time
	Transit       TrainStation `gorm:"FOREIGNKEY:TransitID"`
	TransitID     int          `gorm:"INTEGER; NOT NULL"`
	Departure     TrainStation `gorm:"FOREIGNKEY:DepartmentID"`
	DepartureID   int          `gorm:"INTEGER; NOT NULL"`
	DepartureTime time.Time
	Seat          int `gorm:"INTEGER; NOT NULL"`
	Price         int `gorm:"INTEGER; NOT NULL"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:index`
}

func init() {
	database := connection.GetConnection()
	defer database.Close()

	database.
		AutoMigrate(&Train{}).
		AddForeignKey("departure_id", "train_stations(id)", "CASCADE", "CASCADE").
		AddForeignKey("transit_id", "train_stations(id)", "CASCADE", "CASCADE").
		AddForeignKey("arrival_id", "train_stations(id)", "CASCADE", "CASCADE")

	log.Println("Initialize Train Success")
}

func GetAllTrain() []Train {
	database := connection.GetConnection()
	defer database.Close()

	var train []Train
	database.Find(&train)

	for i, _ := range train {
		database.Model(train[i]).Related(&train[i].Departure, "departure_id")
		database.Model(train[i]).Related(&train[i].Transit, "transit_id")
		database.Model(train[i]).Related(&train[i].Arrival, "arrival_id")
	}

	return train
}

func GetTrainByArrivalDestination(arrival string, destination string, date time.Time) []Train {
	database := connection.GetConnection()
	defer database.Close()

	arrivalStation := SearchTrainStationByName(arrival)
	departureStation := SearchTrainStationByName(destination)


	var train []Train
	database.Where("arrival_id = ? AND departure_id = ? AND DATE_PART('day',arrival_time) = ?",
		arrivalStation.ID, departureStation.ID, date.Day()).Find(&train)

	for i, _ := range train {
		database.Model(train[i]).Related(&train[i].Departure, "departure_id")
		database.Model(train[i]).Related(&train[i].Transit, "transit_id")
		database.Model(train[i]).Related(&train[i].Arrival, "arrival_id")
	}

	return train
}

func InsertTrain(name string, code string, class string, arrival string, arrivalTime time.Time, transit string, departure string, departureTime time.Time, seat int, price int) *Train {
	database := connection.GetConnection()
	defer database.Close()

	arrivalStation := SearchTrainStationByName(arrival)
	departureStation := SearchTrainStationByName(departure)

	var transitStation TrainStation
	var transitID = 1
	if len(transit) != 0 {
		transitStation = SearchTrainStationByName(transit)
		transitID = transitStation.ID
	}

	newTrain := &Train{
		Name:          name,
		Code:          code,
		Class:         class,
		ArrivalID:     arrivalStation.ID,
		ArrivalTime:   arrivalTime.Add(-7 * time.Hour),
		TransitID:     transitID,
		DepartureID:   departureStation.ID,
		DepartureTime: departureTime.Add(-7 * time.Hour),
		Seat:          seat,
		Price:         price,
	}
	database.Save(newTrain)

	database.Model(newTrain).Related(&newTrain.Departure, "departure_id")
	database.Model(newTrain).Related(&newTrain.Transit, "transit_id")
	database.Model(newTrain).Related(&newTrain.Arrival, "arrival_id")

	log.Println("Insert New Train Success")
	return newTrain
}

func GetTrainByID(id int) Train {
	database := connection.GetConnection()
	defer database.Close()

	var train Train
	database.
		Where("id = ?", id).
		First(&train)

	database.Model(train).Related(&train.Departure, "departure_id")
	database.Model(train).Related(&train.Transit, "transit_id")
	database.Model(train).Related(&train.Arrival, "arrival_id")

	return train
}

func UpdateTrain(id int, arrivalTime time.Time, departureTime time.Time, seat int, price int) Train {
	database := connection.GetConnection()
	defer database.Close()

	var train Train
	database.
		Model(&train).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"arrival_time":   arrivalTime.Add(-7 * time.Hour),
			"departure_time": departureTime.Add(-7 * time.Hour),
			"seat":           seat,
			"price":          price,
		})
	var newTrain Train
	database.
		Where("id = ?", id).
		Find(&newTrain)

	database.Model(newTrain).Related(&newTrain.Departure, "departure_id")
	database.Model(newTrain).Related(&newTrain.Transit, "transit_id")
	database.Model(newTrain).Related(&newTrain.Arrival, "arrival_id")

	log.Println("Update Train Success")
	return newTrain
}

func DeleteTrain(id int) *Train {
	database := connection.GetConnection()
	defer database.Close()

	var train Train
	train = GetTrainByID(id)

	if train.ID == 0 {
		log.Println("Delete User Failed")
		return &train
	}

	err := database.Delete(train).Error

	if err != nil {
		panic("Error Delete Train !" + err.Error())
	}

	log.Println("Delete User Success")
	return &train
}
