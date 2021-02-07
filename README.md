# Bippy
Golang implementation of the BIP32/BIP39/BIP43/BIP44/SLIP44/BIP49/BIP84/BIP173 for creating and recoverying keys, mnemonic seeds and Hierarchical Deterministic (HD) addresses. Includes support for BTC and ETH.

Lovingly ripped off from [@modood](https://github.com/modood) via https://github.com/modood/btckeygen.

# Ledger Recovery

Bippy has been verified to recover keys from Ledger devices with only the passphrase. Despite the marketing, you do not have to purchase a new Ledger device to recover your keys. See example [below](https://github.com/jlogelin/bippy#recover-account-keys-from-passphrase).

# Don't Trust. Verify.

I recommend every user of this library audit and verify any underlying code for its validity and suitability.

You can do so by using this tool: https://iancoleman.io/bip39/

or the many existing implementations described in the bip39 document: https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki#reference-implementation

I would also recommend running this command from a cold system, disconnected from the Internet, as it explicitely prints private key information in the terminal.

# Releases

Prebuilt binaries can be found in the releases section: https://github.com/jlogelin/bippy/releases.

  * [bippy for mac](https://github.com/jlogelin/bippy/releases/download/v0.0.3/bippy-0.0.3-darwin-amd64)
  * [bippy for windows](https://github.com/jlogelin/bippy/releases/download/v0.0.3/bippy-0.0.3-windows-amd64.exe)
  * [bippy for linux](https://github.com/jlogelin/bippy/releases/download/v0.0.3/bippy-0.0.3-linux-amd64)

Once you download the binary, open up a terminal and start using it.

# Help
```
% bippy -help

Usage of ./bippy:
  -bip39
    	mnemonic code for generating deterministic keys
  -n int
    	set number of keys to generate (default 10)
  -pass string
    	protect bip39 mnemonic with a passphrase
  -phrase string
    	set the key phase mnemonic
  -s	Use short derivation paths, as used with Ledger devices (ie. m/44'/60'/0'/0)
  -size int
    	Key phase size, valid values are 12 or 24 (default 24)
```

# Examples

## Recover Account Keys from Passphrase

Display Ethereum and Bitcoin keys from a 12 or 24 word passphrase. This can be used with compliant software and hardware wallets, including Ledger, to recover keys.

Example Ledger recovery; must include `-s` flag:

```
% bippy -bip39 -s -n 10 -phrase "echo cool vapor illness drastic citizen damp nurse labor rocket tool verb tower position duck endless tourist struggle ten firm scissors pilot own crouch"

BIP39 Mnemonic:    echo cool vapor illness drastic citizen damp nurse labor rocket tool verb tower position duck endless tourist struggle ten firm scissors pilot own crouch
BIP39 Passphrase:  <none>
BIP39 Seed:        ea1cdd5aa09e7760ceab85e1bb24c61769785c5dc09e44322d0cde3b54056e9326a697a06ca33c28b5dd7ab7d2f041aeeadf7a5d78e706e455be5dae93a92fec
BIP32 Root Key:    xprv9s21ZrQH143K3ByxkcZXiiG7aejrq3ioFQ6yNC4bDf7BJk2aHRpocGtk4Y9XV2M7wFLdUBs26PnW83tRu7tzBa7EKeSZG1iLW12hfN8s52v

Path(BIP44)        Ethereum Address                           Private Key                                 
------------------------------------------------------------------------------------------------------------------------------
m/44'/60'/0'/0     0xE720c0704b808b320F76A943a649d094Da2f0251 a592f34fccfa234d14ba2b550b1f30ed2df735802fe4c98ac60877c57808c404
m/44'/60'/0'/1     0xED31ADB49171B50c2e2fEDd14D810A99De024A2d c2d058624c30662b2a3e43125dedd95f4452ce6603669f93d8cb67d4b3546d6c
m/44'/60'/0'/2     0xe62EA82e64f2c3aF43DBAE0b7fC5826f0C189130 856bafc2361a308508774b7bf273c448c6d08c772d6b1f05e9b5a55de4eefb24
m/44'/60'/0'/3     0xB8555F45749aD194Ec0B3c6723B23cE038c78F21 f3e0c4bfc0715f4ddd85080fb52cfebf1eb86928cbfb6a780a0b55b4473dde16
m/44'/60'/0'/4     0x7A966345BA85C11477C84691249FD6Ba2f441ab6 c9b5dbd7f1f1a5b80e1096d90b9306643e5bc3c9df9acde3fbcb9b60982b741a
m/44'/60'/0'/5     0xc7c3C51571ac58b536d3204b66914A41231546f9 1f31c37e85e8a1207ec3d7453050bc18253ec223d4443ff198b705a264eac387
m/44'/60'/0'/6     0x6CD77fee3cd9e05D806Df15Dcb7a72Cc89aFbBdc 352b4a26ad9ca2b189c63b666ca99001540fd2b3a15b544c94cb4f7186ee4a6c
m/44'/60'/0'/7     0x23e7A1124e104609AD814cEFFe2735af022F3454 2dccb49aa08fecd0ef90f540ad5cadfb5c0c4b18df4a03d5353768dd23585166
m/44'/60'/0'/8     0x63d4543e3bC2CAC5b61A7C0754B4691722E25526 8396baa82e3ffe642aea226c8d8b67889aa6e58f925bc0d43306ff0529063a50
m/44'/60'/0'/9     0x8dcf09d5770E98BdD47E8d4aeC022a95ba05Aed9 13c0cf21e3b1d3f0faecf3fd520de03ef415611e362f61da988fbadc133a56a3

Path(BIP44)        Bitcoin Address                    WIF(Wallet Import Format)                            Private Key
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------
m/44'/0'/0'/0      1LG9h9Gv5W85FwQWuoHTDTDf829LbLrGKD Kwt2A54cB2dgkAmNbXfvTtS6HZ7uvcusJs6a6AuQw59mKENm71Ns 13b542bc3fd663529fedae308bcb7ca990b755b5935f1bffaeb69c67b87dfd7a
m/44'/0'/0'/1      1FEUE2jKAm4feufV25d3yLAHmNe41w8vFa L4MLfNveQbHg22MoarurFde5sHxiPL3BS4tP827n48rzCmGdci8C d4c950a3c075d3970ffecd8e49067adcee4c2e68135a82e2a8de75b8c5535b45
m/44'/0'/0'/2      1KfbTGRvTxM3FQFsv8P5bRxbx582bUxT1K Kyg7notEhHramYRf6f285TJozRqKQoq4zDCS7zCfGAy6B3fNXPsj 4942722abceeec29aa9bef1796b2decc40b3edca290e729c1517f7aa168fba7d
m/44'/0'/0'/3      19dSUJttTPFru4egjxVyvtxx1UBiSz1q9R KxtMPBqMF7cqeuoHmmZvEcVrdBYBnqQUQzEWKcCQv43tsb8X5Ytc 31b74a1f58230eb3360938758a4b1230d7fc920b1968d597abc00022660795d8
m/44'/0'/0'/4      1Q7ZP2nML7bx3RkUxyEjYaDNLeLBgEnUph L4asiptTK52TiD6dJVndz1EhSW27YuDPLNG6HkLknoUnrbzuGyCU dbbfe374bbee6da16fa212029cbb6749c3ead6ef163994a8301ec8d791a5bbc1
m/44'/0'/0'/5      18JkTosA3GKN2hUvzfBhjq3iFPzx8VGiUS L3K1PfqXz8uaKfgjrhdEfrE26KC7r8ZQSMcUqcchhi5hLSo6G2Nw b5bfcae0bfe543c80986e9d91aed420d88d4f92ec23f73a99482dc7afe67a860
m/44'/0'/0'/6      1CRJP6Ck8f1uXuYLDsbmNGz3t8dvZcYnNd L5fh9L7RZ5AKbXvwn8jV5baE9JkgaSEQUnJPsNzieTDX4o3i6jE6 fc10b9967a577e1a591f10c342cf452f3a80c35874c1635e467195e0a3b0b00c
m/44'/0'/0'/7      1C3awP2kYY9GHXyGJsvC6grFab55P6DSFD KyeDBxeGLXXtUn5KYhxxeXXzkzRy1cdqU4BC4yjnf12KzVeXTWCg 4847502efe46acfe0a0c3d92014204cd0635a7f598863a25f088979ef8266870
m/44'/0'/0'/8      1HZ6dzXCMeP7t38KUMrTu6k5NYaJRsk7pf KwbYREy8TdW9Y56dumCK3Nb5HoZ7yfwzkDHtiVhRejj6UbKeNZEi b3b23c022b5977509b03854dae45f92f9bbea8bed6a0fb33be32f6b19c73b7a
m/44'/0'/0'/9      125yo8ncwHpTXrLQexbtNpZ5DRKDqYHheF L1ESSPwydzh3eD7QGm6twmAzqhyTKJqmga6RYnK1xMXN8HLNYekj 77b96feed9329fceb4e7aae975ab619a40466ffba2958a044dd8548457854bc5

Path(BIP49)        SegWit(nested)                     WIF(Wallet Import Format)                            Private key
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------
m/49'/0'/0'/0      3GhedqaYchTqMQc6JgVXjetLHfm6A5LzYi L4JqpgGMHM3VrEqSJxm7RKVJ8eXSCwWGiaMA6W5ubV5QirsGyWV3 d380705557db86ed4fac49bb306f49fac6ac2f4210b00427391a9a7b8a4947dd
m/49'/0'/0'/1      35J6TabXYBeCd61AfPX7igRadfmC8yM1Zt L16sbk66Mcwj8PH8yG6KUUAXheZQmQQArm7W5xYB7JPVqGUvbaRP 73d4fffdc5e1402c2ae139d106bddaa6d0900d6d13a60f9e350f207af72bc213
m/49'/0'/0'/2      37oG3JrLwrxxCwKw5B5FwFt1HtgwaFaxwS L4iNaeGE3CzSw62d9WqFAywcbVziA621zEp8uXk7pttreeP2kgc2 df9b49f258d72e861bd32923fb3544f64274cc810a1d1d1bd5c7cc5ea4763929
m/49'/0'/0'/3      35KGWBDxDUMYwDGDoK2NyoLDnH1PqcNQks L35yQtHu92XjQ1YGSxowVHSnEHUod6cVxw1ccAaYYHJ75XB52zxD af0b3fe361f19f5fe7815eadc506738ead589a12221c72268874f5f7b7050038
m/49'/0'/0'/4      3AkJ7CwJpdRe4xV11yja4437Aq9qDzzhEW L3b2n8K9VuGcj429KETBh1s9fyKFdvXNwrwZqan9bZmNBEiRYFvb bdfe146c800142a4e52dad34dc6bc11addbc92391043a25f3e54c71050e664a3
m/49'/0'/0'/5      3GmQzfUJ75f5HJANMC4vUsJe2T9VdBBdWW L1xWGyFC9ysEtwn5eLEjWyerKRfRgoUZiCWPX5E5EwNVdaPbaNg7 8d5d630002370df34cbcbc6155a5199575111888f0a47a1628b0376f19662489
m/49'/0'/0'/6      3LgDmfuLt7xzbDhBJdwowf7LwsVUsjZEfg L39BzAZZZF15YZbF49CAWpKrxmGxtz2r5AyTvCynvYFQQinsvRJL b0b2e36166fa21aa5ff24e388e52c2c381cab4d3c38cd81f1b665951bc3ad779
m/49'/0'/0'/7      3M43MHUx9T5wDJeQCumENmpGxyj8PaqXdK L5AaWXKQApAsbaGe8L2atBzShGhGA1s2FjCSHp3y7gGxwFDrxJxv ed167938775647447822a28e176e7bfe3bdf0fdb44251c1865a9511e13a4223f
m/49'/0'/0'/8      32uPvZnHV9EBzz5nix2SFYGyDdvvs94yfb L15mqKsM6gAreaUKHyCsZVt2udEST3WiB88mXiVuEG9ZkDYaKnmw 7344363d1fc8f8213ca1571c8773dd8f0ef076acc68a58d3436d28d379e44d9e
m/49'/0'/0'/9      39BY5p5uJpBFK81FxZFeE1Ab7RTzyyMdPh L5BZMKDddDJ6FsuBHPnmiSnhMiHY2TYumyoz8gwv92viseG9rqCr ed9789da38a2fe60480ec707b10a0e8e4f8efc73e47ffcb084d372db3c69fb9c

Path(BIP84)        SegWit(bech32)                             WIF(Wallet Import Format)                            Private key
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
m/84'/0'/0'/0      bc1qwk8l9qptwh0ds0w88jnjrvjrf708yhw9f43j0g L5W7667WSQonx6jT41AHV5D6vr7z9kjNo3Prnn6gvJ6RgXfLioHo f72221cd682e0c2fe6e968aa4402e60d0447532dc4b154d238cb1763149e5a48
m/84'/0'/0'/1      bc1qxcgrzqcr5kxlmcqcspxdxewel2l8hmm3j26tk6 L5nvDjs6zkzUGrur8jEmTAvvQGAL4H18M7M6eTX4AWVNTh7Qxvnc ffc849d079e93c3f475fa1c585d2b390442eb5336467f2f8671f702405c04fae
m/84'/0'/0'/2      bc1qednrl23uct7q75cqdgklfm4n35su0u8ylz2j0n L38rFXVfBJeLQxxJJFzWCacrVqrRuQUh9MtijL41UifTA3QwztbV b08613bdf635ce0c8a9b9286ac794e1eafaa3598b464774feeace4dee5afb5d2
m/84'/0'/0'/3      bc1qe6xfdgfjuzgddgdkqaky9qpnr4aafsvq985xrg KzDrYgrRjJ4Q7SSPUSqZPmgkStUHq9baVKTHSu6Vh4VfGn4mtGx1 59961bddc3a3eb9b320f3657ace03d2af061fdae7f7437d2fa8d09358032013d
m/84'/0'/0'/4      bc1qr8l4f9vq59kmep88xktasdfl3m0k9xux2fkk5w L3EUBNEzMvMYPUZkZzhKruZ9Bqh5FSVgBWANUS6QbrNbev4Cfsx9 b36a231cc55232eb23d3d7f152cae57c83fcd1f884a130f08e02332f063d5155
m/84'/0'/0'/5      bc1q5k322qm3t8df8ykkcfg86cfp5rtnvcy0j9f4dm KwX7Rhit1nA1YBTsAb4m6Hu87gMMx2pFuSeEbXE2SbEWhpUzuwtm 8f39b96c7a44ded1392f43c57ce1d87501ef491c6bd771b8c7321d3cc3cf4e9
m/84'/0'/0'/6      bc1q8vsdqhgumj88p6f9sels99ycq9pdkxvwv5axss KzdSSfUA1NyW884TxziMRet7L97q867dDbUuaJkuefcwWh3fjnXf 65b81599cd3eb9471284fdc65bebfcbdb3db6fd9844f4c7ff84a19c7cff27252
m/84'/0'/0'/7      bc1qvdj4rl5ea08yrcc3lrz87ctszgufntsttay5y4 L2YXxFenq1ek8ywvDSGesUTVYgxmiPPAFcCMabB5vWTCqKVTfKcu 9ededcc9c8f6b26632aac36a4ca55be55b269b8ee42e3606efc63299c155cbf8
m/84'/0'/0'/8      bc1q9jl0z75plyj0mcz9at9v7am0d4x7a3e3jeemmn L49AJ5UnKLoXr1te86SHNsJKjXaMhiqFfyewZy2SqhWJSHCDUyvD ce856bf6c34cc9809c5086b1498f8f600648e6d264fe618931a17a05c4a7430c
m/84'/0'/0'/9      bc1q2z0f7zlu5t55rntnyj56g6tlndh7hz4q4ldty4 L2UTci6rfxjWWfQMSp7JEs3n7L6K7vj86e5Bkxso31E8eM2A8Lsc 9cc63aff9e4b4f3b2173c7dfaa312e03620f8100c800ef6c7ce7ba97d90be995
```


## Generate bip39 Mnemonic with Passphrase:

```
% bippy -bip39 -pass=123456 -n 10

BIP39 Mnemonic:    noodle cigar stamp awkward income worry must life wage canoe ancient harsh key indicate loud service ripple smile pledge flame rather exit theory share
BIP39 Passphrase:  123456
BIP39 Seed:        03c889b029068b4fcb1a122e242948f5b151533c19b7f550364c9745d446d66b2a470554ebc3661d29b840ffe1ea9a5c55f7c8a26caee47ff38c0fa3d59f93f2
BIP32 Root Key:    xprv9s21ZrQH143K466DS3XT9rjanyRZJGwsUAArqg2rDHRFe6tuiuH78ZyyVFdrvC1UDBdWibB6FWfrBUAm7NeodAMxBQfKzDrxZ8n8eq2nSMa

Path(BIP44)        Ethereum Address                           Private Key                                 
------------------------------------------------------------------------------------------------------------------------------
m/44'/60'/0'/0/0   0x0e2a51609C60682885F73E5AF71D8E03B4d73BD1 a017f32e8c95dd2e3c3c57ae4f9b6b9144cf1b026f44a55923fb9d2d3ece3b0d
m/44'/60'/0'/0/1   0x82f8f899A93769A6Ed6e7f233679EA40a8d3d948 9fd712c2f8305a41135abb4c54a63a3a6991595459e7eae206e5d8b66eb1ed2d
m/44'/60'/0'/0/2   0x82FcC7416B4539AdaCc0872439F46e08d52E6A81 8fbcebfaa2a4b2a7876e20438de1bea7c2a8cab66a4b46079a689dc3548950a8
m/44'/60'/0'/0/3   0x61D187d21ff1ACa7e25340859b098D0d10BDE05F a7cb23b5be3379d83ffd6996710081a533c4386332d9505855658964c2b8be0f
m/44'/60'/0'/0/4   0x8431027Bc0520732BAc69bd4F405B54d26C908aD c52285279197ff96b196438b242dad757b8782cf2d3ad1a011b560eaa3655492
m/44'/60'/0'/0/5   0x86dA8C7CE193A1161a53D7E90f9907cd629a7B15 9a6effc2b2084d22adaba35ba669668a032365715a5c9cea2fe2037734425c74
m/44'/60'/0'/0/6   0x409e5D8583C46f18733928BDD0b263ee24C60BEb aff4f6fb4ecc43abf2969ecdb68f692ad3f9dcedac5a4b058d100d673da41e1d
m/44'/60'/0'/0/7   0xee4235De6Cdf548e0920Bcb6ac4db87025A85549 37ac05546affc2da510406045eb49d844135b53f557b75243c4f9d68493a247
m/44'/60'/0'/0/8   0xF19eD0855F7933f445041E6405F1378CD7201d68 8d0cfcb706322f2bd880b67a4c234db3e06433fd82f3d156419293b491c84346
m/44'/60'/0'/0/9   0x8aB012fB1771b31409c82eB2237eB68D6dfd4E19 c45595c0390cd33c904022ac05c17e6e01a47a45fd7b4dd3e486588a0d2719a5

Path(BIP44)        Bitcoin Address                    WIF(Wallet Import Format)                            Private Key
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------
m/44'/0'/0'/0/0    1JSMsUq6uvKUsMHcEMzwxuX9wz4ezXUw6Q KyVozpddPshEwkRKHDYcpTbu6hKPVsKZ3di3HJc2DissFyDNM989 43f51587788f91bd945079048bc0b020ca34eeb89c4e94d77cc84bc4a00b7701
m/44'/0'/0'/0/1    1GaKguE5uCXngTvAEYLJfChppmjCNPhM4h Ky8XxfJ8TFCryP2ycbDnVbUKPSrgtGNtPtyBTpP3XS9SNfEqrEKs 39030d31bfe2414e87563c0d0ac8db20b4401578becabe2f4093829e66e1a93e
m/44'/0'/0'/0/2    1KHdZt998RypkDycjzJ83LkV7E3ZXWYdhe L2sMzgTu5JNyrPoUTLtFRtCAntULSV5ifnT417YMwSa4LSMLSZUj a88e7a2259ed4c8bb615f50d46746ab11f7d9aa288b05aad64b285a185582e96
m/44'/0'/0'/0/3    1727oVZGtMZRdSo7ngppPn7RmUBecy7rTc L49pGGsPPsqAmiKQmRhx3S46HR16TynBmBQGR2QF9K2TYH6ryzz3 cedba2807f0801461c633d5fe28bacdc6bd9ca7a2e257adc3942fd0d220f8bb3
m/44'/0'/0'/0/4    1EjdLb8ZXnNM8E1c5EFKXJLfXMo1MK5fW1 L48mheJngPw5Th9jVm5gV2u3bk9jTcGeUwh8zbcHWkkon6vkbuzF ce521eb8287962cdccceb10b0a794ee6ee910a7d77e2bf979fbecebebf5de204
m/44'/0'/0'/0/5    13zNMTZzTuu5E4jHsBqtCtAPMPcpwKfWid L1o8ghCHgdnbu7qKgQeX92JEpK7r9CjWFj91c4KmNFUXJsiWQGx9 888b18ceb5a20ede76b84a356a32dac498a184b63221b2a5da2dc94a3c2268ac
m/44'/0'/0'/0/6    12FLhAnRJxKcmr5ftQHhEv7Ayj9qk9ZveH L2VR34sC8oSeGoYwbQvucDguPoSk7tSMJkepjwkxgkLvRxohBpPP 9d44117a71b71f2e388384d5fbcd511116eb03105a71234be9790a61b5b3f29c
m/44'/0'/0'/0/7    1PpPovmZWaZmRqocxc89JtD17z4LPMsdJK L3idSHcK7MkV1BGvuipNUhDsB9Y4coDYkyUSF1kYwaKqHNNJ4XLZ c1e6a5c493eccae2f1b6b3e613825b9161b2bec34f7e445dea7bf8ab5b384270
m/44'/0'/0'/0/8    1J6Uwik7RZN7Qz7R3cwAHBQZ9fhpeWogS2 KxguGCt5LFJvHrv8Dm1VhKV6FYpRokz8aM3wuyJDpmZ1XamXRKaY 2bd352aeaac6792fc63e23ad714aebaf980897d64a061a89517c9e9ab393520e
m/44'/0'/0'/0/9    1Ke5ZHS2XmTHwcQyF3GcvZxFWUgvDHCn78 KzHR3x6JFZZRkVtNmsHJ8M7QfzcGEXB3xEA5vFryDtkhJsvMfqj7 5b6b00adb3406442fa0d73813076d8c9f21ec07d412c88c7db065992fa73fd2c

Path(BIP49)        SegWit(nested)                     WIF(Wallet Import Format)                            Private key
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------
m/49'/0'/0'/0/0    3CFUqJEV4Gc97W7JwGWixhkSDGqVfC3adC KzhrQVJxvAx7TzX7KDoG2JSt6dJQH8wz9Nf5g4jad6VXwrNSk9SU 67fd474e7b42ae26129ab6250e59a2255ef88a8fb169736f9596fe8222652142
m/49'/0'/0'/0/1    35UsLii8MLFmXPioG8yTSvgTFhiEDCHjM7 L3LknAxfHDFZnB8tq3gtPAA6UB6xhVn3xKEtr9V83uCJGkxJQmgw b6a6019b42be2725468c373d4fb6ccd510cd06f334c50fc35a5cccea78f5d5a6
m/49'/0'/0'/0/2    31jR7eu3RE5RRPNZxbzKUVB9bZQ9bD94GD L2Nzwsccdc85tq5zdcPopHqf5ATwqfMRyA7KZqqj94CnpqzzZP8L 99f73184904d3fe62272d7388686938c8ef5ff624985b79c05276fa663a4dd5a
m/49'/0'/0'/0/3    3D85vyS4FWp6Tynhf7xYwfiYb4sznbdUK7 L2C3ntqoRGZvSKuGExfJDpQ92saYUu5vB7AgqrzW5NCEGT6F1Wy9 9454ff2a38edbab2b2f74be2bb5fdf12300e9e6a4afaa5299e746ee28f4cc51c
m/49'/0'/0'/0/4    3DzCX6RQCo2bzvQxQWhRZ95smdb6i3vTKL KzXEgVEjw9FUULTf81xDsvMeg29jbcKuRbB2nrNsnuYcUaEsLhHk 628733966459c47e80f628ea34ec195e3f7cc9a2db7b955dbabc72dc720d2c3f
m/49'/0'/0'/0/5    3Mv8VqdpMvvKTFK2iE2UVk9rGyaFfWnr6n L3DD2RUN8Y2W3uuY6KZmHV6VkctdPPqfyZ8EdxG9D4CE8SL1uKNM b2c40814fa1dac409c5e925df643885c0e5d5830d0a5e337e1c93353552735a9
m/49'/0'/0'/0/6    36ZYXeMxkVZfVszqNh7q5eHiNmaUxCuVPD Kxnre8wNnEbCfZFuAPsiCiDjBvtHeBiBur2Gu5HSdM11BM63mDVw 2ee38bd9c50edcb156d3a14d0cd4396b49dabb82a3b6fadf431b48cecf77281d
m/49'/0'/0'/0/7    35RpAPTmsn3ArcSqyZiSj668jUHnX4u3EF KyXy5NXubweUhhEH5rQojtuSRpTG5pv1LkfuwzhGyBWwBDLJ2fT4 4511170b604253d33f4583845aec0049ddad6431757b8cd69d2299e0439dcf7d
m/49'/0'/0'/0/8    3GWTkv29o6qBgp6yhEvchRAMhrSJ6PvKvK L3Wc7HSYMPMAKBKYU89GvcboembricZNX6BjDMWvLAH6tgtqw4we bbb747a2b3b68fd2e2762b38bac4f0e1274a5dd3ed3b81410db79f484bea131e
m/49'/0'/0'/0/9    3Hx5mxPBSDrX8fXeo66RbRkXK88oMxp2Te L2VcS3Puhj3XD8Efio6R51hFun3jo1H7pPdbSrgBRWt8fYiNrL89 9d5df1cdd4a8fb2e91c1e78e27b2dda6ff18844d2c0b4d6171e20cf141372431

Path(BIP84)        SegWit(bech32)                             WIF(Wallet Import Format)                            Private key
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
m/84'/0'/0'/0/0    bc1qy79er2g8vv5lx5unyj23vggdpc9w84h9a043qv L27WLcsWCoXa465aQR6bitehoxBspwPhmsKYbwTW5hoVKFo8cuDo 91fecb5f125d7097cb905420a6d5bb2e0745599ec9f952fa8bb65dfff3a52de9
m/84'/0'/0'/0/1    bc1qjjjaycaewcwpjamf9lchkzcyk7h3ujnef7mkt4 L527t1UK3H6YKzzAXAqQCpwhNKWJeahSihsZD2KtfT9PRGg4V6n7 e8bc70675e4524421fe3aee5f181d41b242067357361c50a2b47b55e141ba39a
m/84'/0'/0'/0/2    bc1qu795xgh5nfzlskdsenm2s2ytes6mw4l6s93730 L21BrQPqJcKb632oPVB9jQgCdJYojp799neirA1ixFK8GZtyQgfq 8ebea45e20f339298b875118b371af9e35f03016160d94f7d60570f1844fe8d5
m/84'/0'/0'/0/3    bc1qt8e7wyu9napfavapmk8fh84xfasw00qr6d2au5 KwG8tAnPhuVxSFTJY85mE39k3np9LyWbixrc4tTcYzyRriq9HqGp 13f791c96d989a1e0ff72aa257a5299b5e4f62cf7b82c36b03017ea0dac3b48
m/84'/0'/0'/0/4    bc1q75whtrgx2924y9fe6lyk8x8u43qxkfh4j9qkxn L3yUDnQaPTyoynmM8rU7D9RvkmpnyFBendsfmT775LmBVgyeJMSj c98929f3094cbacbac2d561b0763bb0eaeeba44f016a57a4887561c50623fac1
m/84'/0'/0'/0/5    bc1qnq47ah6u9v9ed89ad9uydqhjr80e6lctv6yvy0 KwQXZXz1L8w6aDdzRGX6hXyNqwDakLSG9VMZgqumxNjZsZViLvar 590896224a7de14c38e2da2ae82283930b7cf097d0e034cdf6002c00a875687
m/84'/0'/0'/0/6    bc1q7zq8eurj3nfjlwu9rr5j0ep5r8jy4wpaghqcjk KyN6VFbavkq8r8DAqcwCusuD6efjcivyhnGB2DLYCuRpU2RSJ9jR 3ffcf5539047593f685bd997b220e7cefc496167f595a2c55c03904aec771b95
m/84'/0'/0'/0/7    bc1q0rdcf67m6f4t4jedc325x8suk0q950mq4w3y0e L3JZ1HE89LLHonFUipB1w5tC9bVwAo2QCvchx1r5Cq5nFYnv6EWr b583e139344baebeed1d20a25f38b2bc8f4ea9fc5a69513f8ab634103e0d520c
m/84'/0'/0'/0/8    bc1qnp26dvzmqp2vcu2unry59prfxsrtjzl034l5xg L49JNB6QZwYD7KbeBfnXxtYbpWzX8tx2cmyuvnqNNJUhKCVRm5yv ce97bf41bcf02b46e65b20c961c5581ad14fef5b44ed23f437f43d30686f6f9b
m/84'/0'/0'/0/9    bc1q9d3dmj54u9j6c4smm00gfn5vvehhntqjd60wms L2JZoxhV4rfsFpTCqxCEUzqMNXCdtPGM9Tik6SAUVsusM86ARB6V 97af556f709a3c74de0ed070ce15e631bf9afb772da72956bf4c7f7bc0e09018
```
