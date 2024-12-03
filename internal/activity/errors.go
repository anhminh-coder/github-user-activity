package activity

import "fmt"

var (
	ErrInternalServer = fmt.Errorf("internal server error")
	ErrUserNotFound   = fmt.Errorf("user not found")
)

func ErrFetchingData(status int) error {
	return fmt.Errorf("error fetching data: %d", status)
}
