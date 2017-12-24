package helpers

import (
	"log"
	"strconv"
	"strings"
)

// RepresentShelflife Convert shelflife to number of days
func RepresentShelflife(data string) (result int) {
	split := strings.Split(data, " ")
	mod := split[1]
	var value int
	var err error
	if value, err = strconv.Atoi(split[0]); err != nil {
		log.Println(err)
		return result
	}

	switch mod {
	case "дней":
		result = value
		break
	case "день":
		result = value
		break
	case "год":
		result = 365 * value
		break
	case "месяц":
		result = int(float32(365) / float32(12) * float32(value))
		break
	case "месяцев":
		result = int(float32(365) / float32(12) * float32(value))
		break
	}
	return result
}
