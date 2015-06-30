package samjung

import (
	"bufio"
	"fmt"
	"os"
)

const (
	tableFile   = "samjung.table"
	indexFile   = "samjung.index"
	autoIncFile = "samjung.inc"
)

type Samjung struct {
	baseDir string
}

func New(baseDir string) *Samjung {
	return &Samjung{baseDir}
}

func isExistFiles(baseDir string) error {
	// make base directory
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		if err := os.MkdirAll(baseDir, 0755); err != nil {
			return fmt.Errorf("failed to make directory: err=%v", err)
		}
	}

	//	// make table file
	//	if _, err := os.Stat(baseDir + "/" + tableFile); os.IsNotExist(err) {
	//		// TODO : make table file
	//	}
	//
	//	if _, err := os.Stat(baseDir + "/" + tableFile); os.IsNotExist(err) {
	//		// TODO : make index file
	//	}

	return nil
}

func (r *Samjung) Start() {
	// init
	err := isExistFiles(r.baseDir)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// TODO : 파일이 없으면 테이블을 새로 만드는 과정을 거친 후에 작업이 시작되도록 구현

	for {
		fmt.Print("1. Insert 2. Select 3. Finish: ")
		c, err := r.readByte()
		if err != nil {
			fmt.Printf("failed to readByte: err=%v", err)
			return
		}

		switch c {
		default:
			fmt.Printf("Invalid value..")
		case '1':
			r.insertRow()
		case '2':
			r.selectRow()
		case '3':
			return
		}
	}
}

func (r *Samjung) readByte() (byte, error) {
	reader := bufio.NewReader(os.Stdin)
	b, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}

	return b, nil
}

func (r *Samjung) readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return text, nil
}
