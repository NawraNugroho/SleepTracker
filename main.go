//This application helps users record and analyze their sleep patterns and daily health
//activities. The primary data includes sleep history, sleep and wake-up times, and
//sleep quality. The intended users are individuals who want to monitor their sleep
//health

//a. Users can add, modify, and delete sleep records, including sleep and wake-up times.
//b. The system calculates sleep duration and provides sleep improvement suggestions.
//c. Users can search for sleep records by date using Sequential and Binary Search.
//d. Users can sort sleep records by duration or date using Selection and Insertion Sort.
//e. The system displays sleep reports, including, for example:
//• A summary of sleep duration over the past 7 days
//• The average sleep duration per week.

package main

import (
	"fmt"
)

const NMAX int = 1000

type myTime struct {
	hour     int
	minute   int
	totalMin int
}
type sleepData struct {
	date                string
	sleepTime, wakeTime myTime
	sleepDuration       myTime
	sleepQuality        int
}
type record [NMAX]sleepData

func main() {
	var sleepRecord record
	var addN int
	var n int = 0
	var opt int
	var ans int
	var date string
	var idx int

	
	opt = -1
	for opt != 0 {
		fmt.Println("\n--- Sleep Tracker ---")
		fmt.Println("1. Add Record")
		fmt.Println("2. Modify Record")
		fmt.Println("3. Delete Record")
		fmt.Println("4. Search Record by Date")
		fmt.Println("5. View Records")
		fmt.Println("6. View Weekly Summary")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scan(&opt)

		switch opt {
		case 1:
			fmt.Println("\n--- Add Record ---")
			fmt.Print("How many records would you like to add: ")
			fmt.Scan(&addN)
			addRecord(&sleepRecord, addN, &n)
		case 2:
			fmt.Println("\n--- Modify Record ---")
			fmt.Print("Enter the date you want to modify (yy/mm/dd) or x to exit: ")
			fmt.Scan(&date)
			for date != "x" {
				modify(&sleepRecord, n, date)
				fmt.Print("Enter the date you want to modify (yy/mm/dd) or x to exit: ")
				fmt.Scan(&date)
			}
		case 3:
			fmt.Println("\n--- Delete Record ---")
			fmt.Print("Enter the date you want to delete (yy/mm/dd) or x to exit: ")
			fmt.Scan(&date)
			for date != "x" {
				delete(&sleepRecord, &n, date)
				fmt.Print("Enter the date you want to delete (yy/mm/dd) or x to exit: ")
				fmt.Scan(&date)
			}
		case 4:
			fmt.Println("\n--- Search records by date ---")
			insertionSortDate(&sleepRecord, n)
			fmt.Print("Enter the date you want to search (yy/mm/dd) or x to exit: ")
			fmt.Scan(&date)
			for date != "x" {
				idx = binSearch(sleepRecord, n, date)
				if idx != -1 {
					fmt.Printf("\n|%-15s | %-10s | %-12s | %-15s | %-13s|\n", "Date", "Sleep Time", "Wake Up Time", "Sleep Duration", "Sleep Quality")
					fmt.Println("-------------------------------------------------------------------------------")
					displayARecord(sleepRecord, n, idx)
				} else {
					fmt.Println("Date is not found!")
				}
				fmt.Print("\nEnter the date you want to search (yy/mm/dd) or x to exit: ")
				fmt.Scan(&date)
			}
		case 5:
			fmt.Println("\n--- View records ---")
			ans = -1
			for ans != 0 {
				fmt.Println("1. Sort by date")
				fmt.Println("2. Sort by sleeping duration")
				fmt.Println("0. Exit")
				fmt.Print("Choose an option: ")
				fmt.Scan(&ans)

				switch ans {
				case 1:
					sortDate(sleepRecord, n)
				case 2:
					sortDuration(sleepRecord, n)
				}
			}
		case 6:
			fmt.Println("\n--- View Weekly Summary ---")
			if n < 7 {
				fmt.Println("You don't have enough record. You must have at least 7 records to use this feature.")
			} else {
				fmt.Println("The 7 most recent recorded days.")
				viewWeeklySummary(sleepRecord, n)
			}
		}
	}
}

