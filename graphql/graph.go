package main

import "github.com/99designs/gqlgen/graphql"

type Server struct {
	//     accountClient *account.Client
	//     catalogClient *catalog.Client
	//     orderClient  *order.Client
}

// // this takes urls of clients
func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	// 	accountClient, err := account.NewClient(accountUrl)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	catalogClient, err := catalog.NewClient(catalogUrl)
	// 	if err != nil {
	// 		accountClient.Close() // dependent connection
	// 		return nil, err
	// 	}

	// 	orderClient, err := order.NewClient(orderUrl)
	// 	if err != nil {
	// 		accountClient.Close() // dependent connections should be close same for below
	// 		catalogClient.Close()
	// 		return nil, err
	// 	}

	return &Server{
		// accountClient,
		// catalogClient,
		// orderClient,
	}, nil
}

// func (s *Server) Mutation() MutationResolver {
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() QueryResolver {
// 	return &queryResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() AccountResolver {
// 	return &accountResolver{
// 		server: s,
// 	}
// }

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
