package config

import "insider_interview/config/helpers"

type Job struct {
	TimeIntervalInSeconds string
}

func NewJob() Job {
	return Job{
		TimeIntervalInSeconds: helpers.Getenv("TIME_INTERVAL_IN_SECONDS", "5"),
	}
}
