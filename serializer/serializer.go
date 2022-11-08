package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	b, err := protojson.MarshalOptions{
		Indent:          true,
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
	return string(b), err
}
