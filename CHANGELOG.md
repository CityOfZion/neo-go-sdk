# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and this 
project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

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
