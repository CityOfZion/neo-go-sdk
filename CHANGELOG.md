# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and this
project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## v1.9.0 - 2018-01-22

### Added

- New function to select the best node from a array of node URIs.

```golang
client, err := neo.NewClientUsingMultipleNodes(
  []string{
    "http://seed1.neo.org:10332",
    "http://seed2.neo.org:10332",
    "http://seed3.neo.org:10332",
  },
)

err = client.SelectBestNode()
```

## v1.8.0 - 2017-10-27

### Added

- New function to validate a public NEO address.

```golang
isValid, err := client.ValidateAddress("ARnLq3m1jsrZU7SS7jLHAvNm37GmaZbsPy")
```

## v1.7.0 - 2017-10-27

### Added

- New function to fetch a storage value of a given smart contract.

```golang
storage, err := client.GetStorage(
  "0xecc6b20d3ccac1ee9ef109af5a7cdb85706b1df9", // RPX smart contract
  "totalSupply",
)
```

## v1.6.1 - 2017-10-13

### Changed

- Updated client unit tests to adhere to the new JSON RPC hash representations.
- Performed some general cleanup and organization in existing client unit tests.

## v1.6.0 - 2017-09-01

### Added

- Ability to derive signature from private key.
- CLI outputs signature from WIF using new function.

```golang
privateKey, err := neo.NewPrivateKeyFromWIF("L1QqQJnpBwbsPGAuutuzPTac8piqvbR1HRjrY5qHup48TBCBFe4g")
if err != nil {
  log.Fatal(err)
}

signatureBytes, err := privateKey.Signature()
if err != nil {
  log.Fatal(err)
}

signature := hex.EncodeToString(signatureBytes)

log.Println(signature)
```

## v1.5.0 - 2017-09-01

### Added

- Ability to derive public key value from private key.
- CLI outputs public key from WIF using new function.

```golang
privateKey, err := neo.NewPrivateKeyFromWIF("L1QqQJnpBwbsPGAuutuzPTac8piqvbR1HRjrY5qHup48TBCBFe4g")
if err != nil {
  log.Fatal(err)
}

publicKeyBytes, err := privateKey.PublicKey()
if err != nil {
  log.Fatal(err)
}

publicKey := hex.EncodeToString(publicKeyBytes)

log.Println(publicKey)
```

## v1.4.1 - 2017-09-01

### Added

- GIF to README to help with documenting the new CLI.

## Changed

- Add `dist/` to **.gitignore**.

## v1.4.0 - 2017-08-19

### Added

- Documentation and configuration for the CLI.

## v1.3.0 - 2017-08-19

### Added

- CLI for debugging a NEO public and private key pair (see [releases](https://github.com/CityOfZion/neo-go-sdk/releases)):

```
./neo-go-sdk --wif KxQREAjBL6Ga8dw9rPN45pwoZ5dxhAQacEajQke6qmpB7DW6nAWE
```

## v1.2.0 - 2017-08-19

### Added

- [Elliptic curve](https://en.wikipedia.org/wiki/Elliptic_curve) implementation in utility package.
- Base58 encoding support.
- Migrated `neo.WIF` struct into `neo.PrivateKey` struct.
- Derive a public NEO address from a private key (WIF):

```golang
privateKey, err := neo.NewPrivateKeyFromWIF("L1QqQJnpBwbsPGAuutuzPTac8piqvbR1HRjrY5qHup48TBCBFe4g")
if err != nil {
  log.Fatal(err)
}

publicAddress, err := privateKey.PublicAddress()
if err != nil {
  log.Fatal(err)
}

log.Println(publicAddress)
```

## v1.1.2 - 2017-08-12

### Changed

- TCP connections made by `client.Ping()` are now closed to stop memory leaks from happening.

## v1.1.1 - 2017-08-12

### Changed

- **@eramus** fixed:
  - closing response body in wrong place.
  - made Base58 decode much more efficient and clean.
- Slack URI in README was updated.

## v1.1.0 - 2017-08-19

### Added

- New `neo.WIF` struct.
- New `neo.PrivateKey` struct.
- Ability to convert a WIF into a PrivateKey, see example below.

```golang
wif := neo.NewWIF("L1QqQJnpBwbsPGAuutuzPTac8piqvbR1HRjrY5qHup48TBCBFe4g")

privateKey, err := wif.ToPrivateKey()
if err != nil {
  log.Fatal(err)
}

log.Println(privateKey.Value)
```

## v1.0.1 - 2017-08-19

### Added

- Link to GoDoc documentation.

### Changed

- Logo in README to new CoZ logo.

## v1.0.0 - 2017-08-16

### Changed

- Added badges to README.

## v0.2.0 - 2017-08-16

### Added

- Existing code from original repo.
- Full CI job.

## v0.1.0 - 2017-08-16

### Added

- Setup repo.
