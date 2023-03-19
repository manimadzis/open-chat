package repositories

import "fmt"

func UnknownError(err error) error {
	return fmt.Errorf("unknown error: %v", err.Error())
}
