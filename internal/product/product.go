package product

import "errors"

var (
	productCounter int
	productNames   map[string]bool
	allProducts    map[int]*product
	faulty_prod    product
)

func init() {
	productNames = make(map[string]bool)
	allProducts = make(map[int]*product)

	faulty_prod = New("Faulty Product")
	allProducts[faulty_prod.ID] = &faulty_prod
}

func New(name string) product {
	productCounter++
	p := product{
		ID: productCounter,
	}
	err := p.SetName(name)
	if err != nil {
		return faulty_prod
	}
	save(p)
	return p
}

func (p *product) GetName() string {
	return p.ProductName
}

func (p *product) SetName(productName string) error {
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

func (p *product) GetPrice() int {
	if p.SalePercentage != 0 {
		return (100 - p.SalePercentage) * (p.ProductPrice / 100)

	}
	return p.ProductPrice
}

func (p *product) SetPrice(productPrice int) error {
	if productPrice < 0 {
		return errors.New("price cannot be below 0")
	}
	p.ProductPrice = productPrice
	update(*p)
	return nil
}

func (p *product) GetStock() int {
	return p.StockAmount
}

func (p *product) SetStock(stockAmt int) error {
	if stockAmt < 0 {
		return errors.New("stock cannot be less than 0")
	}
	p.StockAmount = stockAmt
	update(*p)
	return nil
}

func (p *product) GetID() int {
	return p.ID
}

func GetByID(ID int) (product, error) {
	pp := allProducts[ID]

	return *pp, nil
}

func update(updated product) {
	allProducts[updated.ID] = &updated
}

func save(p product) error {
	_, prs := allProducts[p.ID]
	if prs {
		return errors.New("product already exists")
	}

	allProducts[p.ID] = &p
	return nil
}

func (p *product) SetSalePercentage(percentage int) error {
	if percentage > 90 {
		return errors.New("cannot set sale percentage over 90")
	}
	p.SalePercentage = percentage
	update(*p)
	return nil
}
