package main

import (
	"fmt"

	"github.com/ukayaj620/learn-go/cylinder"
	"github.com/ukayaj620/learn-go/grade"
	"github.com/ukayaj620/learn-go/loop_sum"
	"github.com/ukayaj620/learn-go/triangle"
)

func ChallengeOne() {
	fmt.Println("Challenge 01")

	cylinder := cylinder.Cylinder{}
	fmt.Print("Input Tinggi Tabung: ")
	fmt.Scan(&cylinder.Height)

	fmt.Print("Input Radius Tabung: ")
	fmt.Scan(&cylinder.Radius)

	fmt.Printf("Luas Permukaan Tabung: %.2f\n", cylinder.SurfaceArea())
	fmt.Printf("Volume Tabung: %.2f\n", cylinder.Volume())
}

func ChallengeTwo() {
	fmt.Println("\nChallenge 02")
	var score int
	fmt.Print("Input Score: ")
	fmt.Scan(&score)
	fmt.Printf("Sum: %s\n", grade.Grade(score))
}

func ChallengeThree() {
	fmt.Println("\nChallenge 03")
	var n int
	fmt.Print("Input N: ")
	fmt.Scan(&n)
	fmt.Printf("Sum: %d\n", loop_sum.LoopSum(n))
}

func ChallengeFour() {
	fmt.Println("\nChallenge 04")
	var height int
	fmt.Print("Tinggi segitiga: ")
	fmt.Scan(&height)
	triangle.LowerRight(height)
	fmt.Println()
}

func main() {
	ChallengeOne()
	ChallengeTwo()
	ChallengeThree()
	ChallengeFour()
}
