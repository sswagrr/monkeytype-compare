package main

import (
	"fmt"
)

func Total(r Records) float64 {
	return r.Time15 + r.Time30 + r.Time60 + r.Time120 + r.Words10 + r.Words25 + r.Words50 + r.Words100
}

func PrintDiff(r1, r2 Records) {
	RecordDiff := Records{
		Name:     "diff",
		Time15:   r1.Time15 - r2.Time15,
		Time30:   r1.Time30 - r2.Time30,
		Time60:   r1.Time60 - r2.Time60,
		Time120:  r1.Time120 - r2.Time120,
		Words10:  r1.Words10 - r2.Words10,
		Words25:  r1.Words25 - r2.Words25,
		Words50:  r1.Words50 - r2.Words50,
		Words100: r1.Words100 - r2.Words100,
	}

	PrintRecords(RecordDiff)

	diff := Total(RecordDiff)
	if diff > 0 {
		fmt.Printf("'%s' wins by %.2f wpm.\n", r1.Name, diff)
	} else if diff < 0 {
		fmt.Printf("'%s' wins by %.2f wpm.\n", r2.Name, -diff)
	} else {
		fmt.Println("Perfect equality!")
	}

}

func PrintHeader() {
	fmt.Printf("%-15s %6s %6s %6s %6s %6s %6s %6s %6s\n",
		"", "15s", "30s", "60s", "120s", "10", "25", "50", "100")
}

func DisplayRecords(r1, r2 Records) {
	PrintHeader()
	PrintRecords(r1)
	PrintRecords(r2)
	PrintDiff(r1, r2)
}

func PrintRecords(r Records) {
	line := fmt.Sprintf("%6.2f %6.2f %6.2f %6.2f %6.2f %6.2f %6.2f %6.2f\n",
		r.Time15, r.Time30, r.Time60, r.Time120, r.Words10, r.Words25, r.Words50, r.Words100)

	fmt.Printf("%-15s %s", r.Name, line)
}
