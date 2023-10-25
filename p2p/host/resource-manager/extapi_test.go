package rcmgr

import (
	"encoding/json"
	"testing"

	"github.com/AstaFrode/go-libp2p/core/network"
	"github.com/AstaFrode/go-libp2p/core/peer"
	"github.com/stretchr/testify/require"
)

func TestResourceManagerStatRoundTrip(t *testing.T) {
	validPeerID, err := peer.Decode("QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN")
	require.NoError(t, err)

	n := ResourceManagerStat{
		Peers: map[peer.ID]network.ScopeStat{validPeerID: {}},
	}
	b, err := json.Marshal(n)
	require.NoError(t, err)

	rt := ResourceManagerStat{}
	require.NoError(t, json.Unmarshal(b, &rt))

	require.Equal(t, n, rt)
}
