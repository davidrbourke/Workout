package models

import (
	"errors"
)

// Exercise struct
type Exercise struct {
	ExerciseID  int
	Name        string
	Description string
	MuscleGroup string
}

// CreateExercise constructor
func CreateExercise(exerciseID int, name, description, muscleGroup string) *Exercise {
	return &Exercise{
		ExerciseID:  exerciseID,
		Name:        name,
		Description: description,
		MuscleGroup: muscleGroup,
	}
}

// PrintExercise return user friendly output of the exercise
func (e Exercise) PrintExercise() string {
	return e.Name + ", " + e.MuscleGroup
}

// UpdateExerciseName updates the exercise name
func (e *Exercise) UpdateExerciseName(name string) error {
	if len(name) == 0 {
		return errors.New("Invalid exercise name")
	}
	e.Name = name
	return nil
}

func (e Exercise) setName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid name provided")
	}
	e.Name = value
	return nil
}
