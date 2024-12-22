package abonent

type Abonent struct {
	ID        string
	Name      string
	ShortName string
	Contract  Contract
	Site      Site
	Rate      float32
	Phone     string
	Equip     Equipment
}

// Договор
type Contract struct {
	ID     string
	Status string
}

// Площадка
type Site struct {
	ID      int
	Status  string
	Address string
}

// Оборудование
type Equipment struct {
	Type      string
	Name      string
	SerialNum string
	Status    string
}
