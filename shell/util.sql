select
  sum("count") as y,
  sum("count") filter (
    where
      product_id in (1, 210, 100, 200, 220, 230, 235, 22, 20, 181, 13)
  ) as zhuankuaihao_product_count,
  sum("count") filter (
    where
      product_id in (500, 501)
  ) as huaxiaozhu_product_count,
  sum("count") filter (
    where
      caller = 'price-estimate'
  ) as price_estimate_count,
  Floor(__time to @__time) as __time
from
  "gocoupon_getcoupons"
where
  