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
	case "5":
		day5(fileMap[day])
	case "6":
		day6(fileMap[day])
	case "7":
		day7(fileMap[day])
	case "8":
		day8(fileMap[day])
	case "10":
		day10(fileMap[day])
	}
}
