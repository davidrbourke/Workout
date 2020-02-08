package repos

import (
	"errors"

	"github.com/davidrbourke/Workout/domain/models"
)

// ExerciseRepo is the exercise repository
type ExerciseRepo interface {
	FindAll() ([]*models.Exercise, error)
	Get(exerciseID int) (*models.Exercise, error)
	IntialiseRepo()
}

type exerciseRepository struct {
	exercises []*models.Exercise
}

// CreateExerciseRepo return ref to repo
func CreateExerciseRepo() ExerciseRepo {
	return &exerciseRepository{}
}

func (r *exerciseRepository) IntialiseRepo() {
	exercise1 := models.CreateExercise(
		1,
		"Deadlift",
		"Steps, 1,2,3",
		"Legs,Back",
	)

	exercise2 := models.CreateExercise(
		2,
		"Squat",
		"Steps, 1,2,3",
		"Legs,Back",
	)

	r.exercises = append(r.exercises, exercise1)
	r.exercises = append(r.exercises, exercise2)
}

func (r *exerciseRepository) FindAll() ([]*models.Exercise, error) {

	return r.exercises, nil
}

func (r *exerciseRepository) Get(exerciseID int) (*models.Exercise, error) {
	for _, e := range r.exercises {
		if e.ExerciseID == exerciseID {
			return e, nil
		}
	}
	return nil, errors.New("Exercise not found")
}
