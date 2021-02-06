package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"

	crypto "github.com/ethereum/go-ethereum/crypto"
)

// Purpose BIP43 - Purpose Field for Deterministic Wallets
// https://github.com/bitcoin/bips/blob/master/bip-0043.mediawiki
//
// Purpose is a constant set to 44' (or 0x8000002C) following the BIP43 recommendation.
// It indicates that the subtree of this node is used according to this specification.
//
// What does 44' mean in BIP44?
// https://bitcoin.stackexchange.com/questions/74368/what-does-44-mean-in-bip44
//
// 44' means that hardened keys should be used. The distinguisher for whether
// a key a given index is hardened is that the index is greater than 2^31,
// which is 2147483648. In hex, that is 0x80000000. That is what the apostrophe (') means.
// The 44 comes from adding it to 2^31 to get the final hardened key index.
// In hex, 44 is 2C, so 0x80000000 + 0x2C = 0x8000002C.
type Purpose = uint32

const (
	PurposeBIP44 Purpose = 0x8000002C // 44' BIP44
	PurposeBIP49 Purpose = 0x80000031 // 49' BIP49
	PurposeBIP84 Purpose = 0x80000054 // 84' BIP84
)

// CoinType SLIP-0044 : Registered coin types for BIP-0044
// https://github.com/satoshilabs/slips/blob/master/slip-0044.md
type CoinType = uint32

const (
	CoinTypeBTC CoinType = 0x80000000
	CoinTypeLTC CoinType = 0x80000002
	CoinTypeETH CoinType = 0x8000003c
	CoinTypeEOS CoinType = 0x800000c2
)

const (
	Apostrophe uint32 = 0x80000000 // 0'
)

type Key struct {
	path     string
	bip32Key *bip32.Key
}

func (k *Key) Encode(compress bool) (wif, address, segwitBech32, segwitNested string, privateKey string, err error) {
	prvKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), k.bip32Key.Key)
	return GenerateFromBytes(prvKey, compress)
}

func (k *Key) EncodeEth() (address string, publicKey string, privateKey string) {
	prvKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), k.bip32Key.Key)
	pub := pubKey.ToECDSA()
	address = crypto.PubkeyToAddress(*pub).Hex()
	publicKey = string(crypto.FromECDSAPub(pub))
	privateKey = fmt.Sprintf("%x", prvKey.D)
	return address, publicKey, privateKey
}

// https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
// bip44 define the following 5 levels in BIP32 path:
// m / purpose' / coin_type' / account' / change / address_index

func (k *Key) GetPath() string {
	return k.path
}

type KeyManager struct {
	mnemonic   string
	passphrase string
	keys       map[string]*bip32.Key
	mux        sync.Mutex
}

// NewKeyManager return new key manager
// bitSize has to be a multiple 32 and be within the inclusive range of {128, 256}
// 128: 12 phrases
// 256: 24 phrases
func NewKeyManager(bitSize int, passphrase string) (*KeyManager, error) {
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	km := &KeyManager{
		mnemonic:   mnemonic,
		passphrase: passphrase,
		keys:       make(map[string]*bip32.Key, 0),
	}
	return km, nil
}

// NewKeyManager return new key manager
// bitSize has to be a multiple 32 and be within the inclusive range of {128, 256}
// 128: 12 phrases
// 256: 24 phrases
func NewKeyManagerWithMnemonic(bitSize int, passphrase string, keyphrase string) (*KeyManager, error) {

	entropy, err := bip39.EntropyFromMnemonic(keyphrase)
	if err != nil {
		return nil, err
	}

	/*
		valid := bip39.IsMnemonicValid(keyphrase)
		if valid == false {
			err := errors.New("mnemonic not valid")
			return nil, err
		}
	*/

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	km := &KeyManager{
		mnemonic:   mnemonic,
		passphrase: passphrase,
		keys:       make(map[string]*bip32.Key, 0),
	}
	return km, nil
}

func (km *KeyManager) GetMnemonic() string {
	return km.mnemonic
}

func (km *KeyManager) GetPassphrase() string {
	return km.passphrase
}

func (km *KeyManager) GetSeed() []byte {
	return bip39.NewSeed(km.GetMnemonic(), km.GetPassphrase())
}

func (km *KeyManager) getKey(path string) (*bip32.Key, bool) {
	km.mux.Lock()
	defer km.mux.Unlock()

	key, ok := km.keys[path]
	return key, ok
}

func (km *KeyManager) setKey(path string, key *bip32.Key) {
	km.mux.Lock()
	defer km.mux.Unlock()

	km.keys[path] = key
}

