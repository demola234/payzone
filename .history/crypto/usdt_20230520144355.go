package crypto

import (
	"github.com/foxnut/go-hdwallet"
)

// Create a usdt wallet for a user
func CreateUsdtWallet() {
	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
	)
	if err != nil {
		panic(err)
	}

	//  1. Generate a new address
	

	//  2. Save the address to the database
	//  3. Return the address to the user

}
