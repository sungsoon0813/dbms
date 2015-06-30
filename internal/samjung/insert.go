package samjung

import (
	"encoding/binary"
	"errors"
	"os"
	"syscall"
	"fmt"
)

// TODO : 픽스 시키지 않고 컬럼 이름으로 구분하여 컬럼 타입에 따라 insert할 수 있도록 변경
func (r *Samjung) insertRow() error {
	fmt.Print("name: ")
	name, err := r.readLine()
	if err != nil {
		return err
	}

	fmt.Print("position: ")
	position, err := r.readLine()
	if err != nil {
		return err
	}
	
	pk, err := r.getAutoIncrement()
	if err != nil {
		return err
	}

	rowData, err := r.makePutArgs(pk, name, position)
	if err != nil {
		return err
	}

	offset, err := r.writeDataToFile(rowData)
	if err != nil {
		return err
	}
	
	err = r.makeIndex(pk, offset)
	if err != nil {
		return nil
	}

	return nil
}

// TODO : 현재는 uint64 및 string 두 가지만 받아들이며 추후 다른 타입들도 추가하도록 한다.
func (r *Samjung) makePutArgs(args ...interface{}) ([]byte, error) {
	if len(args) == 0 || args == nil {
		return nil, errors.New("Invalid parameter..")
	}

	buf := make([]byte, 0)
	// 컬럼 추가
	for _, v := range args {
		if v == nil {
			continue
		}

		switch v := v.(type) {
		case string: // variable length + byte array
			// string length
			tmpBuf := make([]byte, 10)
			l := len(v)
			length := binary.PutUvarint(tmpBuf, uint64(l))
			buf = append(buf, tmpBuf[:length]...)

			// append string
			buf = append(buf, []byte(v)...)
		case uint64: // uint64 to byte
			tmpBuf := make([]byte, 8)
			binary.BigEndian.PutUint64(tmpBuf, v)
			buf = append(buf, tmpBuf...)
		default:
			return nil, errors.New("Unsupported data type")
		}
	}

	return buf, nil
}

func (r *Samjung) writeDataToFile(rowData []byte) (uint64, error) {
	// open file
	f, err := os.OpenFile(r.baseDir+"/"+tableFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// file lock
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
	if err != nil {
		return 0, err
	}
	defer syscall.Flock(int(f.Fd()), syscall.LOCK_UN)

	// seek end offset of file
	offset, err := f.Seek(0, os.SEEK_END)
	if err != nil {
		return 0, err
	}

	// write to file
	_, err = f.WriteAt(rowData, offset)
	if err != nil {
		return 0, err
	}

	return uint64(offset), nil
}
