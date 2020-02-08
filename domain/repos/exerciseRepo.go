package repos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davidrbourke/Workout/domain/models"
)

// ExerciseRepo is the exercise repository
type ExerciseRepo interface {
	FindAll() ([]*models.Exercise, error)
	Get(exerciseID int) (*models.Exercise, error)
	IntialiseRepo()
}

type exerciseRepository struct {
	Exercises []*models.Exercise `json:"exercises"`
}

// CreateExerciseRepo return ref to repo
func CreateExerciseRepo() ExerciseRepo {
	return &exerciseRepository{}
}

func (r *exerciseRepository) IntialiseRepo() {
	jsonFile, err := os.Open("domain/jsonData/exercises.json")
	if err != nil {
		fmt.Println("Error reading json file")
		return
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &r)

	defer jsonFile.Close()
}

func (r *exerciseRepository) FindAll() ([]*models.Exercise, error) {

	return r.Exercises, nil
}

func (r *exerciseRepository) Get(exerciseID int) (*models.Exercise, error) {
	for _, e := range r.Exercises {
		if e.ExerciseID == exerciseID {
			return e, nil
		}
	}
	return nil, errors.New("Exercise not found")
}
