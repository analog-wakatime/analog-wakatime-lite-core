package auths

import (
	"fmt"
	"strconv"
)

func parseTokenUserID(subject string) (uint, error) {
	userID, err := strconv.ParseUint(subject, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid token subject: %w", err)
	}

	return uint(userID), nil
}
