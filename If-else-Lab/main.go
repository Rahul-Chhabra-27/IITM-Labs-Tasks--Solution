package main

import "fmt"

func gradeAssignment(score int) string {
	var grade string = ""
	
	if 90 <= score && score <= 100 {
		grade = "A";
	} else if 80 <= score && score < 90 {
		grade = "B";
	} else if 70 <= score && score < 80 {
		grade = "C";
	} else if 60 <= score && score  < 70 {
		grade = "D";
	} else if 0 <= score && score < 60 {
		grade = "F";
	} else {
		grade = "Grade can't be assign."
	}
	return grade
}
func main() {
	score := 61;
	fmt.Printf(gradeAssignment(score));
}

// 90 <= score <= 100: "A"
// 80 <= score < 90: "B"
// 70 <= score < 80: "C"
// 60 <= score < 70: "D"
// 0 <= score < 60: "F"
