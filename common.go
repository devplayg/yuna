package gofriend

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
