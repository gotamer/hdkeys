---
title: hdkeys
description: hdkeys is a library to create Bitcoin and Nostr secret and public account keys from the same mnemonic seeds, and a command-line tool that uses this library. It allows for the creation of keys, mnemonic seeds, and Hierarchical Deterministic (HD) addresses.
tags: nostr, nips, NIP05, NIP06, NIP19, bitcoin, BIP32, BIP39, BIP43, BIP44, BIP84, BIP86, BIP173, SLIP44
layout: page
---

hdkeys <img src="https://www.buybitcoinworldwide.com/img/segwit.png" width="100">
========

[![license](https://img.shields.io/badge/license-WTFPL%20--%20Do%20What%20the%20Fuck%20You%20Want%20to%20Public%20License-green.svg)](https://github.com/modood/hdkeygen/blob/master/LICENSE)

hdkeys is a library to create Bitcoin and Nostr secret and public account keys from the same mnemonic seeds, and a command-line tool that uses this library. It allows for the creation of keys, mnemonic seeds, and Hierarchical Deterministic (HD) addresses.

- hdkeys allows for the creation of mnemonic seeds, and Hierarchical Deterministic (HD) addresses.
- hdkeys supports BIP39 passphrase protection for HD Wallets
- hdkeys creates Bitcoin and Nostr keys/accounts from the same mnemonic seeds
- hdkeys can create the WIF (Wallet Import Format), and decode private and public keys from WIF
___________

*   BIP32 - Hierarchical Deterministic Wallets
*   BIP39 - Mnemonic code for generating deterministic keys
*   BIP43 - Purpose Field for Deterministic Wallets
*   BIP44 - Multi-Account Hierarchy for Deterministic Wallets
*   BIP49 - Derivation scheme for P2WPKH-nested-in-P2SH based accounts
*   BIP84 - Derivation scheme for P2WPKH based accounts
*   BIP86 - Derivation scheme for Pay-to-Taproot (P2TR) based accounts
*   BIP173 - Base32 address format for native v0-16 witness outputs
*   SLIP44 - Registered coin types for BIP-0044
*   [NIP05](https://nostr-nips.com/nip-05) - Mapping Nostr keys to DNS-based internet identifiers, HEX key output
*   [NIP06](https://nostr-nips.com/nip-06) - Basic key derivation from mnemonic seed phrase
*   [NIP19](https://nostr-nips.com/nip-19) - bech32-encoded entities (nsec, npub)
___________

Can I trust this code?
----------------------

> Don't Trust. Verify.

Install
-------

See [release](https://github.com/gotamer/hdkeys/releases) page for executables.


Manual / help text
------------------

    hdkeys --help

[help output](doc/help.txt)

____________
Examples
--------

### Wall Of Fame

    hdkeys wof -pass "" -mnemonic "leader monkey parrot ring guide accident before fence cannon height naive bean" > doc/wof.txt

[wof output text](doc/wof.txt)

Don't Trust Verify NIP06 [test vector](https://nostr-nips.com/nip-06#test-vectors)

### Wallet Import Format
    $ hdkeys wif L1VZ55UPgF83k4ndU8BBf62eM9prgo4coie5ttZrvS8GBzddzrhD > doc/wif.txt

[wif output text](doc/wif.txt)


### json key set
    $ hdkeys keyset -pass "" -mnemonic "leader monkey parrot ring guide accident before fence cannon height naive bean" > doc/keyset.json

[keyset output json](doc/keyset.json)

___________

Coders
------
Coders quick start: [link](doc/coders.md)

License
-------

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.

Base code stolen from
---------------------

[hdkeygen](https://github.com/modood/hdkeygen) with same license __Thank you [modood](https://github.com/modood)__
