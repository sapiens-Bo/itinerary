package abonent

type Abonent struct {
	ID        string
	Name      string
	ShortName string
	Contract  Contract
	Site      Site
	Rate      float32
	Phone     string
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

type Equipment struct {
	Type      string
	Name      string
	SerialNum int64
	Status    string
}
