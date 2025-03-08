// bitcoin libs
package hdkeys

import (
	"github.com/btcsuite/btcd/btcutil"
)

const CoinTypeBTC CoinType = 0x80000000

type Bitcoin struct {
	Wif     string // Wallet Import Format
	Lagacy  string // Pay-to-Pubkey Hash (P2PKH)
	Nested  string // Pay to Script Hash (P2SH) includes "Pay-to-Witness-Pubkey Hash"
	SegWit  string // Bech32
	Taproot string // Bech32 with Schnorr signing
}

func GetBitcoinFromWif(wif string) (*Bitcoin, error) {

	var btc = new(Bitcoin)

	wiftype, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return btc, err
	}

	wifCompressed, addressCompressed, segwitBech32, segwitNested, taproot, err := CalculateFromPrivateKey(wiftype.PrivKey, true)
	if err != nil {
		return btc, err
	}

	btc.Wif = wifCompressed
	btc.Lagacy = addressCompressed
	btc.Nested = segwitNested
	btc.SegWit = segwitBech32
	btc.Taproot = taproot
	return btc, nil
}
