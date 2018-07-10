package gohq

// Create a new session using a direct login token

// Hello HQ if you're looking at this, you got beanboozled by a 14 year old <3
func New(loginToken string) (account *Account, err error) {
	account = &Account{LoginToken: loginToken}
	tokens, err := account.Tokens()
	if err != nil {
		return
	}

	account.AccessToken = tokens.AccessToken
	account.AuthToken = tokens.AuthToken
	account.LoginToken = tokens.LoginToken
	account.Admin = tokens.Admin
	account.Guest = tokens.Guest
	account.Tester = tokens.Tester

	me, err := account.Me()
	if err != nil {
		return
	}

	account.Username = me.Username
	account.UserID = me.UserID
	account.AvatarURL = me.AvatarURL

	return
}
