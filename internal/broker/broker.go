package broker

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"goqueue/pkg/eventbus"
	"goqueue/pkg/protocol"
	"net"
	"net/rpc/jsonrpc"
)

type BrokerCfg struct {
	ListenAddr string
	ListenPort string
}

type Broker struct {
	events         chan eventbus.Event
	logger         *logrus.Logger
	cfg            BrokerCfg
	stopBrokerChan chan struct{}
}

func (b *Broker) OnEvent(event eventbus.Event) {
	b.events <- event
}

func NewBroker(config BrokerCfg, logger *logrus.Logger) *Broker {
	return &Broker{cfg: config, logger: logger, stopBrokerChan: make(chan struct{}, 0)}
}

func (b *Broker) Run() error {
	listenAddr := b.cfg.ListenAddr + ":" + b.cfg.ListenPort
	b.logger.Debugf("Broker listen on %s", listenAddr)
	l, err := net.Listen("tcp", listenAddr) // TODO: get tls connection
	if err != nil {
		return err
	}
	defer l.Close()
	clients := make([]*net.Conn, 0)

	err = RegisterCommands()
	if err != nil {
		return err
	}

	go func() {
	OUTER:
		for {
			select {
			case <-b.stopBrokerChan:
				break OUTER
			default:
				clientConn, err := l.Accept()
				b.logger.Infof("Accept new client %s: ", clientConn.RemoteAddr().String())
				if err != nil {
					continue
				}
				clients = append(clients, &clientConn)
				jsonrpc.ServeConn(clientConn)
			}
		}
	}()
	<-b.stopBrokerChan
	err = b.sendCloseMessageForAllClient(clients)
	if err != nil {
		return err
	}

	return nil
}

func (b *Broker) sendCloseMessageForAllClient(clients []*net.Conn) error {
	var endMsgBuilder bytes.Buffer
	endMsgBuilder.WriteString(protocol.CLIENT_END_MSG)

	b.logger.Debugf("Number of client to close: %d", len(clients))
	for _, c := range clients {
		_, err := (*c).Write(endMsgBuilder.Bytes())
		if err != nil {
			b.logger.Error(err)
		}
	}
	return nil
}
