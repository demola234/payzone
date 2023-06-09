package gapi

import "google.golang.org/genproto/googleapis/rpc/errdetails"

func fieldViolation(field string, erro string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: description,
	}
}
