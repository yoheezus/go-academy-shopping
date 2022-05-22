package product

import "errors"

var (
	productCounter int
	productNames   map[string]bool
	allProducts    map[int]*Product
	faulty_prod    Product
)

func init() {
	productNames = make(map[string]bool)
	allProducts = make(map[int]*Product)

	faulty_prod = New("Faulty Product")
	allProducts[faulty_prod.ID] = &faulty_prod
}

func New(name string) Product {
	productCounter++
	p := Product{
		ID: productCounter,
	}
	err := p.SetName(name)
	if err != nil {
		return faulty_prod
	}
	save(p)
	return p
}

func (p *Product) GetName() string {
	return p.ProductName
}

func (p *Product) SetName(productName string) error {
	if productName == "" {
		return errors.New("cannot set name to nothing")
	}
	if productName == p.ProductName {
		return nil
	}

	if _, prs := productNames[productName]; prs {
		return errors.New("name of product already exists")
	}
	productNames[productName] = true
	p.ProductName = productName
	update(*p)
	return nil
}

func (p *Product) GetPrice() int {
	if p.SalePercentage != 0 {
		return (100 - p.SalePercentage) * (p.ProductPrice / 100)

	}
	return p.ProductPrice
}

func (p *Product) SetPrice(productPrice int) error {
	if productPrice < 0 {
		return errors.New("price cannot be below 0")
	}
	p.ProductPrice = productPrice
	update(*p)
	return nil
}

func (p *Product) GetStock() int {
	return p.StockAmount
}

func (p *Product) SetStock(stockAmt int) error {
	if stockAmt < 0 {
		return errors.New("stock cannot be less than 0")
	}
	p.StockAmount = stockAmt
	update(*p)
	return nil
}

func (p *Product) GetID() int {
	return p.ID
}

func GetByID(ID int) (Product, error) {
	pp := allProducts[ID]

	return *pp, nil
}

func update(updated Product) {
	allProducts[updated.ID] = &updated
}

func save(p Product) error {
	_, prs := allProducts[p.ID]
	if prs {
		return errors.New("product already exists")
	}

	allProducts[p.ID] = &p
	return nil
}

func (p *Product) SetSalePercentage(percentage int) error {
	if percentage > 90 {
		return errors.New("cannot set sale percentage over 90")
	}
	p.SalePercentage = percentage
	update(*p)
	return nil
}
