package codec

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256r1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/sideprotocol/side/bitcoin/keys/segwit"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	var pk *cryptotypes.PubKey
	registry.RegisterInterface("cosmos.crypto.PubKey", pk)
	registry.RegisterImplementations(pk, &ed25519.PubKey{})
	registry.RegisterImplementations(pk, &segwit.PubKey{})
	registry.RegisterImplementations(pk, &multisig.LegacyAminoPubKey{})

	var priv *cryptotypes.PrivKey
	registry.RegisterInterface("cosmos.crypto.PrivKey", priv)
	registry.RegisterImplementations(priv, &segwit.PrivKey{})
	registry.RegisterImplementations(priv, &ed25519.PrivKey{})
	secp256r1.RegisterInterfaces(registry)
}
