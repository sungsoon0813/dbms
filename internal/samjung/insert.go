package samjung

import (
	"fmt"
)

// TODO : 픽스 시키지 않고 컬럼 이름으로 구분하여 컬럼 타입에 따라 insert할 수 있도록 변경 
func (r *Samjung) insertRow() {
	fmt.Print("name: ")
	name, err := r.readLine()
	if err != nil {
		fmt.Printf("failed to readLine: err=%v", err)
		return
	}
	
	fmt.Print("position: ")
	name, err := r.readLine()
	if err != nil {
		fmt.Printf("failed to readLine: err=%v", err)
		return
	}
	
	r.makePutArgs()
	r.writeToFile()	
}

// TODO : 현재는 uint64 및 string 두 가지만 받아들이며 추후 다른 타입들도 추가하도록 한다.
func (r *Samjung) makePutArgs(args ...interface{}) ([]byte, error) {
	if len(args) == 0 || args == nil {
		return nil, errors.New("Invalid parameter..")
	}

	buf = make([]byte, 0)
	for i, v := range args {
		if v == nil {
			continue
		}

		switch v := v.(type) {
		case *string:	// variable length + byte array
			// string length
			tmpBuf := make([]byte, 10)
			length = binary.PutUvarint(tmpBuf, len(v))
			buf = append(buf, tmpBuf[:length]...)
			
			// append string
			buf = append(buf, []byte(v))
		case *uint64: // uint64 to byte
			tmpBuf := make([]byte, 8)
			binary.BigEndian.PutUint32(tmpBuf, v)
			buf = append(buf, tmpBuf) 
		default:
			return nil, errors.New("Unsupported data type")
		}
	}

	return buf, nil
}