package service

import (
	"context"
	"time"

	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
	"helloworld/internal/pkg/jdl"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer
	jdLClient *jdl.JdLClient
	uc        *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, jdLCLient *jdl.JdLClient) *GreeterService {
	return &GreeterService{
		uc:        uc,
		jdLClient: jdLCLient,
	}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	// go func() {
	// 	time.Sleep(time.Second * 1)

	// 	preCheck := &jdl.PreCheckReq{
	// 		SenderContact: jdl.PlaceContact{
	// 			Name:        "张三",
	// 			Mobile:      "19200194150",
	// 			FullAddress: "北京市亦庄经济技术开发区科创十一街18号院京东大厦",
	// 		},
	// 		ReceiverContact: jdl.PlaceContact{
	// 			Name:        "张先生",
	// 			Mobile:      "19200194150",
	// 			FullAddress: "上海市浦东新区世纪大道1001号环球金融中心",
	// 		},
	// 		OrderOrigin:  1,
	// 		CustomerCode: "27K1234912",
	// 	}

	// 	preCheckResp, err := s.jdLClient.Precheck(context.Background(), preCheck)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	log.Info("preCheckResp: ", preCheckResp)
	// }()

	go func() {
		time.Sleep(time.Second * 1)

		// now := time.Now()

		placeOrder := &jdl.PlaceOrderReq{
			OrderID: "LG202204141438",
			SenderContact: jdl.PlaceContact{
				Name:        "张三",
				Mobile:      "17603474682",
				FullAddress: "北京大兴区亦庄经济开发区京东总部2号楼B座",
			},
			ReceiverContact: jdl.PlaceContact{
				Name:        "张先生",
				Mobile:      "19200194150",
				FullAddress: "北京大兴区亦庄经济开发区京东总部2号楼B座",
			},
			OrderOrigin:  1,
			CustomerCode: "27K1234912",
			ProductsReq: jdl.ProductInfo{
				ProductCode: "ed-m-0059",
			},
			SettleType: 3,
			Cargoes: []jdl.CargoInfo{
				{
					Name:     "服装",
					Quantity: 1,
					Weight:   100,
					Volume:   1,
				},
			},
			CommonChannelInfo: jdl.CommonChannelInfo{
				ChannelCode: "0030001",
			},
			// PickupStartTime: now.Add(time.Hour * 1).Unix(),
			// PickupEndTime:   now.Add(time.Hour * 3).Unix(),
			// ExpectDeliveryStartTime: time.Now().Add(time.Hour * 58).Unix(),
			// ExpectDeliveryEndTime:   time.Now().Add(time.Hour * 72).Unix(),
		}

		placeOrderResp, err := s.jdLClient.PlaceOrder(context.Background(), placeOrder)
		if err != nil {

		}

		log.Info("placeOrderResp: ", placeOrderResp)
	}()
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
