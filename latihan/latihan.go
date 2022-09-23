package latihan

import "fmt"

func ConditionChecking() {
	var currentyear = 2022

	if age := currentyear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat SIM")
	} else {
		fmt.Println("Kamu sudah boleh membuat SIM")
	}
}
