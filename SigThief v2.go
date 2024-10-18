package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Binject/debug/pe"
	"github.com/charmbracelet/log"
)

type SigThief struct {
	SignDir string // 指定目录
	SrcFile string // 未签名的源文件
	DstFile string // 已签名的目标文件
	Signed  string // 窃取签名后的输出文件
	DstCert string // 导出的证书文件
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
	@Warn: The code is for learning purposes only, please do not use it for other purposes
	`)
}

// CertificateTableReader 读取证书表
func CertificateTableReader(file string) ([]byte, error) {
	fileByte, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("<CertificateTableReader> failed to read file: %v", err)
	}

	fileReader := bytes.NewReader(fileByte)
	sourceBytes, err := pe.NewFile(fileReader)
	if err != nil {
		return nil, fmt.Errorf("<CertificateTableReader> failed to parse PE file: %v", err)
	}

	return sourceBytes.CertificateTable, nil
}

// SaveSignedFile 保存签名后的文件
func (CT *SigThief) SaveSignedFile(certTable []byte, fileName string) error {
	peFile, err := pe.Open(filepath.Join(CT.SignDir, CT.SrcFile))
	if err != nil {
		return fmt.Errorf("<SaveSignedFile pe.Open()> failed: %v", err)
	}
	defer peFile.Close()

	peFile.CertificateTable = certTable

	fileDir := filepath.Join(CT.SignDir, fileName)
	if err := peFile.WriteFile(fileDir); err != nil {
		return fmt.Errorf("<SaveSignedFile pe.WriteFile()> failed: %v", err)
	}

	log.Infof("saved signed to %s", fileDir)
	return nil
}

// CertSaver 保存证书到指定目录
func (CT *SigThief) CertSaver() error {
	certificateTable, err := CertificateTableReader(filepath.Join(CT.SignDir, CT.DstFile))
	if err != nil {
		return err
	}

	fileDir := filepath.Join(CT.SignDir, CT.DstCert)
	if err := os.WriteFile(fileDir, certificateTable, os.ModePerm); err != nil {
		return fmt.Errorf("<CertSaver os.WriteFile> failed: %v", err)
	}

	log.Infof("saved cert to %s", fileDir)
	return nil
}

// SigThief 使用窃取的证书进行签名
func (CT *SigThief) SigThief() error {
	DstCert, err := os.ReadFile(filepath.Join(CT.SignDir, CT.DstCert))
	if err != nil {
		return fmt.Errorf("<SigThief os.ReadFile> failed: %v", err)
	}

	return CT.SaveSignedFile(DstCert, CT.Signed)
}

// ExeThief 使用证书签名
func (CT *SigThief) ExeThief() error {
	dstTable, err := CertificateTableReader(filepath.Join(CT.SignDir, CT.DstFile))
	if err != nil {
		return err
	}

	return CT.SaveSignedFile(dstTable, CT.Signed)
}

func main() {
	banner()

	signDir := flag.String("dir", "", "指定目录")
	srcFile := flag.String("src", "", "未签名的源文件")
	dstFile := flag.String("dst", "", "已签名的目标文件")
	signed := flag.String("out", "", "窃取签名后的输出文件")
	dstCert := flag.String("cert", "", "导出的证书文件")
	action := flag.String("act", "", "操作类型 (save, exe, sign)")

	flag.Parse()

	sigThief := SigThief{
		SignDir: *signDir,
		SrcFile: *srcFile,
		DstFile: *dstFile,
		Signed:  *signed,
		DstCert: *dstCert,
	}

	var err error
	switch *action {
	case "save":
		if *dstFile == "" || *dstCert == "" {
			log.Fatalf("dst and cert parameters must be provided")
		}
		err = sigThief.CertSaver()
	case "sign":
		if *dstCert == "" || *srcFile == "" || *signed == "" {
			log.Fatalf("src, cert and out parameters must be provided")
		}
		err = sigThief.SigThief()
	case "exe":
		if *dstFile == "" || *srcFile == "" || *signed == "" {
			log.Fatalf("src, dst and out parameters must be provided")
		}
		err = sigThief.ExeThief()
	default:
		log.Error("Invalid operation type")
		os.Exit(1)
	}

	if err != nil {
		log.Errorf("Operation failed: %v", err)
		os.Exit(1)
	}
}
