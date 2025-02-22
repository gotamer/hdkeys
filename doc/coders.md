Coders
======

- Use as library
- or as bin. Pipe in mnemonic seed and password and get a json set of keys for nostr and bitcoin

Quick start
-----------

```bash
git clone https://github.com/gotamer/hdkeys
cd hdkeys
./make.sh fmt
./make.sh build
./make.sh release
```

Example address path #7, from Mnemonic with Passphrase
```bash
$ hdkeys keyset -no 7 -pass 2345 -mnemonic "leader monkey parrot ring guide accident before fence cannon height naive bean"
```

##### Note: Both Nostr and Bitcoin with the same wif (Wallet Import Format)

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