func viewWeeklySummary(A record, n int) {
	var totSleepDur, totSleepQua int
	var avgSleepDurInHour, avgSleepQua int
	var ans string = "0"
	var improveSleepQuality int
	
	selectionSortDate(&A, n)
	
	fmt.Printf("\n|%-15s | %-10s | %-12s | %-15s | %-13s|\n", "Date", "Sleep Time", "Wake Up Time", "Sleep Duration", "Sleep Quality")
	fmt.Println("-------------------------------------------------------------------------------")
	for i := 0; i < 7; i++ {
		displayARecord(A, n, i)
		totSleepDur += A[i].sleepDuration.totalMin
		totSleepQua += A[i].sleepQuality
	}
	
	avgSleepDurInHour = (totSleepDur/7)/60
	avgSleepQua = totSleepQua/7
	
	fmt.Println("\nYour average sleep duration is ", avgSleepDurInHour)
	fmt.Println("Your average sleep quality is ", avgSleepQua)
	improveSleepQuality = avgSleepQua + 2
	
	for ans != "yes" && ans != "no" && improveSleepQuality < 11{
		fmt.Print("\nDo you want to improve your sleep quality (yes/no)? ")
		fmt.Scan(&ans)
	}
	
	if ans == "yes" {
		fmt.Printf("You can improve your sleep quality rate to %d or %d", improveSleepQuality-1, improveSleepQuality)
		fmt.Println("\n\nIn order to gain that, you have to start: ")
		switch improveSleepQuality{
		case 3 :
			fmt.Println("1. Sleep before 23:00")
			fmt.Println("2. Maintain your average sleeping duration")
		case 4:
			fmt.Println("1. Maintain your sleeping time")
			fmt.Println("2. Increase your average sleeping duration to at least 10 hours")
		case 5:
			fmt.Println("1. Maintain your sleeping time")
			fmt.Println("2. Increase your average sleeping duration to at least 10 hours")	
		case 6:
			fmt.Println("1. Maintain your sleeping time")
			fmt.Println("2. Decrease your average sleeping duration to 6-7 hours")	
		case 7:
			fmt.Println("1. Maintain your sleeping time")
			fmt.Println("2. Decrease your average sleeping duration to 6-7 hours")	
		case 8:
			fmt.Println("1. Maintain your sleeeping time")
			fmt.Println("2. Increase your average sleeping duration to 7-9 hours")	
		case 9:
			fmt.Println("1. Sleep before 21:00")
			fmt.Println("2. Increase your average sleeping duration to 7-9 hours")		
		case 10:
			fmt.Println("1. Sleep before 23:00")
			fmt.Println("2. Maintain your average sleeping duration")
		}
	} 
}
func calcSleepQuality(A *record, idx int) {
	switch {
	// Optimal sleep duration: 7-9 hours
	case A[idx].sleepDuration.hour >= 7 && A[idx].sleepDuration.hour <= 9 && A[idx].sleepTime.hour >= 21 && A[idx].sleepTime.hour < 23:
		A[idx].sleepQuality = 10
	case A[idx].sleepDuration.hour >= 7 && A[idx].sleepDuration.hour <= 9 && A[idx].sleepTime.hour < 21 && A[idx].sleepTime.hour >= 18:
		A[idx].sleepQuality = 9
	case A[idx].sleepDuration.hour >= 7 && A[idx].sleepDuration.hour <= 9 && (A[idx].sleepTime.hour == 23 || A[idx].sleepTime.hour >= 0):
		A[idx].sleepQuality = 8

	// Suboptimal duration: 6-7 hours
	case A[idx].sleepDuration.hour >= 6 && A[idx].sleepDuration.hour < 7 && A[idx].sleepTime.hour >= 21 && A[idx].sleepTime.hour < 23:
		A[idx].sleepQuality = 7
	case A[idx].sleepDuration.hour >= 6 && A[idx].sleepDuration.hour < 7 && (A[idx].sleepTime.hour >= 23 || A[idx].sleepTime.hour < 21):
		A[idx].sleepQuality = 6

	// Oversleeping >9 hours
	case A[idx].sleepDuration.hour > 9 && A[idx].sleepTime.hour >= 21 && A[idx].sleepTime.hour < 23:
		A[idx].sleepQuality = 5
	case A[idx].sleepDuration.hour > 9 && (A[idx].sleepTime.hour >= 23 || A[idx].sleepTime.hour < 21):
		A[idx].sleepQuality = 4

	// Very short sleep <6 hours
	case A[idx].sleepDuration.hour < 6 && A[idx].sleepTime.hour >= 21 && A[idx].sleepTime.hour < 23:
		A[idx].sleepQuality = 3
	case A[idx].sleepDuration.hour < 6 && A[idx].sleepTime.hour < 21 && A[idx].sleepTime.hour >= 18:
		A[idx].sleepQuality = 2
	case A[idx].sleepDuration.hour < 6 && (A[idx].sleepTime.hour == 23 || A[idx].sleepTime.hour >= 0):
		A[idx].sleepQuality = 1

	default:
		A[idx].sleepQuality = 0
	}
}

