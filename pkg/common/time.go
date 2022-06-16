package common

import "time"

func GetParsedDate(
	dateString string,
	timeZone string,
) (*time.Time, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, err
	}
	date, err := time.ParseInLocation("2006-01-02", dateString, loc)
	if err != nil {
		return nil, err
	}
	return &date, nil
}
