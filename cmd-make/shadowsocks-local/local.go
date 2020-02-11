package main 

import (
	"log"
	"flag"	
        "os"
        ss "github.com/shadowsocks/shadowsocks-go/shadowsocks"
)


var debug ss.DebugLog



func main(){
	log.SetOutput(os.Stdout)

        var configFile, cmdServer, cmdURI string
	var cmdConfig ss.Config
	var printVer bool

        flag.BoolVar(&printVer, "version", false, "print version")
	flag.StringVar(&configFile, "c", "config.json", "specify config file")
	flag.StringVar(&cmdServer, "s", "", "server address")
	flag.StringVar(&cmdConfig.LocalAddress, "b", "", "local address, listen only to this address if specified")
	flag.StringVar(&cmdConfig.Password, "k", "", "password")
	flag.IntVar(&cmdConfig.ServerPort, "p", 0, "server port")
	flag.IntVar(&cmdConfig.Timeout, "t", 300, "timeout in seconds")
	flag.IntVar(&cmdConfig.LocalPort, "l", 0, "local socks5 proxy port")
	flag.StringVar(&cmdConfig.Method, "m", "", "encryption method, default: aes-256-cfb")
	flag.BoolVar((*bool)(&debug), "d", false, "print debug message")
	flag.StringVar(&cmdURI, "u", "", "shadowsocks URI")

	flag.Parse()

}


