package testing

type Retriever struct {
	Content string
}

func (Retriever) Get(url string) string {
	return "testing data"
}