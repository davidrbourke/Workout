package repos

import (
	"github.com/davidrbourke/Workout/domain/models"
)

// Repo is the exercise repository
type Repo interface {
	FindAll() ([]*models.Header, error)
}

type exerciseRepository struct {
	exercises []*models.Header
}

// CreateExerciseRepo return ref to repo
func CreateExerciseRepo() Repo {
	return &exerciseRepository{}
}

func (r *exerciseRepository) FindAll() ([]*models.Header, error) {

	exercise1 := models.CreateExercise(
		"Deadlift",
		"Steps, 1,2,3",
		"Legs,Back",
	)

	exercise2 := models.CreateExercise(
		"Squat",
		"Steps, 1,2,3",
		"Legs,Back",
	)

	r.exercises = append(r.exercises, &exercise1)
	r.exercises = append(r.exercises, &exercise2)
	return r.exercises, nil
}
