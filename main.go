package main

import (
	"fmt"
	"hackaton/dbms/internal/config"
	"hackaton/dbms/internal/samjung"
	"os"
)

// 필수기능
// – n개의컬럼으로구성된테이블구현(n>=2)
// – 입력 받은 데이터를 파일로 저장
// – 특정 데이터를 파일에서 검색해서 출력
// – 데이터 파일 인덱싱 (인덱스도 파일로 저장)
//
//추가 기능 (구현시 각각 가산점 부여!)
//– Delete & Update
//– Transaction Log (Redo & Undo)
//– Query cache
//– Remote client
//– Etc.

//만들어진 파일의 첫줄은 컬럼 이름 - 타입을 쌍으로 가지고 있도록 하고 메모리에 순서쌍으로 가지고 있음
//insert 하기 전에 파일 lock 걸고 파일 고치고 flush() 뒤에 lock 해제
//int는 4바이트 고정 , string은 variable length로 길이 체크하고 읽는방법으로 구현.
//검색을 할때는 무조건 primary key를 기준으로 그 row를 전체 다 가지고와서 순서쌍을 비교해서 보여주는 방식으로 함
//인덱스는 어떻게 구현할까 -> map 을 그대로 파일로 써놓고 다시 켜졌을때 다시 불러오는 방법으로 구현

func main() {

	// 설정 파일에서 data_dir을 읽어와서 작업
	c := config.New()
	if err := c.Read(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read configurations: %v\n", err)
		os.Exit(1)
	}

	// TODO : 이미 있는 테이블인지 새로 테이블을 생성할건지 선택하도록 구현

	// TODO : index 파일을 읽어들여 map으로 저장

	// samjung 이라는 테이블이 있는 상태로 구현되었다고 가정하고 진행한다.
	// ID (PK, Auto_increament) , name (string), position(string)

	sds := samjung.New(c.BaseDir)
	sds.Start()
}
