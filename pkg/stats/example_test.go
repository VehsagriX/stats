package stats

import (
	"fmt"

	"github.com/VehsagriX/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{ID: 1, Amount: 100, Status: types.StatusOK},
		{ID: 2, Amount: 200, Status: types.StatusOK},
		{ID: 3, Amount: 300, Status: types.StatusFail},
	}
	result := Avg(payments)
	fmt.Println(result)
	// Output: 100
}

func ExampleAvg_empty() {
	payments := []types.Payment{}
	result := Avg(payments)
	fmt.Println(result)
	// Output: 0
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{ID: 1, Amount: 100, Category: "food", Status: types.StatusOK},
		{ID: 2, Amount: 200, Category: "food", Status: types.StatusFail},
		{ID: 3, Amount: 150, Category: "food", Status: types.StatusOK},
		{ID: 4, Amount: 300, Category: "transport", Status: types.StatusOK},
	}
	result := TotalInCategory(payments, "food")
	fmt.Println(result)
	// Output: 250
}

func ExampleTotalInCategory_empty() {
	payments := []types.Payment{
		{ID: 1, Amount: 100, Category: "food", Status: types.StatusFail},
	}
	result := TotalInCategory(payments, "food")
	fmt.Println(result)
	// Output: 0
}
