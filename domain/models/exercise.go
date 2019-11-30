package models

import (
	"errors"
)

type exercise struct {
	name        string
	description string
	muscleGroup string
}

// Header return heading for exercise
type Header interface {
	PrintExercise() string
	UpdateExerciseName(nam string) error
}

// CreateExercise constructor
func CreateExercise(name, description, muscleGroup string) Header {
	return &exercise{
		name:        name,
		description: description,
		muscleGroup: muscleGroup,
	}
}

func (e exercise) PrintExercise() string {
	return e.name + ", " + e.muscleGroup
}

func (e *exercise) UpdateExerciseName(name string) error {
if len(name) == 0 {
		return errors.New("Invalid exercise name")
	}
	e.name = name
	return nil
}

func (e exercise) setName(value string) error {
if len(value) == 0 {
		return errors.New("Invalid name provided")
	}
	return nil
}
