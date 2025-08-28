package github

type Gist struct {
	Url    string
	Id     string
	Files  map[string]File
	Public bool
}
