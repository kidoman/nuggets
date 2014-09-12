package types

var Nuggets = make([]string, 0)

func Register(url string) {
	Nuggets = append(Nuggets, url)
}
