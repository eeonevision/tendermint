package benchmarks

import (
	"testing"
	"time"

	"github.com/tendermint/tendermint/p2p"

	crypto "github.com/tendermint/go-crypto"
	wire "github.com/tendermint/go-wire"
	proto "github.com/tendermint/tendermint/benchmarks/proto"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func BenchmarkEncodeStatusWire(b *testing.B) {
	b.StopTimer()
	nodeKey := p2p.NodeKey{PrivKey: crypto.GenPrivKeyEd25519().Wrap()}
	status := &ctypes.ResultStatus{
		NodeInfo: p2p.NodeInfo{
			ID:         nodeKey.ID(),
			Moniker:    "SOMENAME",
			Network:    "SOMENAME",
			ListenAddr: "SOMEADDR",
			Version:    "SOMEVER",
			Other:      []string{"SOMESTRING", "OTHERSTRING"},
		},
		PubKey:            nodeKey.PubKey(),
		LatestBlockHash:   []byte("SOMEBYTES"),
		LatestBlockHeight: 123,
		LatestBlockTime:   time.Unix(0, 1234),
	}
	b.StartTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		jsonBytes := wire.JSONBytes(status)
		counter += len(jsonBytes)
	}

}

func BenchmarkEncodeNodeInfoWire(b *testing.B) {
	b.StopTimer()
	nodeKey := p2p.NodeKey{PrivKey: crypto.GenPrivKeyEd25519().Wrap()}
	nodeInfo := p2p.NodeInfo{
		ID:         nodeKey.ID(),
		Moniker:    "SOMENAME",
		Network:    "SOMENAME",
		ListenAddr: "SOMEADDR",
		Version:    "SOMEVER",
		Other:      []string{"SOMESTRING", "OTHERSTRING"},
	}
	b.StartTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		jsonBytes := wire.JSONBytes(nodeInfo)
		counter += len(jsonBytes)
	}
}

func BenchmarkEncodeNodeInfoBinary(b *testing.B) {
	b.StopTimer()
	nodeKey := p2p.NodeKey{PrivKey: crypto.GenPrivKeyEd25519().Wrap()}
	nodeInfo := p2p.NodeInfo{
		ID:         nodeKey.ID(),
		Moniker:    "SOMENAME",
		Network:    "SOMENAME",
		ListenAddr: "SOMEADDR",
		Version:    "SOMEVER",
		Other:      []string{"SOMESTRING", "OTHERSTRING"},
	}
	b.StartTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		jsonBytes := wire.BinaryBytes(nodeInfo)
		counter += len(jsonBytes)
	}

}

func BenchmarkEncodeNodeInfoProto(b *testing.B) {
	b.StopTimer()
	nodeKey := p2p.NodeKey{PrivKey: crypto.GenPrivKeyEd25519().Wrap()}
	id := &proto.ID{id: string(nodeKey.ID())}
	nodeInfo := proto.NodeInfo{
		Id:         id,
		Moniker:    "SOMENAME",
		Network:    "SOMENAME",
		ListenAddr: "SOMEADDR",
		Version:    "SOMEVER",
		Other:      []string{"SOMESTRING", "OTHERSTRING"},
	}
	b.StartTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		bytes, err := nodeInfo.Marshal()
		if err != nil {
			b.Fatal(err)
			return
		}
		//jsonBytes := wire.JSONBytes(nodeInfo)
		counter += len(bytes)
	}

}
