package wordsbank

type Repo struct {
	words map[string]struct{}
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) GetWords() map[string]struct{} {
	return r.words
}

func (r *Repo) SetWords(words map[string]struct{}) {
	r.words = words
}
