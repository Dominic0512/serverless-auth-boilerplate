package domain

import (
	"fmt"
)

type DomainError interface {
	error
	DomainError()
}

type EntityCreationError struct {
	Entity string
}

func (e *EntityCreationError) Error() string {
	return fmt.Sprintf("Failed to create %s", e.Entity)
}

func (e *EntityCreationError) DomainError() {}

type NotFoundError struct {
	Entity string
	ID     string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %s not found", e.Entity, e.ID)
}

func (e *NotFoundError) DomainError() {}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field %s: %s", e.Field, e.Message)
}

func (e *ValidationError) DomainError() {}

type InvalidError struct {
	Entity  string
	Message string
}

func (e *InvalidError) Error() string {
	return fmt.Sprintf("Invalid %s: %s", e.Entity, e.Message)
}

func (e *InvalidError) DomainError() {}
