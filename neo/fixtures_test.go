package neo_test

var (
	testBlocks = []struct {
		hash                 string
		id                   string
		index                int64
		merkleRoot           string
		numberOfTransactions int
	}{
		{
			hash:                 "3f0b498c0d57f73c674a1e28045f5e9a0991f9dac214076fadb5e6bafd546170",
			id:                   "1",
			index:                316675,
			merkleRoot:           "d51ef6237173eee1d422811c2e79d2e30928ed7c487ff9be60c493a9901b03b8",
			numberOfTransactions: 2,
		},
	}

	testTransactions = []struct {
		hash string
		id   string
		size int64
	}{
		{
			hash: "8aaf766179c07941f24a08157d7e6796e6d4aa999d3eaf83e74024c28d081af0",
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
			asset: "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7",
			hash:  "8aaf766179c07941f24a08157d7e6796e6d4aa999d3eaf83e74024c28d081af0",
			id:    "1",
			index: 0,
			value: "50",
		},
	}
)
