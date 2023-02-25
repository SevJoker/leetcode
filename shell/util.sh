


grep '/foundation/coupon/v1/couponinterface/getAvailableCoupons'  dirpc_0921_22_24.log | grep "_tie_to_sale_order_id" | grep 'coupon_type\\":\\"100' > dirpc_0921_22_24.log.step1
egrep 'total_count\\":([2-9]+|[0-9][0-9]+)' dirpc_0921_22_24.log.step1 > dirpc_0921_22_24.log.step2
cat  dirpc_0922_09_10.log.step2 | awk -F 'response=' '{print $2}' | jq -c | head


awk -v a="251" '{print index($0, a)}' dirpc_0922_09_10.log.step2


cat dirpc_0922_09_10.log.step2  | awk -v a="_tie_to_sale_order_id" '{pos = index($0, a); if (pos > 3) print pos;else print -1}' | head
cat  dirpc_0921_22_24.log.step2 | awk -F 'response=' '{print $2}' | awk -v a="_tie_to_sale_order_id" '{pos = index($0, a); if (pos < 10000) print pos}' > dirpc_0921_22_24.log.result1


cat /tmp/risk_voucher_filter_v5.log | awk -F '\\|\\|' '{print $2}' 


cat a | awk -F'\\|\\|' '{print $2}' | awk -F 'orderid=' '{print $1}' | awk -F '&' '{print $1}'



sort -n 1123 | uniq