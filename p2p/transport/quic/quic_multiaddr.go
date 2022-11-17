package libp2pquic

import (
	"errors"
	"net"

	"github.com/lucas-clemente/quic-go"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

var quicV1MA ma.Multiaddr = ma.StringCast("/quic-v1")
var quicDraft29MA ma.Multiaddr = ma.StringCast("/quic")

func toQuicMultiaddr(na net.Addr, version quic.VersionNumber) (ma.Multiaddr, error) {
	udpMA, err := manet.FromNetAddr(na)
	if err != nil {
		return nil, err
	}
	switch version {
	case quic.VersionDraft29:
		return udpMA.Encapsulate(quicDraft29MA), nil
	case quic.Version1:
		return udpMA.Encapsulate(quicV1MA), nil
	default:
		return nil, errors.New("unknown QUIC version")
	}
}

func fromQuicMultiaddr(addr ma.Multiaddr) (net.Addr, quic.VersionNumber, error) {
	var version quic.VersionNumber
	var partsBeforeQUIC []ma.Multiaddr
	ma.ForEach(addr, func(c ma.Component) bool {
		switch c.Protocol().Code {
		case ma.P_QUIC:
			version = quic.VersionDraft29
			return false
		case ma.P_QUIC_V1:
			version = quic.Version1
			return false
		default:
			partsBeforeQUIC = append(partsBeforeQUIC, &c)
			return true
		}
	})
	if len(partsBeforeQUIC) == 0 {
		return nil, version, errors.New("no addr before QUIC component")
	}
	if version == 0 {
		// Not found
		return nil, version, errors.New("unknown QUIC version")
	}
	netAddr, err := manet.ToNetAddr(ma.Join(partsBeforeQUIC...))
	if err != nil {
		return nil, version, err
	}
	return netAddr, version, err
}
