package redirect

type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) (*Redirect, error)
}
