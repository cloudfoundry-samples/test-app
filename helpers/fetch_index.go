package helpers

import (
	"os"
	"strconv"
)

func FetchIndex() (int, error) {
	index := os.Getenv("CF_INSTANCE_INDEX")
	return strconv.Atoi(index)
}