func sortDate(A record, n int) {
	var ans int

	fmt.Println("\n--- Sort by Date ---")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Choose an option: ")
	fmt.Scan(&ans)

	switch ans {
	case 1:
		fmt.Print("\nSort by date in ascending order: ")
		insertionSortDate(&A, n)
		displayAll(A, n)
		fmt.Println()
	case 2:
		fmt.Print("\nSort by date in descending order: ")
		selectionSortDate(&A, n)
		displayAll(A, n)
		fmt.Println()
	}
}

func insertionSortDate(A *record, n int) {
	//sort by date in ascending order
	var i, pass int
	var temp sleepData

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.date < A[i-1].date {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}

func selectionSortDate(A *record, n int) {
	//sort by date in descending order
	var i, idx, pass int
	var temp sleepData

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].date > A[idx].date {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}
func sortDuration(A record, n int) {
	var ans int

	fmt.Println("\n--- Sort by Sleeping Duration ---")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Choose an option: ")
	fmt.Scan(&ans)

	switch ans {
	case 1:
		fmt.Print("\nSort by sleeping duration in ascending order: ")
		insertionSortDuration(&A, n)
		displayAll(A, n)
		fmt.Println()
	case 2:
		fmt.Print("\nSort by sleeping duration in descending order: ")
		selectionSortDuration(&A, n)
		displayAll(A, n)
		fmt.Println()
	}
}

func addRecord(A *record, addN int, n *int) {
	var sleepHour, sleepMinute int
	var wakeHour, wakeMinute int
	var i int
	var max int

	max = *n
	if max == 0 {
		i = 0
	} else {
		i = max
	}

	for i < addN+max {
		fmt.Print("Date (yy/mm/dd): ")
		fmt.Scan(&A[i].date)
		fmt.Print("Sleep Time in 24-hour format (h m): ")
		fmt.Scan(&sleepHour, &sleepMinute)
		A[i].sleepTime.hour = sleepHour
		A[i].sleepTime.minute = sleepMinute
		fmt.Print("Wake Up Time in 24-hour format (h m): ")
		fmt.Scan(&wakeHour, &wakeMinute)
		A[i].wakeTime.hour = wakeHour
		A[i].wakeTime.minute = wakeMinute
		fmt.Println()
		A[i].sleepDuration = calcSleepDur(sleepHour, sleepMinute, wakeHour, wakeMinute)
		calcSleepQuality(&*A, i)
		i++
		*n = *n + 1
	}
}

func calcSleepDur(sleepHour int, sleepMinute int, wakeHour int, wakeMinute int) myTime {
	var duration myTime

	sleepTotMins := sleepHour*60 + sleepMinute
	wakeTotMins := wakeHour*60 + wakeMinute

	duration.totalMin = wakeTotMins - sleepTotMins

	if duration.totalMin < 0 {
		duration.totalMin = 1440 + duration.totalMin
	}

	duration.hour = duration.totalMin / 60
	duration.minute = duration.totalMin % 60

	return duration
}

func searchSeq(A record, n int, date string) int {
	var i, idx int

	idx = -1
	for A[i].date != date && i < n {
		i++
	}

	if A[i].date == date {
		idx = i
	}

	return idx
}

func binSearch(A record, n int, x string) int {
	var left, right, mid int
	
	left = 0
	right = n-1
	mid = (left+right)/2
	for left <= right {
		if A[mid].date == x {
			return mid
		}else if A[mid].date < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left+right)/2
	}
	return -1
}
func selectionSortDuration(A *record, n int) {
	//sort by sleep duration in descending order
	var i, idx, pass int
	var temp sleepData

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].sleepDuration.totalMin > A[idx].sleepDuration.totalMin {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func insertionSortDuration(A *record, n int) {
	//sort by sleep duration in ascending order
	var i, pass int
	var temp sleepData

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.sleepDuration.totalMin < A[i-1].sleepDuration.totalMin {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}

func modify(A *record, n int, date string) {
	var idx int
	var ans int
	var sleepHour, sleepMinute, wakeHour, wakeMinute int

	idx = searchSeq(*A, n, date)
	if idx != -1 {
		ans = -1
		for ans != 0 {
			fmt.Printf("\n|%-15s | %-10s | %-12s | %-15s | %-13s|\n", "Date", "Sleep Time", "Wake Up Time", "Sleep Duration", "Sleep Quality")
			fmt.Println("-------------------------------------------------------------------------------")
			displayARecord(*A, n, idx)
			fmt.Println("\nWhich attribute would you like to edit? ")
			fmt.Println("1. Date")
			fmt.Println("2. Sleep time")
			fmt.Println("3. Wake up time")
			fmt.Println("0. Exit")
			fmt.Print("Choose an option: ")
			fmt.Scan(&ans)

			switch ans {
			case 1:
				fmt.Print("\nDate (yy/mm/dd): ")
				fmt.Scan(&A[idx].date)
			case 2:
				fmt.Print("\nSleep Time in 24-hour format (h m): ")
				fmt.Scan(&sleepHour, &sleepMinute)
				A[idx].sleepTime.hour = sleepHour
				A[idx].sleepTime.minute = sleepMinute
				wakeHour = A[idx].wakeTime.hour
				wakeMinute = A[idx].wakeTime.minute
				A[idx].sleepDuration = calcSleepDur(sleepHour, sleepMinute, wakeHour, wakeMinute)
				calcSleepQuality(&*A, idx)
			case 3:
				fmt.Print("\nWake Up Time in 24-hour format (h m): ")
				fmt.Scan(&wakeHour, &wakeMinute)
				A[idx].wakeTime.hour = wakeHour
				A[idx].wakeTime.minute = wakeMinute
				sleepHour = A[idx].sleepTime.hour
				sleepMinute = A[idx].sleepTime.minute
				A[idx].sleepDuration = calcSleepDur(sleepHour, sleepMinute, wakeHour, wakeMinute)
				calcSleepQuality(&*A, idx)

			}
		}
	} else {
		fmt.Println("The record is not found!")
	}
}

func delete(A *record, n *int, date string) {
	var idx int
	var i int

	idx = searchSeq(*A, *n, date)

	if idx != -1 {
		i = idx
		for i < *n-1 {
			A[i] = A[i+1]
			i++
		}
		*n = *n -1
		fmt.Println("Record deleted successfully!")
	} else {
		fmt.Println("The record is not found!")
	}
}
func displayAll(A record, n int) {
	fmt.Printf("\n|%-15s | %-10s | %-12s | %-15s | %-13s|\n", "Date", "Sleep Time", "Wake Up Time", "Sleep Duration", "Sleep Quality")
	fmt.Println("-------------------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		displayARecord(A, n, i)
	}
}

func displayARecord(A record, n int, idx int) {
	fmt.Printf("|%-10s      | %02d:%02d      | %02d:%02d        | %2d h %2d m       | %-13d|\n", 
	A[idx].date, 
	A[idx].sleepTime.hour, A[idx].sleepTime.minute, 
	A[idx].wakeTime.hour, A[idx].wakeTime.minute, 
	A[idx].sleepDuration.hour, A[idx].sleepDuration.minute, 
	A[idx].sleepQuality)
} 
