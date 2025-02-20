// nostr libs
package hdkeys

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/btcsuite/btcd/chaincfg"
)

const CoinTypeNostr CoinType = 0x800004d5

type Nostr struct {
	Path string
	Wif  string
	NSec string
	NPub string
}

func (km *KeyManager) GetNostr(no uint32) (nostr *Nostr, err error) {

	pathKey, err := km.GetKey(PurposeBIP44, CoinTypeNostr, no, 0, 0)
	if err != nil {
		return
	}

	// private key
	privateKey, _ := btcec.PrivKeyFromBytes(pathKey.Bip32Key.Key)

	// generate the wif(wallet import format) string
	nostrWif, err := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, true)
	if err != nil {
		return
	}

	nsec, err := NostrGetPrivateKey(pathKey.Bip32Key.Key)
	if err != nil {
		return
	}

	npub, err := NostrGetPublicKey(pathKey.Bip32Key.Key)
	if err != nil {
		return
	}

	nostr = new(Nostr)
	nostr.Path = pathKey.Path
	nostr.Wif = nostrWif.String()
	nostr.NSec = nsec
	nostr.NPub = npub
	return
}

func NostrGetPrivateKey(privateKeyBytes []byte) (nsec string, err error) {

	bits5, err := bech32.ConvertBits(privateKeyBytes, 8, 5, true)
	if err != nil {
		return "", err
	}
	nsec, err = bech32.Encode("nsec", bits5)
	return nsec, err
}

func NostrGetPublicKey(privateKeyBytes []byte) (npub string, err error) {
	var pub string
	if pub, err = nostrGetPublicKey(privateKeyBytes); err == nil {
		npub, err = nostrEncodePublicKey(pub) // Encode to HEX
	}
	return
}

func nostrGetPublicKey(privateKey []byte) (string, error) {
	_, pubkey := btcec.PrivKeyFromBytes(privateKey)
	return hex.EncodeToString(schnorr.SerializePubKey(pubkey)), nil
}

// nip19
func nostrEncodePublicKey(publicKeyHex string) (string, error) {
	b, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", err
	}

	bits5, err := bech32.ConvertBits(b, 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode("npub", bits5)
}
