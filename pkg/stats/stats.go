package stats

import "github.com/VehsagriX/bank/pkg/bank/types"

// Avg расчитывает среднюю сумму платежей
func Avg(payments []types.Payment) types.Money {
	if len(payments) == 0 {
		return 0
	}
	var sum types.Money
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount

		}
	}
	averangeAmount := sum / types.Money(len(payments))
	return averangeAmount
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}
	return sum
}
