package types

//Money представляет собой денежнуж сумму в минимальных единицах.
type Money int64

//Category представляет собой категорию, в которой был соверщён платёж.
type Category string


//payment представляет информацию о платеже.
type Payment struct {
	Id       int
	Amount   Money
	Category Category
}
