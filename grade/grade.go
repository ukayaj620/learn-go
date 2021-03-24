package grade

func Grade(score int) string {
	if score >= 0 {
		if score < 55 {
			return "TIDAK LULUS"
		} else if score < 65 {
			return "D"
		} else if score < 75 {
			return "C"
		} else if score < 85 {
			return "B"
		} else {
			return "A"
		}
	} else {
		return "Nilai tidak bisa minus!"
	}
}
