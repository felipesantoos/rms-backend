package credentials

var _ Credentials = &credentials{}

type Credentials interface {
	Email() string
	Password() string

	SetEmail(string)
	SetPassword(string)
}

type credentials struct {
	email    string
	password string
}

func (instance *credentials) Email() string {
	return instance.email
}

func (instance *credentials) Password() string {
	return instance.password
}

func (instance *credentials) SetEmail(email string) {
	instance.email = email
}

func (instance *credentials) SetPassword(password string) {
	instance.password = password
}
