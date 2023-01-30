package internal

import "strconv"

func atoi(x string) (int, error) {
	y, err := strconv.ParseInt(x, 10, 32)
	return int(y), err
}
