package samples

type Sample struct {
	Text   string `bson:"text"`
	Number string `bson:"number"`
	Found  bool   `bson:"found"`
	Type   string `bson:"type"`
}
