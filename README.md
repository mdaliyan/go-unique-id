#### Unique ID Generators benchmarked

| Package                        | Generated ID sample                     |         op    | ns/op  |   B/op   |  allocs/op  |
|--------------------------------|-----------------------------------------|--------------:|-------:|---------:|------------:|
| github.com/rs/xid              | bjl5lbn75t2u0ivl8ffg                    |	10000000   |   104  |      32  |          1  |
| github.com/segmentio/ksuid     | 1LlV8MQHsewh2yTAzRTsGmALEza             |	 1000000   |  1132  |      32  |          1  |
| github.com/kjk/betterguid      | -LfnRYdIOFbNebdOJRMS                    |	20000000   |   110  |      32  |          1  |
| github.com/oklog/ulid          | 01DBSQ4EJKDZW3MWF9RZPSHRS9              |	  200000   | 10859  |    5472  |          4  |
| github.com/sony/sonyflake      | 37a17f03b000167                         |	   20000   | 82055  |   32616  |        165  |
| github.com/chilts/sid          | 1558862510675456033-0814957555925037708 |	 5000000   |   330  |     115  |          3  |
| github.com/lithammer/shortuuid | 2CsivtDgXFSTsCuKsnZTXT                  |	  200000   |  7839  |    2953  |        136  |
| github.com/satori/go.uuid      | ccee5b29-beb2-430d-b9dd-fb8caffb5c4a    |	 2000000   |   886  |      64  |          2  |
| github.com/nu7hatch/gouuid     | b027cb8b-62de-4ccd-697f-77778a532ebe    |	 1000000   |  1291  |     224  |          7  |
| github.com/google/uuid         | 3b00cc3b-07ab-412d-9266-1c9e4aefb9f4    |	 2000000   |   884  |      64  |          2  |

The above benchmark tests were ran on desktop computer

```
 *-cpu
      product: Intel(R) Core(TM) i3-4160 CPU @ 3.60GHz
      vendor: Intel Corp
      size: 3383MHz
      capacity: 3600MHz
      width: 64 bits
     
 *-memory
      size: 15GiB
```
