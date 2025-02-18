package hdkeys

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil/bech32"
)

const CoinTypeNostr CoinType = 0x800004d5

func NostrGetPrivateKey(privateKeyBytes []byte) (address string, err error) {

	bits5, err := bech32.ConvertBits(privateKeyBytes, 8, 5, true)
	if err != nil {
		return "", err
	}
	address, err = bech32.Encode("nsec", bits5)
	return address, err
}

func NostrGetPublicKey(sk []byte) (npub string, err error) {
	var pub string
	if pub, err = nostrGetPublicKey(sk); err == nil {
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

/*
func EncodePrivateKey(privateKeyHex string) (string, error) {
	b, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key hex: %w", err)
	}

	bits5, err := bech32.ConvertBits(b, 8, 5, true)
	if err != nil {
		return "", err
	}

	return bech32.Encode("nsec", bits5)
}

func EncodePublicKey(publicKeyHex string) (string, error) {
	b, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key hex: %w", err)
	}

	bits5, err := bech32.ConvertBits(b, 8, 5, true)
	if err != nil {
		return "", err
	}

	return bech32.Encode("npub", bits5)
}

*/
