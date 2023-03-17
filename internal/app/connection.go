package app

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
)

type Connection struct {
	listener     net.Listener
	clientsCount int        // count of current connected clients
	clients      []net.Conn // current connected clients
}

type ConnectionOptions struct {
	UseTls       bool
	CaCertPath   string
	CertFilePath string
	KeyFilePath  string
	MaxClients   uint16
	IP           string
	Port         uint16
}

func (o ConnectionOptions) Validate() error {
	return nil
}

func MustNewConnection(opts ConnectionOptions) (Connection, error) {
	var conn Connection
	l, err := createListener(opts)
	if err != nil {
		return conn, fmt.Errorf("error create listener %w", err)
	}
	conn.listener = l

	return conn, nil
}

func (c *Connection) AcceptNewClient() (net.Conn, error) {
	conn, err := c.listener.Accept()
	if err != nil {
		return nil, err
	}
	c.clientsCount++
	c.clients = append(c.clients, conn)
	return conn, nil
}
func (c *Connection) ClientLeave() {
	c.clientsCount--
}

func createListener(opts ConnectionOptions) (net.Listener, error) {
	var listener net.Listener
	addrString := opts.IP + ":" + strconv.Itoa(int(opts.Port))
	if opts.UseTls {
		//TODO: use commented config below if mutual TLS will be implemented
		//tlsConfig := tls.Config{ClientCAs: ,ClientAuth: }
		cert, err := tryLoadCerts(opts.CaCertPath, opts.CertFilePath, opts.KeyFilePath)
		if err != nil {
			return nil, err
		}
		tlsConfig := tls.Config{Certificates: []tls.Certificate{cert}}
		listener, err = tls.Listen("tcp", addrString, &tlsConfig)
		if err != nil {
			return nil, err
		}
	}
	listener, err := net.Listen("tcp", addrString)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func tryLoadCerts(caFilePath string, certFilePath string, keyFilePath string) (tls.Certificate, error) {
	//TODO: load client CA cert for mutual tls
	return tls.LoadX509KeyPair(certFilePath, keyFilePath)
}
