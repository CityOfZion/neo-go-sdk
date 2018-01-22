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
			hash:                 "0x9ceb257832478ef778c1f26ad916ef8cbf116c71fe3cd6c5a8f672c1663539ae",
			id:                   "1",
			index:                1511369,
			merkleRoot:           "0x04bb7e7c56711b3387f1593c36dcdc36516b6ccd06d0e0c15adeba3c33643ebe",
			numberOfTransactions: 17,
		},
	}

	testTransactions = []struct {
		hash string
		id   string
		size int64
	}{
		{
			hash: "0xc515c4d2db27e06fd2305a5c5378f820d2c4cc04477ebe40ffa40b956eb4f8b5",
			id:   "1",
			size: 202,
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
			asset: "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
			hash:  "0x96fd0fc8a3cddbaac868647624c32cb8ac27f35cf249e4e6c8123601113d4017",
			id:    "1",
			index: 0,
			value: "2",
		},
	}

	testBlockHashes = []struct {
		id    string
		hash  string
		index int64
	}{
		{
			id:    "1",
			index: 1511369,
			hash:  "0x9ceb257832478ef778c1f26ad916ef8cbf116c71fe3cd6c5a8f672c1663539ae",
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
