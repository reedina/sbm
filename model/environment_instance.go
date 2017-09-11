package model

import (
	"time"

	"github.com/araddon/dateparse"
)

//EnvironmentInstance  (TYPE)
type EnvironmentInstance struct {
	ID                   int         `json:"id"`
	Environment          Environment `json:"environment"`
	ExpirationStringTime string      `json:"expiration_time"`      // "2017-12-01 22:43:22" - Local time
	ExpirationUtcEpoch   int64       `json:"expiration_utc_epoch"` // 1513434000 - UTC time
}

/*
{
   "environment":  {
     "name": "test01",
     "type": "qa"
  },
   "expiration_time":  "2017-12-01 22:43:22"
}
*/

//CreateEnvironmentInstance - Store in database
func CreateEnvironmentInstance(environmentInstance *EnvironmentInstance) error {
	err := db.QueryRow(
		"INSERT INTO sbm_environment_instances(environment_id, expiration_string, expiration_time"+
			") VALUES($1,$2,$3) RETURNING id", environmentInstance.Environment.ID, environmentInstance.ExpirationStringTime,
		environmentInstance.ExpirationUtcEpoch).Scan(&environmentInstance.ID)

	if err != nil {
		return err
	}

	return nil
}

//ConvertStringToUtcEpoch -
func ConvertStringToUtcEpoch(datetime string) (int64, error) {
	timeValue, err := convertStringToTime(datetime)

	if err != nil {
		return 0, err
	}

	return convertLocaltimetoUtcEpoch(timeValue), nil
}

//IsUtcEpoch24HoursFromNow -
func IsUtcEpoch24HoursFromNow(utcEpoch int64) bool {
	plus24Hours := time.Now().UTC().Unix() + (60 * 60 * 24)

	return utcEpoch > plus24Hours
}

//Internal Function - convertStringToTime
func convertStringToTime(datetime string) (time.Time, error) {
	t, err := dateparse.ParseLocal(datetime)
	if err != nil {
		return time.Time{}, err // time.Time{} is the empty value for time.Time
	}
	return t, nil
}

//Internal Function - convertLocaltimetoUtcEpoch
func convertLocaltimetoUtcEpoch(localtime time.Time) int64 {
	return localtime.UTC().Unix()

}