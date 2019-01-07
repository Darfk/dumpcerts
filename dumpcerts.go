package main

import (
	"crypto/tls"
	"encoding/pem"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

var (
	flagError      = errors.New("failed to parse command line flags")
	noChainError   = errors.New("no verified chains exist")
	multChainError = errors.New("multiple verified chains exist")
)

type config struct {
	address  string
	protocol string
	verbose  bool
}

func init() {
	log.SetFlags(0)
}

func configure() (cfg *config, err error) {
	cfg = &config{}
	flag.StringVar(&cfg.protocol, "p", "tcp", "dial protocol")
	flag.BoolVar(&cfg.verbose, "v", false, "verbose")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.PrintDefaults()
		log.Println()
		return nil, flagError
	}
	cfg.address = flag.Args()[0]
	return
}

func dumpcerts(cfg *config) error {
	if cfg.verbose {
		log.Printf("connecting to %s\n", cfg.address)
	}
	conn, err := tls.Dial(cfg.protocol, cfg.address, nil)
	if err != nil {
		return err
	}
	if cfg.verbose {
		log.Printf("established a tls connection to %s\n", cfg.address)
	}
	defer conn.Close()

	state := conn.ConnectionState()
	if cfg.verbose {
		log.Printf("%d verified certificate chain(s) for %s\n", len(state.VerifiedChains), cfg.address)
	}
	if len(state.VerifiedChains) < 1 {
		return noChainError
	} else if len(state.VerifiedChains) > 1 {
		return multChainError
	}

	var outputWriter io.Writer = os.Stdout

	if cfg.verbose {
		log.Printf("writing pem to stdout\n")
	}

	for _, chain := range state.VerifiedChains {
		for _, crt := range chain {
			if cfg.verbose {
				log.Printf("write certificate for %s issued by %s", crt.Subject.CommonName, crt.Issuer.CommonName)
			}

			pem.Encode(outputWriter, &pem.Block{
				Type:  "CERTIFICATE",
				Bytes: crt.Raw,
			})
		}
	}
	return nil
}

func main() {
	var cfg *config
	var err error

	cfg, err = configure()
	if err != nil {
		log.Fatal(err)
	}

	err = dumpcerts(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
