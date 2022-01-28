package services

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct {
}

type itemsServiceInterface interface {
	GetItems(id int) []interface{}
	DeleteItems(id int)
}

func (i *itemsService) GetItems(id int) []interface{} {
	return nil
}

func (i *itemsService) DeleteItems(id int) {

}
