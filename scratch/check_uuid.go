package main

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	var u pgtype.UUID
	copy(u.Bytes[:], "1234567890123456")
	u.Valid = true
	fmt.Println(u.String())
}
