### 프로토콜 버퍼란?
프로토콜 버퍼는 구글에서 개발하고 오픈소스로 공개한 직렬화 데이터 구조로서 다양한 언어(python js php ruby go c++)을 지원하며 
직렬화 속도가 빠르고 직렬화 된 파일의 크기가 작기도해서 Apache Avro 파일 포맷과 함께 이용한다.

이렇게 직렬화된 데이터를 전송하기 위해 gRPC로도 사용하는 모습을 볼 수 있다.

### Go 설치하기
https://golang.org/dl/ 에서 golang을 설치할 수 있습니다.


### 프로토콜 버퍼 컴파일러 설치하기
https://github.com/protocolbuffers/protobuf/releases/tag/v3.5.1 에서 운영체제에 맞는 컴파일러를 다운로드 받아주시면 됩니다.


### 프로토콜 버퍼 golang 표준 라이브러리 다운로드

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

### proto File 컴파일하기

현 저장소 클론 이후 .proto 파일의 컴파일 명령어는 아래와 같습니다.

```bash
protoc -I .\api\ --go_out=./api .\api\Person.proto
```

아래에 [protoname.pb.go]가 나오면 생성이 완료된겁니다!

### 더 많은 걸 알고 싶다면?
아래의 문서을 한번 읽어보세요!

(https://developers.google.com/protocol-buffers/docs/reference/proto3-spec)[protobuf specification]


### 트러블 슈팅
만약에 위 명령어를 쳐도 아무런 반응이 없다면 GOPATH가 환경변수에 잡혀있는지 확인해주세요.
