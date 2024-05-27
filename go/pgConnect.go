package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Parent struct {
	val1 int
	val2 int
}

func (parent Parent) Adder() int {
	return parent.val1 + parent.val2
}

type Child struct {
	Parent
	val3 string
}

type Movie struct {
	gorm.Model
	Name        string
	Stars       string
	Description string
}

func main() {
	// Simple Composition Test
	compositionTest()
	// Functions inside Functions
	// Function as argument
	fmt.Println(functionAsArgument(true, func(flag bool) string {
		if flag {
			return "true"
		} else {
			return "false"
		}
	}))
	// Return Function
	fmt.Println(createIncrementer(5)(7))
	//Connecting To Database
	db := connectToDB()
	//Creating Schema If Needed
	db.AutoMigrate(&Movie{})
	//Insert Some Value
	db.Create(&Movie{Name: "RRR", Stars: "4.8", Description: "Action Epic Historical"})
	//Retrieve Some Value
	movie := retrieveTest(db)
	//Update retrieved value (other attributes remain same only Stars changes)
	db.Model(&movie).Updates(Movie{Stars: "4.1"})
	//Delete the movie
	db.Delete(&movie)
	fmt.Println("DONE")
}

func retrieveTest(db *gorm.DB) Movie {
	// specify variable to receive input
	var movie Movie
	// retrieve first value based on name, (assume ordering by primary key)
	db.First(&movie, "name = ?", "RRR")
	fmt.Printf("%+v", movie)
	return movie
}

func connectToDB() *gorm.DB {
	dsn := "host=localhost user=shashank password=noob dbname=movie port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db
}

func compositionTest() {
	var parent Parent = Parent{1, 2}
	child := Child{Parent: parent, val3: ""}
	fmt.Println(child.Adder())
}

func functionAsArgument(flag bool, compute func(input bool) string) string {
	return compute(flag)
}

func createIncrementer(baseInput int) func(input int) int {
	return func(input int) int {
		return baseInput + input
	}
}