func (km *KeyManager) GetMasterKey() (*bip32.Key, error) {
	path := "m"

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	key, err := bip32.NewMasterKey(km.GetSeed())
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

func (km *KeyManager) GetPurposeKey(purpose uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'`, purpose-Apostrophe)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetMasterKey()
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(purpose)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

func (km *KeyManager) GetCoinTypeKey(purpose, coinType uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'`, purpose-Apostrophe, coinType-Apostrophe)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetPurposeKey(purpose)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(coinType)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

func (km *KeyManager) GetAccountKey(purpose, coinType, account uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'`, purpose-Apostrophe, coinType-Apostrophe, account)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetCoinTypeKey(purpose, coinType)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(account + Apostrophe)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

// GetChangeKey ...
// https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki#change
// change constant 0 is used for external chain
// change constant 1 is used for internal chain (also known as change addresses)
func (km *KeyManager) GetChangeKey(purpose, coinType, account, change uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'/%d`, purpose-Apostrophe, coinType-Apostrophe, account, change)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	parent, err := km.GetAccountKey(purpose, coinType, account)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(change)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

func (km *KeyManager) GetShortAccountKey(purpose uint32, coinType uint32, account uint32) (*bip32.Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d`, purpose-Apostrophe, coinType-Apostrophe, account)

	key, ok := km.getKey(path)
	if ok {
		return key, nil
	}

	key, err := km.GetAccountKey(purpose, coinType, account)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return key, nil
}

func (km *KeyManager) GetShortKey(purpose uint32, coinType uint32, account uint32, index uint32) (*Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'/%d`, purpose-Apostrophe, coinType-Apostrophe, account, index)

	key, ok := km.getKey(path)
	if ok {
		return &Key{path: path, bip32Key: key}, nil
	}

	parent, err := km.GetShortAccountKey(purpose, coinType, account)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(index)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return &Key{path: path, bip32Key: key}, nil
}

func (km *KeyManager) GetKey(purpose, coinType, account, change, index uint32) (*Key, error) {
	path := fmt.Sprintf(`m/%d'/%d'/%d'/%d/%d`, purpose-Apostrophe, coinType-Apostrophe, account, change, index)

	key, ok := km.getKey(path)
	if ok {
		return &Key{path: path, bip32Key: key}, nil
	}

	parent, err := km.GetChangeKey(purpose, coinType, account, change)
	if err != nil {
		return nil, err
	}

	key, err = parent.NewChildKey(index)
	if err != nil {
		return nil, err
	}

	km.setKey(path, key)

	return &Key{path: path, bip32Key: key}, nil
}

func Generate(compress bool) (wif, address, segwitBech32, segwitNested string, privateKey string, err error) {
	prvKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", "", "", "", err
	}
	return GenerateFromBytes(prvKey, compress)
}

func GenerateFromBytes(prvKey *btcec.PrivateKey, compress bool) (wif, address, segwitBech32, segwitNested string, privateKey string, err error) {
	// generate the wif(wallet import format) string
	btcwif, err := btcutil.NewWIF(prvKey, &chaincfg.MainNetParams, compress)
	if err != nil {
		return "", "", "", "", "", err
	}
	wif = btcwif.String()

	// generate a normal p2pkh address
	serializedPubKey := btcwif.SerializePubKey()
	addressPubKey, err := btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", "", "", "", err
	}
	address = addressPubKey.EncodeAddress()

	// generate a normal p2wkh address from the pubkey hash
	witnessProg := btcutil.Hash160(serializedPubKey)
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", "", "", "", err
	}
	segwitBech32 = addressWitnessPubKeyHash.EncodeAddress()

	// generate an address which is
	// backwards compatible to Bitcoin nodes running 0.6.0 onwards, but
	// allows us to take advantage of segwit's scripting improvments,
	// and malleability fixes.
	serializedScript, err := txscript.PayToAddrScript(addressWitnessPubKeyHash)
	if err != nil {
		return "", "", "", "", "", err
	}
	addressScriptHash, err := btcutil.NewAddressScriptHash(serializedScript, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", "", "", "", err
	}
	segwitNested = addressScriptHash.EncodeAddress()

	privateKey = fmt.Sprintf("%x", prvKey.D)

	return wif, address, segwitBech32, segwitNested, privateKey, nil
}

func main() {
	compress := true // generate a compressed public key
	bip39 := flag.Bool("bip39", false, "mnemonic code for generating deterministic keys")
	pass := flag.String("pass", "", "protect bip39 mnemonic with a passphrase")
	number := flag.Int("n", 10, "set number of keys to generate")
	mnemonic := flag.String("phrase", "", "set the key phase mnemonic")
	size := flag.Int("size", 24, "Key phase size, valid values are 12 or 24")

	flag.Parse()

	var bitsize int
	if *size == 24 {
		bitsize = 256
	} else if *size == 12 {
		bitsize = 128
	} else {
		log.Fatal("Invalid size, value must be 12 or 24")
		return
	}

	if !*bip39 {
		fmt.Printf("\n%-34s %-34s %-52s %-42s %s\n", "Bitcoin Address", "Private Key", "WIF(Wallet Import Format)", "SegWit(bech32)", "SegWit(nested)")
		fmt.Println(strings.Repeat("-", 165))

		for i := 0; i < *number; i++ {
			wif, address, segwitBech32, segwitNested, pk, err := Generate(compress)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%-34s %s %s %s %s\n", address, pk, wif, segwitBech32, segwitNested)
		}
		fmt.Println()
		return
	}

	var err error
	var km *KeyManager
	if *mnemonic != "" {
		km, err = NewKeyManagerWithMnemonic(bitsize, *pass, *mnemonic)
	} else {
		km, err = NewKeyManager(bitsize, *pass)
	}
	if err != nil {
		log.Fatal(err)
	}
	masterKey, err := km.GetMasterKey()
	if err != nil {
		log.Fatal(err)
	}
	passphrase := km.GetPassphrase()
	if passphrase == "" {
		passphrase = "<none>"
	}
	fmt.Printf("\n%-18s %s\n", "BIP39 Mnemonic:", km.GetMnemonic())
	fmt.Printf("%-18s %s\n", "BIP39 Passphrase:", passphrase)
	fmt.Printf("%-18s %x\n", "BIP39 Seed:", km.GetSeed())
	fmt.Printf("%-18s %s\n", "BIP32 Root Key:", masterKey.B58Serialize())

	fmt.Printf("\n** Ledger Derivation **")
	fmt.Printf("\n%-18s %-42s %-44s\n", "Path(BIP44)", "Ethereum Address", "Private Key")
	fmt.Println(strings.Repeat("-", 126))
	for i := 0; i < *number; i++ {
		key, err := km.GetShortKey(PurposeBIP44, CoinTypeETH, 0, uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		address, _, privateKey := key.EncodeEth()

		fmt.Printf("%-18s %-34s %s\n", key.GetPath(), address, privateKey)
	}

	fmt.Printf("\n%-18s %-42s %-44s\n", "Path(BIP44)", "Ethereum Address", "Private Key")
	fmt.Println(strings.Repeat("-", 126))
	for i := 0; i < *number; i++ {
		key, err := km.GetKey(PurposeBIP44, CoinTypeETH, 0, 0, uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		address, _, privateKey := key.EncodeEth()

		fmt.Printf("%-18s %-34s %s\n", key.GetPath(), address, privateKey)
	}

	fmt.Printf("\n%s %22s %44s %38s\n", "Path(BIP44)", "Bitcoin Address", "WIF(Wallet Import Format)", "Private Key")
	fmt.Println(strings.Repeat("-", 171))
	for i := 0; i < *number; i++ {
		key, err := km.GetKey(PurposeBIP44, CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		wif, address, _, _, pk, err := key.Encode(compress)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-18s %-34s %44s %s\n", key.GetPath(), address, wif, pk)
	}

	fmt.Printf("\n%-18s %-34s %24s %38s\n", "Path(BIP49)", "SegWit(nested)", "WIF(Wallet Import Format)", "Private key")
	fmt.Println(strings.Repeat("-", 171))
	for i := 0; i < *number; i++ {
		key, err := km.GetKey(PurposeBIP49, CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		wif, _, _, segwitNested, pk, err := key.Encode(compress)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-18s %s %s %s\n", key.GetPath(), segwitNested, wif, pk)
	}

	fmt.Printf("\n%-18s %-42s %s %38s\n", "Path(BIP84)", "SegWit(bech32)", "WIF(Wallet Import Format)", "Private key")
	fmt.Println(strings.Repeat("-", 179))
	for i := 0; i < *number; i++ {
		key, err := km.GetKey(PurposeBIP84, CoinTypeBTC, 0, 0, uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		wif, _, segwitBech32, _, pk, err := key.Encode(compress)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-18s %s %s %s\n", key.GetPath(), segwitBech32, wif, pk)
	}
	fmt.Println()
}
