package repository

import (
	"errors"
	"fmt"

	"github.com/eggsbenjamin/piemapping/commons"
	"github.com/eggsbenjamin/piemapping/models"
)

var (
	jErr = errors.New("unable to get journeys")
)

type JourneyRepositor interface {
	GetByDriverAvailability(string) ([]*models.Journey, error)
}

type JourneyRepository struct {
	conn Queryer
	log  commons.LevelledLogWriter
}

//	constructor
func NewJourneyRepository(conn Queryer, llw commons.LevelledLogWriter) *JourneyRepository {
	return &JourneyRepository{
		conn: conn,
		log:  llw,
	}
}

//	queries the db by driverId and returns all journeys within which the driver is available for the duration.
//	returns a generic error 'unable to get journeys' if an error ocurs while querying.
func (jr *JourneyRepository) GetByDriverAvailability(driverId string) ([]*models.Journey, error) {
	qry := fmt.Sprintf(`
		SELECT 
		j.id, j.departure_location, j.arrival_location,
		j.departure_time, j.arrival_time
		FROM drivers AS d 
		INNER JOIN journeys AS j 
		ON d.available_from <= j.departure_time 
		AND d.available_till >= j.arrival_time 
		WHERE d.id = '%s';
		`,
		driverId,
	)
	return jr.query(qry)
}

func (jr *JourneyRepository) query(qry string) ([]*models.Journey, error) {
	rows, err := jr.conn.Query(qry)
	if err != nil {
		jr.log.Errorf("Query error: %s", err.Error())
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
			jr.log.Errorf("Scan error: %s", err.Error())
			return nil, jErr
		}
		jrnys = append(jrnys, j)
	}
	if err := rows.Err(); err != nil {
		jr.log.Errorf("Row error: %v", err.Error())
		return nil, jErr
	}
	return jrnys, nil
}
