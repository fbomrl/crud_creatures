package database

import "fmt"

func BuildDSN(user, pass, host, instance, db string) string {
	if instance != "" {
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s?database=%s&instance=%s",
			user, pass, host, db, instance,
		)
	}
	return fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		user, pass, host, db,
	)
}
