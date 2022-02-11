package helper

import "log"

func CekConnectionDB(err error) {
	if err != nil {
		log.Fatal("Database connection failed")
	}
}
