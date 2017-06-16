package repository

import (
	"errors"

	"github.com/eggsbenjamin/piemapping/models"
	"github.com/spf13/viper"
)

var (
	jErr = errors.New("unable to get journeys")
)

type JourneyRepositor interface {
	GetByDriverAvailability(string) ([]*models.Journey, error)
}

type JourneyRepository struct {
	conn Queryer
}

//	constructor
func NewJourneyRepository(conn Queryer) *JourneyRepository {
	return &JourneyRepository{
		conn: conn,
	}
}

//	queries the db by driverId and returns all journeys within which the driver is available for the duration.
//	returns a generic error 'unable to get journeys' if an error ocurs while querying.
func (j *JourneyRepository) GetByDriverAvailability(driverId string) ([]*models.Journey, error) {
	qry := viper.GetString("common.queries.journeysByDriverAvailability")
	rows, err := j.conn.Query(qry, driverId)
	if err != nil {
		return nil, jErr
	}
	defer rows.Close()
	jrnys := []*models.Journey{}
	for rows.Next() {
		j := &models.Journey{}
		if err := rows.Scan(
			&j.Id,
			&j.Departure_location,
			&j.Arrival_location,
			&j.Departure_time,
			&j.Arrival_time,
		); err != nil {
			return nil, jErr
		}
		jrnys = append(jrnys, j)
	}
	if rows.Err() != nil {
		return nil, jErr
	}
	return jrnys, nil
}
