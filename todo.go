// 1. Figure out how to map db res to a go struct and then go struct to my frontend in the same 
// format that will render correctly out the box
// 2. Figure out how to use *mongo.Next instead of Al, which will not scale
func getCardsData(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URL))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mongo db connected succesfully")
	}

	collection := client.Database("haikyuuDB").Collection("haikyu")

	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var allData AllCardData
	for cur.Next(context.Background()) {

		var data CardData
		err := cur.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		allData.CardsData = append(allData.CardsData, data)

	}
	json.NewEncoder(w).Encode(allData)


}

type CardData struct {
	Anime    string   `bson:"anime"`
	Fname    string   `bson:"fname"`
	Lname    string   `bson:"lname"`
	Position string   `bson:"position"`
	Quote    string   `bson:"quote"`
	Images   []string `bson:"images"`
}

type AllCardData struct {
	CardsData []CardData `bson:"results"`
}