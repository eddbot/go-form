package main

type SignupForm struct {
	Username string
	Errors   []string
}

func (sf *SignupForm) Validate() {
	if sf.Username != "terry" {
		sf.Errors = append(sf.Errors, "username must be terry")
	}

	if sf.Username == "" {
		sf.Errors = append(sf.Errors, "username cannot be blank")
	}

	if len(sf.Username) < 3 {
		sf.Errors = append(sf.Errors, "username must be at least 3 characters")
	}

	if len(sf.Username) > 10 {
		sf.Errors = append(sf.Errors, "username must be at most 10 characters")
	}

}
