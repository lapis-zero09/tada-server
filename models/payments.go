package models

type Payments struct {
	Id      int
	PlaceId int
	Cost    int
}

func SamplePayments() []Payments {
	payments := make([]Payments, 0, 10)
	for i := 0; i < 10; i++ {
		payments = append(payments, Payments{Id: i, PlaceId: i, Cost: i * 1000})
	}
	return payments
}
