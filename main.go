package main

import (
	"fmt"

	"github.com/davidrbourke/Workout/domain/repos"
)

func main() {
	fmt.Println("Starting...")

	// exercise2 := exercise.CreateExercise(
	// 	"Squat",
	// 	"Steps, 1,2,3",
	// 	"Legs,Back",
	// )

	// fmt.Println(exercise1.PrintExercise())
	// fmt.Println(exercise2.PrintExercise())

	// exercise1.UpdateExerciseName("Deadlift 2")
	// fmt.Println(exercise1.PrintExercise())

	//controller.ConfigRouting(exercise1)

	repo := repos.CreateExerciseRepo()
	all, _ := repo.FindAll()
	// e := *(all[0])

	for _, e := range all {
		a := *e
		fmt.Println(a.PrintExercise())
	}
}