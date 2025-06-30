package salesforceclient

import (
	"context"

	"github.com/sanservices/salesforce-client/grpc/salesforce-proto/pb"
	"google.golang.org/protobuf/types/known/structpb"
)

func (c *SfmcAgentClient) InsertDataRows(ctx context.Context, rows []interface{}, brand, extensionKey string) (*pb.InsertRowsResponse, error) {
	payload := map[string]interface{}{"items": rows}
	pbStruct, err := structpb.NewStruct(payload)
	if err != nil {
		return nil, err
	}

	in := &pb.InsertRowsRequest{
		Brand:        brand,
		Payload:      pbStruct,
		ExtensionKey: extensionKey,
	}
	res, err := c.Client.InsertRows(ctx, in)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *SfmcAgentClient) InsertDataRow(ctx context.Context, row interface{}, brand, extensionKey string) (*pb.InsertRowsResponse, error) {

	payload := map[string]interface{}{"items": row}
	pbStruct, err := structpb.NewStruct(payload)
	if err != nil {
		return nil, err
	}

	in := &pb.InsertRowsRequest{
		Brand:        brand,
		Payload:      pbStruct,
		ExtensionKey: extensionKey,
	}
	res, err := c.Client.InsertRows(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *SfmcAgentClient) CheckInsert(ctx context.Context, requestId, brand string) (*pb.CheckInsertRowsResponse, error) {
	in := &pb.CheckInsertRowsRequest{
		Brand:     brand,
		RequestId: requestId,
	}
	res, err := c.Client.CheckInsertRows(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}
