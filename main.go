package main

import (
	"fmt"
	"strconv"

	"github.com/davidrbourke/Workout/domain/repos"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting...")

	exerciseRepo := repos.CreateExerciseRepo()
	exerciseRepo.IntialiseRepo()
	all, _ := exerciseRepo.FindAll()

	for _, e := range all {
		a := *e
		fmt.Println(a.PrintExercise())
	}

	r := gin.Default()
	r.GET("/exercises", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": all,
		})
	})

	r.GET("/exercises/:id", func(c *gin.Context) {
		exerciseID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("Invalid input")
			c.JSON(500, gin.H{
				"message": "Invalid id",
			})
			return
		}

		e, err := exerciseRepo.Get(exerciseID)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"data": e,
		})
	})

	r.Run()
}
