package stringutils

import (
	"errors"
	"strconv"
	"strings"
	"time"

	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func StrToInt64(str string) (int64, error) {
	if strings.TrimSpace(str) == "" {
		err := errors.New("invalid number")
		return 0, err
	}

	number, getErr := strconv.ParseInt(str, 10, 64)
	if getErr != nil {
		err := errors.New("invalid number")
		return 0, err
	}
	return number, nil
}

func NullStringToString(ns adm.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func NullTimeToTime(ns adm.NullTime) time.Time {
	if ns.Valid {
		return ns.Time
	}
	return time.Time{}
}

// NullInt32ToInt converts sql.NullInt32 to int32. Returns 0 if invalid.
func NullInt32ToInt(ni adm.NullInt32) int32 {
	if ni.Valid {
		return ni.Int32
	}
	return 0
}
