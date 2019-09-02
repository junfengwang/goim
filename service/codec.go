package service

import (
	"bytes"
	"encoding/binary"
)

//2字节总长度 2字节版本 2字节操作 后面为

func Encode(msg *Message) ([]byte, error) {
	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.BigEndian, msg.datazise)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buffer, binary.BigEndian, msg.ve)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buffer, binary.BigEndian, msg.op)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buffer, binary.BigEndian, msg.data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func Decode(buf []byte) (*Message, error)  {
	bufReader := bytes.NewReader(buf)
	dataSize := len(buf)

	var ve int16
	err := binary.Read(bufReader, binary.BigEndian, &ve)
	if err != nil {
		return nil, err
	}

	var op int16
	err = binary.Read(bufReader, binary.BigEndian, &op)
	if err != nil {
		return nil, err
	}


	msgDataLen := dataSize - 4 -4
	msgBuf := make([]byte, msgDataLen)
	err = binary.Read(bufReader, binary.BigEndian, &msgBuf)
	if err != nil {
		return nil, err
	}

	message := &Message{
		ve:ve,
		op:op,
		data:msgBuf,
	}

	return message, nil
}
