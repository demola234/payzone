package crypto

import (
	"fmt"

	"github.com/foxnut/go-hdwallet"
)

var (
	mnemonic = "range sheriff try enroll deer over ten level bring display stamp recycle"
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
	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.USDC), hdwallet.AddressIndex(1))
	address, _ := wallet.GetAddress()
	addressP2WPKH, _ := wallet.GetKey().AddressP2WPKH()
	addressP2WPKHInP2SH, _ := wallet.GetKey().AddressP2WPKHInP2SH()
	fmt.Println("BTC: ", address, addressP2WPKH, addressP2WPKHInP2SH)

	//  2. Save the address to the database
	//  3. Return the address to the user

}
