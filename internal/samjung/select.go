package samjung

import (
	"fmt"
	"os"
	"encoding/binary"
	"strconv"
)

type column struct {
	pk uint64
	name string
	position string
}

func (r *Samjung) selectRow() error {
	fmt.Print("pk: ")
	pkStr, err := r.readLine()
	if err != nil {
		return err
	}

	pkLen := len(pkStr) - 1
	pk, err := strconv.ParseUint(string(pkStr[:pkLen]), 10, 64)
	if err != nil {
		return err
	}
	
	err = r.selectOne(pk)
	if err != nil {
		return err
	}
	
	return nil
}

func (r *Samjung) selectOne(pk uint64) error {
	v, ok := r.indexMap[pk]
	if ok == false {
		fmt.Println("Not found matched row..")
		return nil
	}
	
	f, err := os.Open(r.baseDir+"/"+tableFile)
	if err != nil {
		return err
	}
	defer f.Close()
	
	col, err := r.readRow(f, int64(v), false)
	if err != nil {
		return err
	}
	
	fmt.Printf("ID : %v, Name : %v, Position : %v\n", col.pk, col.name, col.position)
	return nil
}

func (r *Samjung) readRow(f *os.File, offset int64, recursive bool) (column, error) {
	// pk 위치
	_, err := f.Seek(offset, 0)
	if err != nil {
		return column{}, err
	}
	
	pkBuf := make([]byte, 8)
	n, err := f.Read(pkBuf)
	if err != nil {
		return column{}, err
	}
	if n != 8 {
		return column{}, fmt.Errorf("Why not return 8 bytes?")
	}
	pk := binary.BigEndian.Uint64(pkBuf)
	
	// name의 variable integer
	offset = offset + int64(8)
	_, err = f.Seek(offset, 0)
	if err != nil {
		return column{}, err
	}
	
	nameLenBuf := make([]byte, 10)
	_, err = f.Read(nameLenBuf)
	if err != nil {
		return column{}, err
	}
	nameLen, varintLen := binary.Uvarint(nameLenBuf)
	
	// name value
	offset = offset + int64(varintLen)
	_, err = f.Seek(offset, 0)
	if err != nil {
		return column{}, err
	}
	
	nameBuf := make([]byte, nameLen)
	_, err = f.Read(nameBuf)
	if err != nil {
		return column{}, err
	}
	name := string(nameBuf)
	
	// position의 variable integer
	offset = offset + int64(nameLen)
	_, err = f.Seek(offset, 0)
	if err != nil {
		return column{}, err
	}
	
	positionLenBuf := make([]byte, 10)
	_, err = f.Read(positionLenBuf)
	if err != nil {
		return column{}, err
	}
	positionLen, varintLen := binary.Uvarint(positionLenBuf)
	
	// position value
	offset = offset + int64(varintLen)
	_, err = f.Seek(offset, 0)
	if err != nil {
		return column{}, err
	}
	
	positionBuf := make([]byte, positionLen)
	_, err = f.Read(positionBuf)
	if err != nil {
		return column{}, err
	}
	position := string(positionBuf)
	
	offset = offset + int64(positionLen)
	return column{pk, name, position}, nil
}
