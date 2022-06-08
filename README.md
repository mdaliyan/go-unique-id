#### Unique ID Generators benchmarked
to run the benchmarks clone and:

```bash
go mod download
go test -bench=.
```
### results:

| Package | Generated ID sample | op | ns/op | B/op | allocs/op |
|--------------------------------|-----------------------------------------|--------------:|-------:|---------:|------------:|
| [rs/xid](https://github.com/rs/xid) | <sub>bk3nbif75t2rio46gqpg<br>bk3nbif75t2rio46gqq0<br>bk3nbif75t2rio46gqqg</sub> | <sub>9213346</sub> | 123.9 | 24 | 1 |
| [segmentio/ksuid](https://github.com/segmentio/ksuid) | <sub>1MlrOfKT8cD6XzaBnohu9nUTBxE<br>1MlrOcSmL9MbMokyvQcpZNCZN5C<br>1MlrOgKNJQu3gM34KOs98VLdI3J</sub> | <sub>1214566</sub> | 981.3 | 32 | 1 |
| [kjk/betterguid](https://github.com/kjk/betterguid) | <sub>-Lh_80ZiiWoATBE8exs2<br>-Lh_80ZiiWoATBE8exs3<br>-Lh_80ZiiWoATBE8exs4</sub> | <sub>9499146</sub> | 126.2 | 24 | 1 |
| [oklog/ulid](https://github.com/oklog/ulid) | <sub>01DDJJ869E9KZS3NVMGQC8D6DN<br>01DDJJ869EZTGPDR68HCKAS0DW<br>01DDJJ869E9BJM8N9SBASPN7Q1</sub> | <sub>128148</sub> | 9215 | 5472 | 4 |
| [sony/sonyflake](https://github.com/sony/sonyflake) | <sub>385768684000066<br>385768684010066<br>385768684020066</sub> | <sub>23869</sub> | 50834 | 18024 | 217 |
| [chilts/sid](https://github.com/chilts/sid) | <sub>1560769993008531236-6649545531703618877<br>1560769993008535930-4802730258495174943<br>1560769993008537813-4507423309930203142</sub> | <sub>4780513</sub> | 250.1 | 98 | 3 |
| [lithammer/shortuuid](https://github.com/lithammer/shortuuid) | <sub>BrjHNqCxbHkwBRLhkSRz78<br>HDHjMwEHjnoLYjodMvDALg<br>guPA33LzknFSQYNbKJtm4R</sub> | <sub>157112</sub> | 6381 | 2873 | 135 |
| [satori/go](https://github.com/satori/go.uuid) | <sub>9ce723e4-5737-4e5b-a54a-e17086bf66d5<br>f4989905-508a-447e-9be4-3b44d5b14f1b<br>694b0805-23e6-4b70-89ee-b5735b59f90a</sub> | <sub>1605798</sub> | 747.6 | 64 | 2 |
| [nu7hatch/gouuid](https://github.com/nu7hatch/gouuid) | <sub>721f0f92-13a1-4eb6-6442-a12984f0a131<br>58dcddc6-21f2-451e-4035-7098300b8398<br>b2223556-d9b4-4fe4-4134-41b55bb3a399</sub> | <sub>1000000</sub> | 1118 | 184 | 7 |
| [google/uuid](https://github.com/google/uuid) | <sub>aa1fe2d6-7360-48ec-9c06-997f75664658<br>eea39318-5d61-450c-b867-e64f65fad0fa<br>d9d82f8a-7139-44a6-8702-b0c4c5424129</sub> | <sub>1577893</sub> | 759.0 | 64 | 2 |
| [lucsky/cuid](https://github.com/lucsky/cuid) | <sub>ckv5z1nw800004uaqlegq4tns<br>ckv5z25bx00005vaqyr1wqb4h<br>ckv5z2lsn00006waqwc0b6qcv</sub> | <sub>3193407</sub> | 398.3 | 55 | 4 |
| [matoous/go-nanoid/v2](https://github.com/matoous/go-nanoid/v2) | <sub>4Oxn94IiH2RO10y7Rd6p4<br>BQKZ1qhr2ij52ACLGem_s<br>cfhM5t9ykJbu1_3xx8JP8</sub> | <sub>1206334</sub> | 910.0 | 144 | 3 |

The above benchmark results are for Go 1.17 on:

```
  MacBook Pro (13-inch, 2019, Two Thunderbolt 3 ports)
  Processor: 1.7 GHz Quad-Core Intel Core i7
  Memory: 16 GB 2133 MHz LPDDR3
```
