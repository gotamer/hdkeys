hdkeys <img src="https://www.buybitcoinworldwide.com/img/segwit.png" width="100">
========

[![license](https://img.shields.io/badge/license-WTFPL%20--%20Do%20What%20the%20Fuck%20You%20Want%20to%20Public%20License-green.svg)](https://github.com/modood/hdkeygen/blob/master/LICENSE)

A very simple and easy to use bitcoin(btc) and nostr key/wallet generator.

- hdkeys allows for the creation of mnemonic seeds, and Hierarchical Deterministic (HD) addresses.
- hdkeys supports BIP39 passphrase protection.
- creates Bitcoin and Nostr accounts from the same mnemonic seeds
- hdkeys can create WIF (Wallet Import Format), and decode private keys from WIF

*   BIP32 - Hierarchical Deterministic Wallets
*   BIP39 - Mnemonic code for generating deterministic keys
*   BIP43 - Purpose Field for Deterministic Wallets
*   BIP44 - Multi-Account Hierarchy for Deterministic Wallets
*   BIP49 - Derivation scheme for P2WPKH-nested-in-P2SH based accounts
*   BIP84 - Derivation scheme for P2WPKH based accounts
*   BIP86 - Derivation scheme for Pay-to-Taproot (P2TR) based accounts
*   BIP173 - Base32 address format for native v0-16 witness outputs
*   SLIP44 - Registered coin types for BIP-0044
*   NIP06 - Basic key derivation from mnemonic seed phrase (https://nostr-nips.com/nip-06)

Can I trust this code?
----------------------

> Don't Trust. Verify.
> We recommend every user of this library audit and verify any underlying code for its validity and suitability.

Install
-------
See release page for executables. Link at the right ->

Coders quick start
------------------

```bash
git clone https://github.com/gotamer/hdkeys
cd hdkeys
./make.sh fmt
./make.sh build
./make.sh release
```

Cli Help
--------

```bash
hdkeys - Mnemonic seeds and Hierarchical Deterministic (HD) addresses.

SYNOPSIS
	hdkeys [OPTIONS]
	hdkeys [OPTIONS] [COMMAND] [COMMAND OPTIONS]

ENVIRONMENT VARIABLE OPTIONS
	HDKEYS_PASSWORD
	HDKEYS_MNEMONIC
	HDKEYS_WIF

DESCRIPTION
hdkeys allows for the creation of mnemonic seeds, and Hierarchical Deterministic (HD) addresses.

- hdkeys supports BIP39 passphrase protection.
- hdkeys creates Bitcoin and Nostr accounts from the same mnemonic seeds
- hdkeys can create WIF (Wallet Import Format), and decode private keys from WIF

    BIP32 - Hierarchical Deterministic Wallets
    BIP39 - Mnemonic code for generating deterministic keys
    BIP43 - Purpose Field for Deterministic Wallets
    BIP44 - Multi-Account Hierarchy for Deterministic Wallets
    BIP49 - Derivation scheme for P2WPKH-nested-in-P2SH based accounts
    BIP84 - Derivation scheme for P2WPKH based accounts
    BIP86 - Derivation scheme for Pay-to-Taproot (P2TR) based accounts
    BIP173 - Base32 address format for native v0-16 witness outputs
    SLIP44 - Registered coin types for BIP-0044
    NIP06 - Basic key derivation from mnemonic seed phrase
	...

COMMANDS
	wof
		Wall Of Fame, prints a whole set of keys
		COMMAND OPTIONS
			-count [int] (default = 1)
				Set number of keys to generate.

	wif [string] or [Environment variable]
		Decode the private key from wif(Wallet Import Format), then generate the address.

	keyset
		Gets a Bitcoin and Nostr key set with the same WIF (Wallet Import Format) as JSON.
		COMMAND OPTIONS
			-no [int] (default = 0)
				Nostr Account number to generate

OPTIONS
	-mnemonic [string]
		Mnemonic words

	-pass [string]
		Protect bip39 mnemonic with a passphrase via flag,
		or use environment variable,
		or you will be asked to enter at a prompt.

	-h or --help [bool]
		Report usage information and exit.

	-v [bool]
		Print version tag and exit.

	--version [bool]
		Print detailed version and exit.

	-d or --debug [bool]
		Print debug information about the process.

	--verbose [bool]
		Print extra verbose debug information about the process.

REPORTING BUGS
	npub12jjczvd2mzstyhr468fyas7vzmsm5d2x3tv5l9tev6q0jakk9djqx4uk7x

REPO
	https://github.com/gotamer/hdkeys

COPYRIGHT
	© 2025 WTFPL – Do What the Fuck You Want to Public License. (http://www.wtfpl.net)
```


Example
-------

**generate bip39 mnemonic with passphrase:**

```txt
$ hdkeys -n 3 -pass very_secret_bip39_passphrase

BIP39 Mnemonic:    unlock hole vessel genius tube topic twenty slim frequent crash obey keep
BIP39 Passphrase:  very_secret_bip39_passphrase
BIP39 Seed:        a977a9a88acdf7c166db91a3b85d7d66636ce59e25d54f844b8b455cbb7f89e4aced81048c6c798b4f4d825f50c31460bf774d5882c7a769795c348070b721c7
BIP32 Root Key:    xprv9s21ZrQH143K45DNnkfXyJnm2c9t97ed6mLR2FyYeU5G4LGoBhwHtatqnzzmfvXS3kbpKHFTqtZKAMF7JPFYR7bXhV4mrT3NoHjyjmFRrSj

Path(BIP44)        Legacy(P2PKH, compresed)           WIF(Wallet Import Format)
----------------------------------------------------------------------------------------------------------
m/44'/0'/0'/0/0    182Y79r2nvyysvDQF6H1PiApCfCCTasngh L4odrT4EvvfEpLoxtcas8xrQEqhXmxqfovRSf8oo3gJWSLDBzNzs
m/44'/0'/0'/0/1    1M3zgBtCknJ98w6iXKTu25hgNMkf73dBss L4N2mHwsowJzPiqMooeNpmTJ3Um2PiopFrWRgw43i65YguYJV4o4
m/44'/0'/0'/0/2    1EvicJdixb9aRSjbrL7V1SgMQ88LXztxi7 KzvE4Ea36VunGEFAsN7p6KsQWxmBNv5btRYXZFfKHAGPURDUfUL9

Path(BIP49)        SegWit(P2WPKH-nested-in-P2SH)      WIF(Wallet Import Format)
----------------------------------------------------------------------------------------------------------
m/49'/0'/0'/0/0    34A2LmikUmch525awgPNiGZoTtkpwK1jFH L2CgG2uoUpAh3RMxHDgMKeiZzyRpsT2KkQtikpTBZexP4RqjZoQh
m/49'/0'/0'/0/1    3EKe4TsUM7wvtnphRvAFrcZH4eLpyyFwzJ KwdUZUbNRmV42vqEDZdXEUNA4anurxi6ZyoWSew7MDNyJhVuyH2V
m/49'/0'/0'/0/2    3FvGV8YrWTKWLqvoKqfFbPvkdQE73HYyQ2 L4YFui7KgGPF6ThwzY9XyKRhTRvirZTjq3zzotZRcGJvCLGrtXnU

Path(BIP84)        SegWit(P2WPKH, bech32)                     WIF(Wallet Import Format)
------------------------------------------------------------------------------------------------------------------
m/84'/0'/0'/0/0    bc1qtm476yzdvq4utp8qvcfl8dt8t7n5fajnlk5chr L2tmKALybcbskuGzoJDGnnn4QNwjWChR1MwKcRFH8A44xb4Q2qx1
m/84'/0'/0'/0/1    bc1qvwkut6j8guejw5970amrtq4e8l8k3728h80gyz Kxo8vovuKzQ5RwydkhsvzJjP1rL5hToRccJSboWTiuYxgxLA7M7P
m/84'/0'/0'/0/2    bc1qkut6v75kucmp6n5my4vjg2qahn4rs95qjx2enk KxvkGXkZBKHWDLrQ3vwPe94BGAweZhWvSpJmtD3imWpASQnYiYdQ

Path(BIP86)        Taproot(P2TR, bech32m)                                         WIF(Wallet Import Format)
--------------------------------------------------------------------------------------------------------------------------------------
m/86'/0'/0'/0/0    bc1ptckpf58pvsa07fqaj3gfawjtceahh6gk7d7c374adhlsc8sy7gas5yla7m L2PpsRHqQQdD7BW1L7DZVyiJDsjSdXDNAB4fQBHzyNorYaqQkpfw
m/86'/0'/0'/0/1    bc1p8d0adwzaycl4f2vcuf2uxgvf3qmvqk7lt2wkwvgk0t7af92fpsxsfhmj46 L5RsQhb929ZMoM2Jk7ZhEybAu5BkdbYpds7oeCtd2kRtXH2RufXg
m/86'/0'/0'/0/2    bc1p0pgatlg43wmepyr3e4kxhvhmfqruqcczatfhwhzydmgay6g8t5wsz3hvjw KwbLrW2qJUXRgn3e5PtpEMrdyegBMH62N1n2BD3Tp9m2DveKrHEP

Path(BIP44)        Nostr
-----------------------------------------------------------------------------------------
m/44'/1237'/0'/0/0 wif:   L4YJqAH2nCjNG2xMgaw7uHLj1Dy11zbXH7B7Sh3nbSbmEah7cPtB
m/44'/1237'/0'/0/0 Nostr: nsec1mfkaq5sx58ht30js9n2lux0gwn9xeqk3rnca00nf42ujs6ljv0hqh3pr8h

m/44'/1237'/1'/0/0 wif:   Kx1ZjxtKq9ThktKTXnJ3fDemMg1GMATQVRTULKgD7rwm5aTUU6YK
m/44'/1237'/1'/0/0 Nostr: nsec1z7tdj7krz7z47q8w6ervy0srfty0mq00dyw6sl4z8643hh7kj7cqfa27gm

m/44'/1237'/2'/0/0 wif:   KyYVFvYAXdvtRN76aQMRRTXcVQKCjah8VgpaGwBYygtgK6EEbYbE
m/44'/1237'/2'/0/0 Nostr: nsec1g42e7t9ewww4cgwrvtml4gjj9ggkz6p4q64yj9t05ydmwfas7d7q5fh0yy

```

**re-generate a root key with existing mnemonic:**

```txt
$ hdkeys -mnemonic "ordinary volcano company hedgehog usage success awkward filter state energy wool point" -pass 123456
```

**generate with existing wif:**

```txt
$ hdkeys -wif KyHXurGfBovpZpgQjG37ujZaNQobN8rUcamafamJWXwXHkumzVEV

 Wallet Import Format:
 *   WIF(compressed):         KyHXurGfBovpZpgQjG37ujZaNQobN8rUcamafamJWXwXHkumzVEV
 *   WIF(uncompressed):       5JHS8SfSRCYE3TD6q6DiSYPnd7kdna3boaUmnSmvp27v5p7V7E8

 Public Addresses:
 *   Legacy(compresed):       19ZZTqYbz4jgGJjTmsMYujTYfsWTEMZPKE
 *   Legacy(uncompressed):    1FP329UGwtcguc3k24ygfsu5v9TokvSmvf
 *   SegWit(nested):          3G1MmjVWALjdCD6WxR5g4BXQ85JoV76KQF
 *   SegWit(bech32):          bc1qth5k5cvq8c5zpp0mczrh0cxulvafm2vlvgkkvz
 *   Taproot(bech32m):        bc1p6xlcwdx2e27plvmhaxh0nnrztw69k5ccw4xvckc43kezq5l2xy0qju6uje

```

```bash
$ hdkeys keyset -mnemonic "leader monkey parrot ring guide accident before fence cannon height naive bean"
{"Nostr":{"Path":"m/44'/1237'/0'/0/0","Wif":"L1VZ55UPgF83k4ndU8BBf62eM9prgo4coie5ttZrvS8GBzddzrhD","NSec":"nsec10allq0gjx7fddtzef0ax00mdps9t2kmtrldkyjfs8l5xruwvh2dq0lhhkp","NPub":"npub1zutzeysacnf9rru6zqwmxd54mud0k44tst6l70ja5mhv8jjumytsd2x7nu"},"Bitcoin":{"Wif":"L1VZ55UPgF83k4ndU8BBf62eM9prgo4coie5ttZrvS8GBzddzrhD","Lagacy":"1FyKyNJKMVDLkhH29KmDPUNubcnSrQcoWR","Nested":"3Me6rM5XAfKNGDhVpcozk7NcRj7CDnGYow","SegWit":"bc1q5suwgj96zvqgjq2d8df6wrkxuffz93fxsx3n3l","Taproot":"bc1pmupac93le8w8ncqmsupfmcvc9667t836p9q78zc26pqff5zf5xxswh9re2"}}
```
```json
{
  "Nostr": {
    "Path": "m/44'/1237'/0'/0/0",
    "Wif": "L1VZ55UPgF83k4ndU8BBf62eM9prgo4coie5ttZrvS8GBzddzrhD",
    "NSec": "nsec10allq0gjx7fddtzef0ax00mdps9t2kmtrldkyjfs8l5xruwvh2dq0lhhkp",
    "NPub": "npub1zutzeysacnf9rru6zqwmxd54mud0k44tst6l70ja5mhv8jjumytsd2x7nu"
  },
  "Bitcoin": {
    "Wif": "L1VZ55UPgF83k4ndU8BBf62eM9prgo4coie5ttZrvS8GBzddzrhD",
    "Lagacy": "1FyKyNJKMVDLkhH29KmDPUNubcnSrQcoWR",
    "Nested": "3Me6rM5XAfKNGDhVpcozk7NcRj7CDnGYow",
    "SegWit": "bc1q5suwgj96zvqgjq2d8df6wrkxuffz93fxsx3n3l",
    "Taproot": "bc1pmupac93le8w8ncqmsupfmcvc9667t836p9q78zc26pqff5zf5xxswh9re2"
  }
}
```

Example address path #7 with Passphrase
```bash
$ hdkeys keyset -no 7 -pass 2345 -mnemonic "leader monkey parrot ring guide accident before fence cannon height naive bean"
```
```json
{
  "Nostr": {
    "Path": "m/44'/1237'/7'/0/0",
    "Wif": "L35EeWsQHijBKXJQJMdCnGrE8L2xLsYTicomb8kAgujKrcigZFYw",
    "NSec": "nsec1464z2lqk9q5a7ct79wunnax3w2urrkpmsz4u9c6zttws4ueg89ssfz8xa6",
    "NPub": "npub15jqt9xfzf7awawtls6qq9cy5hj4rpgxrs6meqnae7eaw2nvgsqksdr0lcp"
  },
  "Bitcoin": {
    "Wif": "L35EeWsQHijBKXJQJMdCnGrE8L2xLsYTicomb8kAgujKrcigZFYw",
    "Lagacy": "1LGj6Vyi24uZyLWURK2Gm3eu26vrn1L43D",
    "Nested": "3GygZBMK6vfQ4R2Vcw9K6dG7GWwfsV4dkq",
    "SegWit": "bc1q6d3uftavdr6hresf5xl3e7clr6kt2l5ednwlf8",
    "Taproot": "bc1pulhwyfv8e6akfksqsdugn2wyk8fn4n0hft3rdx0mtke3npmwzrxqgjc2sf"
  }
}
```

License
-------

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.

Base code stolen from
---------------------

https://github.com/modood/hdkeygen
