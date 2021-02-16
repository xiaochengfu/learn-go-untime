package main

import (
	"crypto/tls"
	"log"
	"time"
)

type Server struct {
	Addr       string
	Port       int
	Protocol   string
	Timeout    time.Duration
	MaxConnect int
	TLS        *tls.Config
}

//定义一个函数类型
type Option func(server *Server)

//定义一个函数，返回一个函数，也叫高阶函数
func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(t time.Duration) Option {
	return func(s *Server) {
		s.Timeout = t
	}
}

func MaxConnect(m int) Option {
	return func(s *Server) {
		s.MaxConnect = m
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(s *Server)) (*Server, error) {
	srv := Server{
		Addr:       addr,
		Port:       port,
		Protocol:   "tcp",
		Timeout:    30 * time.Second,
		MaxConnect: 1000,
		TLS:        nil,
	}
	for _, option := range options {
		log.Printf("%v", &srv)
		option(&srv)
	}

	return &srv, nil
}

func main() {
	s1, _ := NewServer("127.0.0.1", 8000, Protocol("udp"))
	s2, _ := NewServer("localhsot", 9000, Timeout(300*time.Second), MaxConnect(10), TLS(nil))
	log.Println(s1, s2)
}

//结论：以上使用了函数式编程
//类似的选项配置，可以使用以上方法
