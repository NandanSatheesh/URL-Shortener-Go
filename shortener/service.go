package shortener

type RedirectService interface {
	Find(Code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
