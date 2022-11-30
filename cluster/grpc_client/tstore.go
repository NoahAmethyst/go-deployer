package grpc_client

import (
	"alliance-widget/constant"
	"alliance-widget/entity/entity_pb/tstore_pb"
	"alliance-widget/utils/log"
	"context"
	"errors"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"time"
)

var storeConn *grpc.ClientConn

func InitTStoreClient(address string) {
	if conn, err := StartConnection(address); err != nil {
		log.Error().Fields(map[string]interface{}{
			"action":  "init tstore client",
			"address": address,
			"error":   err,
		}).Send()
	} else {
		log.Info().Fields(map[string]interface{}{
			"action":  "init tstore client",
			"address": address,
		}).Send()
		storeConn = conn
	}
}

func GetTStoreClient() tstore_pb.TStoreServiceClient {
	return tstore_pb.NewTStoreServiceClient(storeConn)
}

func Save(v *tstore_pb.SaveParam) (*tstore_pb.SaveResp, error) {
	client := GetTStoreClient()
	// 设定请求超时时间 3s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return client.Save(ctx, v)
}

func Fetch(key string, path string) (*tstore_pb.FetchResp, error) {
	client := GetTStoreClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return client.Fetch(ctx, &tstore_pb.FetchParam{
		Uid:  key,
		Path: path,
	})
}

func Keys(key string, path string) (*tstore_pb.FetchResp, error) {
	client := GetTStoreClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return client.Keys(ctx, &tstore_pb.FetchParam{
		Uid:  "",
		Path: "",
	})
}

func Del(key string, path string) (*tstore_pb.FetchResp, error) {
	client := GetTStoreClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return client.Delete(ctx, &tstore_pb.FetchParam{
		Uid:  "",
		Path: "",
	})
}

func PBSave(uid string, path string, value proto.Message) error {

	_, err := Save(&tstore_pb.SaveParam{
		Uid:    uid,
		Path:   path,
		IValue: tstore_pb.NewIValue(value),
	})
	return err
}

func PBSaveString(uid string, path string, value string) error {
	_, err := Save(&tstore_pb.SaveParam{
		Uid:    uid,
		Path:   path,
		IValue: tstore_pb.NewIValue(value),
	})
	return err
}

func PBSaveInt(uid string, path string, value int32) error {

	_, err := Save(&tstore_pb.SaveParam{
		Uid:    uid,
		Path:   path,
		IValue: tstore_pb.NewIValue(value),
	})
	return err
}

func PBGetInt(uid string, path string) (int32, error) {

	v, err := Fetch(uid, path)
	if err != nil {
		log.Error().Msgf("fetch tstore error,%s", err.Error())
		return 0, err
	}
	if v.Code == 404 {
		return 0, nil
	}
	if v.IValue.Itype != tstore_pb.IValue_int || v.Code != constant.Success {
		log.Error().Msgf("fetch tstore error,%s", err)
		return 0, errors.New("pb get error")
	}
	return v.IValue.IntValue, nil
}

func PBGetString(uid string, path string) (string, error) {

	v, err := Fetch(uid, path)
	if err != nil {
		log.Error().Msgf("fetch tstore error,%s", err.Error())
		return "", err
	}
	if v.Code == 404 {
		return "", nil
	}
	if v.IValue.Itype != tstore_pb.IValue_str || v.Code != constant.Success {
		log.Error().Msgf("fetch tstore error,%s", err)
		return "", errors.New("pb get error")
	}
	return v.IValue.StrValue, nil
}

func PBGet(uid string, path string) ([]byte, error) {
	v, err := Fetch(uid, path)
	log.Info().Fields(map[string]interface{}{
		"pb get value": v,
	}).Send()
	if err != nil {
		log.Error().Msgf("fetch tstore error,%s", err.Error())
		return nil, err
	}
	if v.Code == 404 {
		return nil, nil
	}

	if v.IValue.Itype != tstore_pb.IValue_any || v.Code != constant.Success {
		log.Error().Msgf("fetch tstore error,%s", err)
		return []byte{}, errors.New("pb get error")
	}

	return v.IValue.AnyValue.Value, nil
}
