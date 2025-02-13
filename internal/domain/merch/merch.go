package merch

import "fmt"

type Merch struct {
	ID    string
	Name  string
	Price int
}

var ErrMerchNotFound = fmt.Errorf("merch not found")
