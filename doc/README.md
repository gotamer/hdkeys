Hierarchical Deterministic Keys
===============================

Install
-------

See [release](https://github.com/gotamer/hdkeys/releases) page for executables.


Manual / help text
------------------

    hdkeys --help

[help output](https://github.com/gotamer/hdkeys/tree/master/doc/help.txt)


Examples
--------

### Wall Of Fame

    hdkeys wof -pass "" -mnemonic "leader monkey parrot ring guide accident before fence cannon height naive bean" > doc/wof.txt

[wof output text](https://github.com/gotamer/hdkeys/tree/master/doc/wof.txt)

Don't Trust Verify NIP06 [test vector](https://nostr-nips.com/nip-06#test-vectors)

### Wallet Import Format
    $ hdkeys wif L1VZ55UPgF83k4ndU8BBf62eM9prgo4coie5ttZrvS8GBzddzrhD > doc/wif.txt

[wif output text](https://github.com/gotamer/hdkeys/tree/master/doc/wif.txt)


### json key set
    $ hdkeys keyset -pass "" -mnemonic "leader monkey parrot ring guide accident before fence cannon height naive bean" > doc/keyset.json

[keyset output json](https://github.com/gotamer/hdkeys/tree/master/doc/keyset.json)


For Coders
----------
[Quick start link](https://github.com/gotamer/hdkeys/tree/master/doc/coders.md)


Base code stolen from:
----------------------

[hdkeygen](https://github.com/modood/hdkeygen) with same license Thank you [modood](https://github.com/modood)

