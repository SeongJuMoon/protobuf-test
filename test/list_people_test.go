package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/golang/protobuf/proto"

	pb "../api"
)

func TestWritePersonWritesPerson(t *testing.T) {
	buf := new(bytes.Buffer)

	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	// writePerson(buf, &p)
	got := buf.String()
	want := `Person ID: 1234
  Name: John Doe
  E-mail address: jdoe@example.com
  Home phone #: 555-4321
`

	if got != want {
		t.Errorf("writePerson(%s) => \n \t%q want %q", p.String(), got, want)
	}

}

func TestWritePersonWithProtoBuf(t *testing.T) {
	person := &pb.Person{
		Id:    1234,
		Name:  "SeongjuMoon",
		Email: "arcuer.dev@gmail.com",
	}

	fname := "output.proto"

	out, err := proto.Marshal(person)
	if err != nil {
		log.Fatalln("encoding person err", err)
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("failed file write person", err)
	}
}

func TestReadPersonWithProtoBuf(t *testing.T) {

	fname := "output.proto"

	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("failed to read file", err)
	}

	person := &pb.Person{}

	if err := proto.Unmarshal(in, person); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Printf("id : %d \n Name: %s, email-address: %s\n", person.Id, person.Name, person.Email)

}
