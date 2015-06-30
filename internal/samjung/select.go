package samjung

import (
	"fmt"
	"os"
	"encoding/binary"
)

type column struct {
	pk uint64
	name string
	position string
}

func (r *Samjung) selectRow() error {
	r.selectOne(9)
	return nil
}

func (r *Samjung) selectAll() {

}

func (r *Samjung) selectOne(pk uint64) {
	v, ok := r.indexMap[pk]
	if ok == false {
		fmt.Println("Not found matched row..")
		return
	}
	
	err := r.readRow(int64(v))
	if err != nil {
		fmt.Printf("%v", err)
	}
	// TODO : 파일 오픈
	// TODO : 파일 읽음어서 뿌려줌
}

func (r *Samjung) readRow(offset int64) error {
	f, err := os.Open(r.baseDir+"/"+tableFile)
	if err != nil {
		return err
	}
	defer f.Close()
	
	// pk 위치
	_, err = f.Seek(int64(offset), 0)
	if err != nil {
		return err
	}
	
	pkBuf := make([]byte, 8)
	n, err := f.Read(pkBuf)
	if err != nil {
		return err
	}
	if n != 8 {
		return fmt.Errorf("Why not return 8 bytes?")
	}
	pk := binary.BigEndian.Uint64(pkBuf)
	
	// name의 variable integer
	offset = offset + int64(8)
	_, err = f.Seek(offset, 0)
	if err != nil {
		return err
	}
	
	nameLenBuf := make([]byte, 10)
	_, err = f.Read(nameLenBuf)
	if err != nil {
		return err
	}
	nameLen, varintLen := binary.Uvarint(nameLenBuf)
	
	// name value
	offset = offset + int64(varintLen)
	_, err = f.Seek(offset, 0)
	if err != nil {
		return err
	}
	
	nameBuf := make([]byte, nameLen)
	_, err = f.Read(nameBuf)
	if err != nil {
		return err
	}
	name := string(nameBuf)
	
	
	fmt.Printf("pk: %v, name : %v", pk, name)
	
	return nil
}
