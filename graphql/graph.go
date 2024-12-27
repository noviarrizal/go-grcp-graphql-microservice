package main

// type Server Struct {
// 	accountClient 	*account.Client
// 	catalogClient 	*catalog.Client
// 	orderClient 	*order.Client
// }

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	catalogclient, err := catalog.NewClient(catalogUrl)
	if err != nil {
		accountClient.Close()
		return nil, err
	}

	orderClient, err := order.NewClient(orderUrl)
	if err != nil {
		accountClient.Close()
		catalogclient.Close()
		return nil, err
	}

	return &Server{
		accountClient,
		catalogClient,
		orderClient,
	}, nil
}
