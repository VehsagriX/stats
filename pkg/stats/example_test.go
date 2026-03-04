package stats

import (
	"fmt"

	"github.com/VehsagriX/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{Amount: 100},
		{Amount: 200},
		{Amount: 300},
	}
	result := Avg(payments)
	fmt.Println(result)
	// Output: 200
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{Amount: 100, Category: "food"},
		{Amount: 200, Category: "transport"},
		{Amount: 150, Category: "food"},
		{Amount: 300, Category: "transport"},
	}
	result := TotalInCategory(payments, "food")
	fmt.Println(result)
	// Output: 250
}
