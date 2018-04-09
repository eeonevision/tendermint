package benchmarks

import (
	"testing"
	"time"

	"github.com/tendermint/go-amino"
	"github.com/tendermint/go-crypto"

	proto "github.com/tendermint/tendermint/benchmarks/proto"
	"github.com/tendermint/tendermint/p2p"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func BenchmarkEncodeStatusWire(b *testing.B) {
	b.StopTimer()
	cdc := amino.NewCodec()
	ctypes.RegisterAmino(cdc)
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
		jsonBytes, err := cdc.MarshalJSON(status)
		if err != nil {
			panic(err)
		}
		counter += len(jsonBytes)
	}

}

func BenchmarkEncodeNodeInfoWire(b *testing.B) {
	b.StopTimer()
	cdc := amino.NewCodec()
	ctypes.RegisterAmino(cdc)
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
		jsonBytes, err := cdc.MarshalJSON(nodeInfo)
		if err != nil {
			panic(err)
		}
		counter += len(jsonBytes)
	}
}

func BenchmarkEncodeNodeInfoBinary(b *testing.B) {
	b.StopTimer()
	cdc := amino.NewCodec()
	ctypes.RegisterAmino(cdc)
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
		jsonBytes := cdc.MustMarshalBinaryBare(nodeInfo)
		counter += len(jsonBytes)
	}

}

func BenchmarkEncodeNodeInfoProto(b *testing.B) {
	b.StopTimer()
	nodeKey := p2p.NodeKey{PrivKey: crypto.GenPrivKeyEd25519().Wrap()}
	nodeID := string(nodeKey.ID())
	someName := "SOMENAME"
	someAddr := "SOMEADDR"
	someVer := "SOMEVER"
	someString := "SOMESTRING"
	otherString := "OTHERSTRING"
	nodeInfo := proto.NodeInfo{
		Id:         &proto.ID{Id: &nodeID},
		Moniker:    &someName,
		Network:    &someName,
		ListenAddr: &someAddr,
		Version:    &someVer,
		Other:      []string{someString, otherString},
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
