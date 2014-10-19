package http

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)
const RAW_HEADER_NAME="x-raw-header-name"
var DEBUG_PRINT_HEADER bool=false
func NewHttpServer(laddr string, handler http.Handler) error {
	ln, err := net.Listen("tcp", laddr)
	if err != nil {
		return err
	}
	listerProxy := &listenerProxy{
		originLister: ln,
		conn:         make(chan net.Conn, 100),
	}

	go http.Serve(listerProxy, handler)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go (func(conn net.Conn) {
			listerProxy.newConn(conn)
		})(conn)
	}
}

type listenerProxy struct {
	conn         chan net.Conn
	originLister net.Listener
}

func (l *listenerProxy) Accept() (net.Conn, error) {
	conn := <-l.conn
	return newConnProxy(conn)
}

func (l *listenerProxy) newConn(conn net.Conn) {
	l.conn <- conn
}

func (l *listenerProxy) Close() error {
	return l.originLister.Close()
}

func (l *listenerProxy) Addr() net.Addr {
	return l.originLister.Addr()
}

func newConnProxy(conn net.Conn) (*connProxy, error) {
	newConn := &connProxy{originConn: conn}
	var err error
	newConn.rawHeader, err = newRawHeader(bufio.NewReader(conn))
	if err != nil {
		return nil, err
	}
	return newConn, nil
}

type connProxy struct {
	originConn net.Conn
	rawHeader  *rawHttpHeader
	N          int
}

func (conn *connProxy) Read(b []byte) (n int, err error) {
	l := len(conn.rawHeader.fixedHeader)
	bufLen := len(b)
	if conn.N < l {
		for _, v := range conn.rawHeader.fixedHeader[conn.N:] {
			b[n] = v
			n++
			if n >= bufLen {
				break
			}
		}
		conn.N += n
		if n == l || n == bufLen {
			return n, nil
		}
	}
	if n != bufLen {
		m, _ := conn.originConn.Read(b[n:])
		n += m
		conn.N += m
	}

	return n, err
}

func (conn *connProxy) Write(b []byte) (n int, err error) {
	return conn.originConn.Write(b)
}
func (conn *connProxy) Close() error {
	return conn.originConn.Close()
}
func (conn *connProxy) LocalAddr() net.Addr {
	return conn.originConn.LocalAddr()
}
func (conn *connProxy) RemoteAddr() net.Addr {
	return conn.originConn.RemoteAddr()
}
func (conn *connProxy) SetDeadline(t time.Time) error {
	return conn.originConn.SetDeadline(t)
}
func (conn *connProxy) SetReadDeadline(t time.Time) error {
	return conn.originConn.SetReadDeadline(t)
}
func (conn *connProxy) SetWriteDeadline(t time.Time) error {
	return conn.originConn.SetWriteDeadline(t)
}

type rawHttpHeader struct {
	rawHeader   []byte
	fixedHeader []byte
}

func newRawHeader(reader *bufio.Reader) (*rawHttpHeader, error) {
	bf := make([]byte, http.DefaultMaxHeaderBytes)
	n, err := reader.Read(bf)
	if err != nil {
		return nil, err
	}
	datas:=strings.SplitN(string(bf[:n]),"\r\n\r\n",2)
	rawHeader := &rawHttpHeader{rawHeader: []byte(datas[0])}

	rawHeaderSlice := strings.Split(strings.TrimSpace(string(rawHeader.rawHeader)), "\r\n")
	names := make([]string, 0)
	for _, v := range rawHeaderSlice[1:] {
		tmp := strings.SplitN(v, ":", 2)
		if len(tmp) == 2 {
			names = append(names, strings.TrimSpace(tmp[0]))
		}
	}
	if(DEBUG_PRINT_HEADER){
		fmt.Println("rawHeader:\n",string(rawHeader.rawHeader),"\n\n")
	}
	headerNewBf := bytes.NewBuffer(rawHeader.rawHeader)
	headerNewBf.WriteString("\r\n")
	headerNewBf.WriteString(fmt.Sprintf("%s :%s\r\n",RAW_HEADER_NAME,strings.Join(names, "|")))
	headerNewBf.WriteString("\r\n")
	if(len(datas)==2){
		headerNewBf.WriteString(datas[1])
	}
	rawHeader.fixedHeader = headerNewBf.Bytes()
	return rawHeader, nil
}
