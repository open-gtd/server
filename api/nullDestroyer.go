package api

type NullDestroyer struct {
}

func (NullDestroyer) Destroy() error {
	return nil
}
