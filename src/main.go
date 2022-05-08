package main

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	// if err != nil {
	// 	panic(err)
	// }

	// defer client.Disconnect(ctx)

	// err = client.Ping(ctx, readpref.Primary())

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("GREAT!!!")

	// udCollection := client.Database("basec").Collection("users-datas")

	// res, err := udCollection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})

	// if err != nil {
	// 	fmt.Errorf("Erro de inserção", err)
	// }

	// id := res.InsertedID

	// fmt.Println(id, " Criado com sucesso :D")

	config := NewConfig()

	client, collection, err := ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	bcRepository := NewBaseCRepository(client, collection)

	bcService := NewBaseCService(config, bcRepository)

	server := NewServer(config, bcService)

	server.Run()
}
