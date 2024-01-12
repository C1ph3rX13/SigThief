package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Binject/debug/pe"
)

type Opts struct {
	Target string
	Cert   string
	Save   string
	Thief  string
	Output string
}

func banner() {
	fmt.Println(`
██████  ██▓  ▄████  ███▄    █ ▄▄▄█████▓ ██░ ██  ██▓▓█████   █████▒
▒██    ▒ ▓██▒ ██▒ ▀█▒ ██ ▀█   █ ▓  ██▒ ▓▒▓██░ ██▒▓██▒▓█   ▀ ▓██   ▒ 
░ ▓██▄   ▒██▒▒██░▄▄▄░▓██  ▀█ ██▒▒ ▓██░ ▒░▒██▀▀██░▒██▒▒███   ▒████ ░ 
  ▒   ██▒░██░░▓█  ██▓▓██▒  ▐▌██▒░ ▓██▓ ░ ░▓█ ░██ ░██░▒▓█  ▄ ░▓█▒  ░ 
▒██████▒▒░██░░▒▓███▀▒▒██░   ▓██░  ▒██▒ ░ ░▓█▒░██▓░██░░▒████▒░▒█░    
▒ ▒▓▒ ▒ ░░▓   ░▒   ▒ ░ ▒░   ▒ ▒   ▒ ░░    ▒ ░░▒░▒░▓  ░░ ▒░ ░ ▒ ░    
░ ░▒  ░ ░ ▒ ░  ░   ░ ░ ░░   ░ ▒░    ░     ▒ ░▒░ ░ ▒ ░ ░ ░  ░ ░      
░  ░  ░   ▒ ░░ ░   ░    ░   ░ ░   ░       ░  ░░ ░ ▒ ░   ░    ░ ░    
      ░   ░        ░          ░           ░  ░  ░ ░     ░  ░        

@Auth: C1ph3rX13
@Blog: https://c1ph3rx13.github.io
@Note：SigThief - Go
@Warn: 代码仅供学习使用，请勿用于其他用途
	`)
}

func flagOpts() *Opts {
	options := &Opts{}
	flag.StringVar(&options.Target, "t", "", "Target (Unsign) Exe")
	flag.StringVar(&options.Cert, "c", "", "Cert File")
	flag.StringVar(&options.Thief, "f", "", "Signed Exe")
	flag.StringVar(&options.Save, "s", "", "Save Cert")
	flag.StringVar(&options.Output, "o", "", "Output Signed File / Cert")
	flag.Parse()
	return options
}

func saveCert(thief string, dstCert string) error {
	cert, err := getCert(thief)
	if err != nil {
		return fmt.Errorf("getCert() err: %v", err)
	}

	err = os.WriteFile(dstCert, cert, os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.WriteFile() err: %v", err)
	}

	return nil
}

func getCert(signed string) ([]byte, error) {
	peFile, err := pe.Open(signed)
	if err != nil {
		return nil, fmt.Errorf("pe.Open() err: %v", err)
	}
	defer peFile.Close()

	if len(peFile.CertificateTable) == 0 {
		return nil, fmt.Errorf("CertificateTable is empty")
	}

	return peFile.CertificateTable, nil
}

func sigThief(unSigned string, signed string, thief []byte) error {
	peFile, err := pe.Open(unSigned)
	if err != nil {
		return fmt.Errorf("pe.Open() err: %v", err)
	}
	defer peFile.Close()

	peFile.CertificateTable = thief

	err = peFile.WriteFile(signed)
	if err != nil {
		return fmt.Errorf("pe.WriteFile() err: %v", err)
	}

	return nil
}

func certThief(unSigned string, signed string, thief string) error {
	thiefTable, err := os.ReadFile(thief)
	if err != nil {
		return fmt.Errorf("os.ReadFile() err: %v", err)
	}

	err = sigThief(unSigned, signed, thiefTable)
	if err != nil {
		return fmt.Errorf("sigThief() err: %v", err)
	}

	return nil
}

func exeThief(unSigned string, signed string, thief string) error {
	peFile, err := pe.Open(thief)
	if err != nil {
		return fmt.Errorf("pe.Open() err: %v", err)
	}
	defer peFile.Close()

	err = sigThief(unSigned, signed, peFile.CertificateTable)
	if err != nil {
		return fmt.Errorf("sigThief() err: %v", err)
	}

	return nil
}

func main() {
	banner()

	options := flagOpts()

	if options.Output == "" {
		fmt.Println("[*] Use -h to get help")
		fmt.Println("[*] Output path is required")
		return
	}

	if options.Target != "" && options.Cert != "" {
		err := certThief(options.Target, options.Output, options.Cert)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[*] Cert Signed: %s", options.Output)
	} else if options.Target != "" && options.Thief != "" {
		err := exeThief(options.Target, options.Output, options.Thief)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[*] SigThief: %s", options.Output)
	} else if options.Save != "" {
		err := saveCert(options.Save, options.Output)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[*] Saved: %s", options.Output)
	} else {
		log.Fatal("Invalid combination of options")
	}
}
