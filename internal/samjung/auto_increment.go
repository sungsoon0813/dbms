package samjung

import (
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

func (r *Samjung) getAutoIncrement() (uint64, error) {
	isExist := true
	if _, err := os.Stat(r.baseDir + "/" + autoIncFile); os.IsNotExist(err) {
		isExist = false
	}

	// autoIncFile 없으면 파일 생성하면서 1 집어넣음
	if isExist == false {
		f, err := os.OpenFile(r.baseDir+"/"+autoIncFile, os.O_WRONLY|os.O_CREATE, 0664)
		if err != nil {
			return 0, err
		}
		defer f.Close()

		_, err = f.WriteString("1")
		if err != nil {
			return 0, err
		}

		return 1, nil
	}

	// 있으면 + 1 해서 리턴해주고 저장
	f, err := os.OpenFile(r.baseDir+"/"+autoIncFile, os.O_WRONLY, 0664)
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

	numByte, err := ioutil.ReadFile(r.baseDir + "/" + autoIncFile)
	if err != nil {
		return 0, err
	}

	number, err := strconv.ParseUint(string(numByte), 10, 64)
	number = number + 1
	_, err = f.WriteString(strconv.FormatUint(number, 10))
	if err != nil {
		return 0, err
	}

	return number, nil
}

//func (r *Samjung) getAutoIncrement() (uint64, error) {
//	isExist := true
//	if _, err := os.Stat(r.baseDir + "/" + autoIncFile); os.IsNotExist(err) {
//		isExist = false
//	}
//
//	// 처음 생성되면 1로 셋팅
//	if isExist == false {
//		f, err := os.OpenFile(r.baseDir+"/"+autoIncFile, os.O_WRONLY|os.O_CREATE, 0664)
//		if err != nil {
//			return 0, err
//		}
//		defer f.Close()
//
//		buf := make([]byte, 8)
//		binary.BigEndian.PutUint64(buf, 1)
//		_, err = f.Write(buf)
//		if err != nil {
//			return 0, err
//		}
//
//		return 1, nil
//	}
//
//	// 이미 파일이 있으면 기존 값에서 +1
//	f, err := os.OpenFile(r.baseDir+"/"+autoIncFile, os.O_WRONLY, 0664)
//	if err != nil {
//		return 0, err
//	}
//	defer f.Close()
//
//	// file lock
//	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
//	if err != nil {
//		return 0, err
//	}
//	defer syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
//
//	var number uint64
//	err = binary.Read(f, binary.BigEndian, &number)
//	if err != nil {
//		return 0, err
//	}
//	fmt.Printf("number : %v", number)
//
//	buf := make([]byte, 8)
//	binary.BigEndian.PutUint64(buf, number+1)
//	_, err = f.Write(buf)
//	if err != nil {
//		return 0, err
//	}
//
//	return number, nil
//}
