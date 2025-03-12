
Hierarchical Deterministic Keys
===============================

hdkeys is a library to create Bitcoin and Nostr secret and public account keys from the same mnemonic seeds, and a command-line tool that uses this library.  

- It allows for the creation of __mnemonic seeds__ (or bring your own).  
- From the __mnemonic seeds__ it will create __Hierarchical Deterministic (HD)__ addresses, AKA __HD Wallets__.
- __HD Wallets__ means you may create nearly unlimited __subkeys__ for Bitcoin and Nostr
- hdkeys supports __BIP39 passphrase protection__ for __HD Wallets__
- hdkeys can create the __Wallet Import Format (WIF)__, and decode private and public keys from __WIF__


Why
------- 

It is hard to keep secrets specially multiple secrets, with __hdkeys__ you have to keep only your __mnemonic seeds__ secret. All other Bitcoun and Nostr keys are derived from your __mnemonic seeds__ and can be reproduced at any time.

This alows you to have multiple identieties. And you will be able to prove selectivly that those identieties belong to you.  
More identieties equals more privacy, means more freedom!


Standards
---------

- BIPxx = Bitcoin standards
- NIPxx = Nostr standards


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


Can I trust this code?
----------------------

> Don't Trust. Verify.


Install and documentation
-------------------------

[Documents](https://github.com/gotamer/hdkeys/tree/master/doc/README.md)
