package config

import (
	"os"
)

func EnvIsDev() bool {
	return os.Getenv("ENV") == "dev"
}

func EnvIsTest() bool {
	return os.Getenv("ENV") == "test"
}

func GetPodName() string {
	return os.Getenv("POD_NAME")
}

func IgnoreQuizJob() bool {
	return os.Getenv("IGNORE_QUIZ_JOB") == "1"
}
