package utils

import "github.com/lib/pq"

const (
	PqErrCodeNameForeignKeyViolation = "foreign_key_violation"
	PqErrCodeNameUniqueViolation     = "unique_violation"
)

const (
	PqErrCodeForeignKeyViolation = "23503"
	PqErrCodeUniqueViolation     = "23505"
)

var PqErrUniqueViolation = &pq.Error{
	Code: PqErrCodeUniqueViolation,
}
