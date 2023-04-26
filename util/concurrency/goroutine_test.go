package concurrency

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	str := "{\n  \"Event\": \"DELIVER_WAIT_TAKING_ORDER\",\n  \"OrderId\": \"8000007649009477890\",\n  \"Content\": \"{\\\"OrderBase\\\":null,\\\"OrderDeliver\\\":{\\\"OrderId\\\":\\\"8000007649009477890\\\",\\\"OrderDeliverId\\\":\\\"800000802836771021014767890\\\",\\\"Address\\\":\\\"{\\\\\\\"UserAddressId\\\\\\\":7215109214165811459,\\\\\\\"LocationAddress\\\\\\\":{\\\\\\\"LocationId\\\\\\\":\\\\\\\"22535660540650921\\\\\\\",\\\\\\\"LocationAddress\\\\\\\":\\\\\\\"庐山路华晨大拇指商业广场5楼\\\\\\\",\\\\\\\"LocationName\\\\\\\":\\\\\\\"测试6\\\\\\\",\\\\\\\"ProvinceCode\\\\\\\":\\\\\\\"650000\\\\\\\",\\\\\\\"ProvinceName\\\\\\\":\\\\\\\"新疆维吾尔自治区\\\\\\\",\\\\\\\"CityCode\\\\\\\":\\\\\\\"652900\\\\\\\",\\\\\\\"CityName\\\\\\\":\\\\\\\"阿克苏地区\\\\\\\",\\\\\\\"DistrictCode\\\\\\\":\\\\\\\"652924\\\\\\\",\\\\\\\"DistrictName\\\\\\\":\\\\\\\"沙雅县\\\\\\\",\\\\\\\"TownCode\\\\\\\":\\\\\\\"\\\\\\\",\\\\\\\"TownName\\\\\\\":\\\\\\\"\\\\\\\",\\\\\\\"LodeLade\\\\\\\":{\\\\\\\"Lode\\\\\\\":82.582757,\\\\\\\"Lade\\\\\\\":40.101899}},\\\\\\\"DoorPlateNum\\\\\\\":\\\\\\\"1层101\\\\\\\",\\\\\\\"ConnectName\\\\\\\":\\\\\\\"宋先生\\\\\\\",\\\\\\\"Phone\\\\\\\":\\\\\\\"18595063852\\\\\\\",\\\\\\\"Gender\\\\\\\":1}\\\",\\\"IsBook\\\":1,\\\"ReceiverName\\\":\\\"宋先生\\\",\\\"ReceiverPhone\\\":\\\"@3jLSl08iN5xpuyjoxRwvcMgOaqomn9/+obSeG5MQt9w=\\\",\\\"IsVirtualPhone\\\":2,\\\"DeliverMode\\\":1,\\\"UserExpectTime\\\":\\\"\\\",\\\"ShopNumber\\\":\\\"9\\\",\\\"PoiId\\\":7208401449636595749,\\\"OrderShopId\\\":\\\"7208440370433099813\\\",\\\"OrderReceiveTimeOut\\\":600,\\\"SysExpectTime\\\":\\\"1679921908-1679923108\\\"},\\\"Status\\\":500,\\\"UpdateTime\\\":1679919508914}\",\n  \"UpdateTime\": 1679919508914\n}"
	s, _ := json.Marshal(str)
	t.Log(s)
}

func TestRace(t *testing.T) {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func TestMultiError(t *testing.T) {
	var (
		err error
		wg  sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Second)
		err = errors.New("safasfasf")
	}()
	wg.Wait()

	if err != nil {
		t.Error(err)
	}

}
