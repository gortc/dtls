package dtls

import (
	"net"
	"sync"
	"time"
)

type connKey struct {
	IP   [net.IPv6len]byte
	Port int
}

type packetConn struct {
	Conn
	R net.Conn
}

// PacketServer implements DTLS wrapping the PacketConn.
type PacketServer struct {
	conn net.PacketConn

	mutex sync.Mutex
	conns map[connKey]*packetConn
}

func Listen() (*PacketServer, error) {
	// pass
	return &PacketServer{}, nil
}

type udpHalfConn struct {
	conn   net.PacketConn
	remote net.Addr

	remoteR, remoteL net.Conn
}

func (c *udpHalfConn) Read(b []byte) (n int, err error) {
	return
}

func (udpHalfConn) Write(b []byte) (n int, err error) {
	panic("implement me")
}

func (udpHalfConn) Close() error {
	panic("implement me")
}

func (udpHalfConn) LocalAddr() net.Addr {
	panic("implement me")
}

func (udpHalfConn) RemoteAddr() net.Addr {
	panic("implement me")
}

func (udpHalfConn) SetDeadline(t time.Time) error {
	panic("implement me")
}

func (udpHalfConn) SetReadDeadline(t time.Time) error {
	panic("implement me")
}

func (udpHalfConn) SetWriteDeadline(t time.Time) error {
	panic("implement me")
}

func (s *PacketServer) Read(buf []byte) error {
	n, addr, err := s.conn.ReadFrom(buf)
	if err != nil {
		return err
	}

	var k connKey
	switch a := addr.(type) {
	case *net.UDPAddr:
		copy(k.IP[:], a.IP)
		k.Port = a.Port
	default:
		panic("unexpected addr type")
	}

	s.mutex.Lock()
	conn, ok := s.conns[k]
	if !ok {
		l, r := net.Pipe()
		conn = &packetConn{
			Conn: Conn{
				conn: l,
			},
			R: r,
		}
		s.conns[k] = conn
		go func() {
			rbuff := make([]byte, 1024)
			for {
				n, err := r.Read(rbuff)
				if err != nil {
					return
				}
				_, _ = s.conn.WriteTo(rbuff[:n], addr)
			}
		}()
	}
	s.mutex.Unlock()

	_, err = conn.R.Write(buf[:n])
	return err
}
