package mundipagg

// Mundipagg struct that holds all functions working in the lib
type Mundipagg interface {
	NewSubscription(s Subscription, indepotencyKey string) (string, error)
	NewCustomer(c Customer, indepotencyKey string) (string, error)
}

// Login A struct that holds the login information
type mundipagg struct {
	BasicSecretAuthKey string
}

// New Create a Mundipagg object
func New(basicSecretAuthKey string, testBasicSecretKey string, isProduction bool) (Mundipagg, error) {
	secret := basicSecretAuthKey
	if !isProduction {
		secret = testBasicSecretKey
	}

	// Create the mundipagg object
	mund := mundipagg{
		BasicSecretAuthKey: secret,
	}

	return mund, nil
}

func main() {}