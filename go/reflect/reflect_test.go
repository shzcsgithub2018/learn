package reflect

import (
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	o := "{\"order\":{\"order_id\":\"8000008171558200578\",\"shop_number\":\"7\",\"is_book\":1,\"sys_expect_time\":\"1680184086-1680185286\",\"remark\":\"\",\"create_time\":1680181680,\"pay_time\":1680181686,\"is_self_deliver\":1,\"deliver_mode\":1,\"table_ware\":\"无需餐具\",\"is_deliver_later\":false},\"merchant\":{\"account_id\":\"7208440370433099813\",\"account_name\":\"深圳汇艺谷网络科技有限公司\"},\"receiver_info\":{\"receiver_name\":\"宋*生先生\",\"receiver_phone\":\"185****3852\",\"secret_number\":\"17882953655_9949\",\"province\":\"新疆维吾尔自治区\",\"city\":\"阿克苏地区\",\"district\":\"沙雅县\",\"town\":\"\",\"location_address\":\"庐山路华晨大拇指商业广场5楼\",\"lat\":40.101899,\"lng\":82.582757,\"location_name\":\"测试6\",\"door_plate_num\":\"1层101\"},\"amount\":{\"pay_amount\":20,\"origin_amount\":20,\"merchant_discount_amount\":0,\"platform_discount_amount\":0,\"product_origin_amount\":10,\"pay_discount_amount\":0,\"origin_deliver_fee\":10,\"pay_deliver_fee\":10},\"poi\":{\"poi_id\":\"7208401449636595749\",\"poi_name\":\"测试POI-saas外卖9813\"},\"products\":[{\"product_id\":\"1760422021173271\",\"sku_id\":\"1760422021173271\",\"product_name\":\"测试商品-勿购\",\"num\":1,\"commodities\":[{\"group_name\":\"测试商品\",\"total_count\":1,\"option_count\":1,\"items\":[{\"name\":\"测试商品\",\"count\":1,\"unit\":\"份\",\"price\":10}]}]}],\"amount_info\":{\"origin_amount\":20,\"pay_amount\":20,\"discount_amount\":0,\"discounts\":null},\"sub_order_amount_infos\":[{\"origin_amount\":10,\"pay_amount\":10,\"discount_amount\":0,\"discounts\":[{\"discount_type\":0,\"discount_amount\":0,\"merchant_discount_amount\":0,\"platform_discount_amount\":0}],\"sub_order_type\":100,\"sub_order_id\":\"800000817155820083413630578\"},{\"origin_amount\":10,\"pay_amount\":10,\"discount_amount\":0,\"discounts\":[{\"discount_type\":0,\"discount_amount\":0,\"merchant_discount_amount\":0,\"platform_discount_amount\":0}],\"sub_order_type\":1,\"sub_order_id\":\"800000817155820109015490578\"}]}"
	n := "{\"order\":{\"order_id\":\"8000008171558200578\",\"shop_number\":\"7\",\"is_book\":1,\"sys_expect_time\":\"1680184086-1680185286\",\"remark\":\"\",\"create_time\":1680181680,\"pay_time\":1680181686,\"is_self_deliver\":1,\"deliver_mode\":1,\"table_ware\":\"无需餐具\",\"is_deliver_later\":false},\"merchant\":{\"account_id\":\"7208440370433099813\",\"account_name\":\"深圳汇艺谷网络科技有限公司\"},\"receiver_info\":{\"receiver_name\":\"宋*生先生\",\"receiver_phone\":\"185****3852\",\"secret_number\":\"17882953655_9949\",\"province\":\"新疆维吾尔自治区\",\"city\":\"阿克苏地区\",\"district\":\"沙雅县\",\"town\":\"\",\"location_address\":\"庐山路华晨大拇指商业广场5楼\",\"lat\":40.101899,\"lng\":82.582757,\"location_name\":\"测试6\",\"door_plate_num\":\"1层101\"},\"amount\":{\"pay_amount\":20,\"origin_amount\":20,\"merchant_discount_amount\":0,\"platform_discount_amount\":0,\"product_origin_amount\":10,\"pay_discount_amount\":0,\"origin_deliver_fee\":10,\"pay_deliver_fee\":10},\"poi\":{\"poi_id\":\"7208401449636595749\",\"poi_name\":\"测试POI-saas外卖9813\"},\"products\":[{\"product_id\":\"1760422021173271\",\"sku_id\":\"1760422021173271\",\"product_name\":\"测试商品-勿购\",\"num\":1,\"commodities\":[{\"group_name\":\"测试商品\",\"total_count\":1,\"option_count\":1,\"items\":[{\"name\":\"测试商品\",\"count\":1,\"unit\":\"份\",\"price\":10}]}]}],\"amount_info\":{\"origin_amount\":20,\"pay_amount\":20,\"discount_amount\":0,\"discounts\":null},\"sub_order_amount_infos\":[{\"origin_amount\":10,\"pay_amount\":10,\"discount_amount\":0,\"discounts\":[{\"discount_type\":0,\"discount_amount\":0,\"merchant_discount_amount\":0,\"platform_discount_amount\":0}],\"sub_order_type\":100,\"sub_order_id\":\"800000817155820083413630578\"},{\"origin_amount\":10,\"pay_amount\":10,\"discount_amount\":0,\"discounts\":[{\"discount_type\":0,\"discount_amount\":0,\"merchant_discount_amount\":0,\"platform_discount_amount\":0}],\"sub_order_type\":1,\"sub_order_id\":\"800000817155820109015490578\"}]}"
	t.Log(reflect.DeepEqual(o, n))
}