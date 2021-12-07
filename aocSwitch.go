package main

func aocSwitch(day string, fileMap map[string]string) {
	switch day {
	case "1":
		day1(fileMap[day])
	case "2":
		day2(fileMap[day])
	case "3":
		day3(fileMap[day])
	case "4":
		day4(fileMap[day])
	}
}
