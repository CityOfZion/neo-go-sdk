package neo_test

var (
	testAccounts = []struct {
		privateKey       string
		privateKeyBase64 string
		publicAddress    string
		publicKey        string
		signature        string
		wif              string
	}{
		{
			privateKey:       "7d128a6d096f0c14c3a25a2b0c41cf79661bfcb4a8cc95aaaea28bde4d732344",
			privateKeyBase64: "fRKKbQlvDBTDolorDEHPeWYb/LSozJWqrqKL3k1zI0Q=",
			publicAddress:    "ALq7AWrhAueN6mJNqk6FHJjnsEoPRytLdW",
			publicKey:        "02028a99826edc0c97d18e22b6932373d908d323aa7f92656a77ec26e8861699ef",
			signature:        "3775292229eccdf904f16fff8e83e7cffdc0f0ce",
			wif:              "L1QqQJnpBwbsPGAuutuzPTac8piqvbR1HRjrY5qHup48TBCBFe4g",
		},
		{
			privateKey:       "dc5e273f370113018217c876056011dad0c897f6eca074bf741807f35ed271d2",
			privateKeyBase64: "3F4nPzcBEwGCF8h2BWAR2tDIl/bsoHS/dBgH817ScdI=",
			publicAddress:    "AVzgMjviERgZSCVoerzaGYhZhKoecd9RXk",
			publicKey:        "02c39609dc92eba0221809f0420e80728649c2515f95536288f848e324751cf738",
			signature:        "9bfdbeb5ac91a220742b1dac66b402e8b7cd9107",
			wif:              "L4c5RVHp3FyG7duisGHpSCBzEqnTUqvYw4Bv6hnwCgbDTsLYgt9o",
		},
		{
			privateKey:       "07fbee0481bb32441e9fca5d52c66e2554c06591c7f747e5d77bcb6217d1952d",
			privateKeyBase64: "B/vuBIG7MkQen8pdUsZuJVTAZZHH90fl13vLYhfRlS0=",
			publicAddress:    "AN2SiiLndiLsX9sYyVYmn3LYyjgozfUnb4",
			publicKey:        "038d92ca035622f04be0c6d23a65bcad04ce638c4ebd18b87facbd101323d349f9",
			signature:        "44922f93bf285dfc5199f28a028e82387d869306",
			wif:              "KwVEM6nK44RvgsohK5zxYAgRbut9mfGYaTcb7jkqBBfnApCLRJys",
		},
	}

	testBlocks = []struct {
		hash                 string
		id                   string
		index                int64
		merkleRoot           string
		numberOfTransactions int
	}{
		{
			hash:                 "0x3f0b498c0d57f73c674a1e28045f5e9a0991f9dac214076fadb5e6bafd546170",
			id:                   "1",
			index:                316675,
			merkleRoot:           "0xd51ef6237173eee1d422811c2e79d2e30928ed7c487ff9be60c493a9901b03b8",
			numberOfTransactions: 2,
		},
	}

	testTransactions = []struct {
		hash string
		id   string
		size int64
	}{
		{
			hash: "0x8aaf766179c07941f24a08157d7e6796e6d4aa999d3eaf83e74024c28d081af0",
			id:   "1",
			size: 262,
		},
	}

	testTransactionOutputs = []struct {
		asset string
		hash  string
		id    string
		index int64
		value string
	}{
		{
			asset: "0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7",
			hash:  "0x8aaf766179c07941f24a08157d7e6796e6d4aa999d3eaf83e74024c28d081af0",
			id:    "1",
			index: 0,
			value: "50",
		},
	}

	testBlockHashes = []struct {
		id    string
		hash  string
		index int64
	}{
		{
			id:    "1",
			index: 316675,
			hash:  "0x3f0b498c0d57f73c674a1e28045f5e9a0991f9dac214076fadb5e6bafd546170",
		},
	}

	testPings = []struct {
		description string
		uri         string
	}{
		{
			description: "InvalidURI",
			uri:         ")£*&%(£*&Q",
		},
		{
			description: "OfflineURI",
			uri:         "/foo",
		},
	}
)
