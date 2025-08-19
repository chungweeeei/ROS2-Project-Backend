package main

import "time"

func convertStringToTime(s string) time.Time {

	layout := "2006-01-02"
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}
	}
	return t
}
