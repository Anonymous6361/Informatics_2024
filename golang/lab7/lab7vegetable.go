package lab7

type Vegetable struct {
	Name  string
	Price float64
	Sort string
}

func (v *Vegetable) GetName() string {
	return v.Name
}

func (v *Vegetable) GetPrice() float64 {
	return v.Price
}

func (v *Vegetable) SetPrice(price float64) {
	v.Price = price
}

func (v *Vegetable) ApplyDiscount(discount float64) {
	v.Price -= v.Price * discount / 100
}
func (v *Vegetable) GetSort(newSort string) {
	v.Sort = newSort
}