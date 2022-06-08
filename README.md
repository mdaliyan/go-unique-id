#### Unique ID Generators benchmarked
to run the benchmarks clone and:

```bash
go mod download
go test -bench=.
```
### results:

#### Machine:

```
  MacBook Pro (13-inch, 2019, Two Thunderbolt 3 ports)
  Processor: Intel(R) Core(TM) i7-8557U CPU @ 1.70GHz
  Memory: 16 GB 2133 MHz LPDDR3
```

#### For Go `v1.18`

| Package                                                         | Generated ID sample                                                                                                                      |                 op |  ns/op |  B/op | allocs/op |
|-----------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|-------------------:|-------:|------:|----------:|
| [rs/xid](https://github.com/rs/xid)                             | <sub>bk3nbif75t2rio46gqpg<br>bk3nbif75t2rio46gqq0<br>bk3nbif75t2rio46gqqg</sub>                                                          | <sub>9415598</sub> |  116.0 |    24 |         1 |
| [segmentio/ksuid](https://github.com/segmentio/ksuid)           | <sub>1MlrOfKT8cD6XzaBnohu9nUTBxE<br>1MlrOcSmL9MbMokyvQcpZNCZN5C<br>1MlrOgKNJQu3gM34KOs98VLdI3J</sub>                                     | <sub>1000000</sub> |   1105 |    32 |         1 |
| [kjk/betterguid](https://github.com/kjk/betterguid)             | <sub>-Lh_80ZiiWoATBE8exs2<br>-Lh_80ZiiWoATBE8exs3<br>-Lh_80ZiiWoATBE8exs4</sub>                                                          | <sub>9581594</sub> |  122.4 |    24 |         1 |
| [oklog/ulid](https://github.com/oklog/ulid)                     | <sub>01DDJJ869E9KZS3NVMGQC8D6DN<br>01DDJJ869EZTGPDR68HCKAS0DW<br>01DDJJ869E9BJM8N9SBASPN7Q1</sub>                                        |  <sub>131324</sub> |   8765 |  5472 |         4 |
| [sony/sonyflake](https://github.com/sony/sonyflake)             | <sub>385768684000066<br>385768684010066<br>385768684020066</sub>                                                                         |   <sub>30249</sub> |  39217 | 13760 |       180 |
| [chilts/sid](https://github.com/chilts/sid)                     | <sub>1560769993008531236-6649545531703618877<br>1560769993008535930-4802730258495174943<br>1560769993008537813-4507423309930203142</sub> | <sub>5079477</sub> |  234.3 |    98 |         3 |
| [lithammer/shortuuid](https://github.com/lithammer/shortuuid)   | <sub>BrjHNqCxbHkwBRLhkSRz78<br>HDHjMwEHjnoLYjodMvDALg<br>guPA33LzknFSQYNbKJtm4R</sub>                                                    |  <sub>179775</sub> |   6442 |  2873 |       135 |
| [satori/go](https://github.com/satori/go.uuid)                  | <sub>9ce723e4-5737-4e5b-a54a-e17086bf66d5<br>f4989905-508a-447e-9be4-3b44d5b14f1b<br>694b0805-23e6-4b70-89ee-b5735b59f90a</sub>          | <sub>1373052</sub> |  877.9 |    64 |         2 |
| [nu7hatch/gouuid](https://github.com/nu7hatch/gouuid)           | <sub>721f0f92-13a1-4eb6-6442-a12984f0a131<br>58dcddc6-21f2-451e-4035-7098300b8398<br>b2223556-d9b4-4fe4-4134-41b55bb3a399</sub>          |  <sub>894649</sub> |   1217 |   184 |         7 |
| [google/uuid](https://github.com/google/uuid)                   | <sub>aa1fe2d6-7360-48ec-9c06-997f75664658<br>eea39318-5d61-450c-b867-e64f65fad0fa<br>d9d82f8a-7139-44a6-8702-b0c4c5424129</sub>          | <sub>1341765</sub> |  853.9 |    64 |         2 |
| [lucsky/cuid](https://github.com/lucsky/cuid)                   | <sub>ckv5z1nw800004uaqlegq4tns<br>ckv5z25bx00005vaqyr1wqb4h<br>ckv5z2lsn00006waqwc0b6qcv</sub>                                           | <sub>3356278</sub> |  358.7 |    55 |         4 |
| [matoous/go-nanoid/v2](https://github.com/matoous/go-nanoid/v2) | <sub>4Oxn94IiH2RO10y7Rd6p4<br>BQKZ1qhr2ij52ACLGem_s<br>cfhM5t9ykJbu1_3xx8JP8</sub>                                                       | <sub>1000000</sub> |   1031 |   144 |         3 |

#### For Go `v1.17`

| Package                                                         | Generated ID sample                                                                                                                      |                   op |   ns/op |  B/op | allocs/op |
|-----------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|---------------------:|--------:|------:|----------:|
| [rs/xid](https://github.com/rs/xid)                             | <sub>bk3nbif75t2rio46gqpg<br>bk3nbif75t2rio46gqq0<br>bk3nbif75t2rio46gqqg</sub>                                                          |  <sub>10151221</sub> |   114.1 |    24 |         1 |
| [segmentio/ksuid](https://github.com/segmentio/ksuid)           | <sub>1MlrOfKT8cD6XzaBnohu9nUTBxE<br>1MlrOcSmL9MbMokyvQcpZNCZN5C<br>1MlrOgKNJQu3gM34KOs98VLdI3J</sub>                                     |    <sub>979638</sub> |    1103 |    32 |         1 |
| [kjk/betterguid](https://github.com/kjk/betterguid)             | <sub>-Lh_80ZiiWoATBE8exs2<br>-Lh_80ZiiWoATBE8exs3<br>-Lh_80ZiiWoATBE8exs4</sub>                                                          |   <sub>9532070</sub> |   122.1 |    24 |         1 |
| [oklog/ulid](https://github.com/oklog/ulid)                     | <sub>01DDJJ869E9KZS3NVMGQC8D6DN<br>01DDJJ869EZTGPDR68HCKAS0DW<br>01DDJJ869E9BJM8N9SBASPN7Q1</sub>                                        |    <sub>132578</sub> |    8529 |  5472 |         4 |
| [sony/sonyflake](https://github.com/sony/sonyflake)             | <sub>385768684000066<br>385768684010066<br>385768684020066</sub>                                                                         |     <sub>31164</sub> |   38191 | 13760 |       180 |
| [chilts/sid](https://github.com/chilts/sid)                     | <sub>1560769993008531236-6649545531703618877<br>1560769993008535930-4802730258495174943<br>1560769993008537813-4507423309930203142</sub> |   <sub>5181452</sub> |   231.6 |    98 |         3 |
| [lithammer/shortuuid](https://github.com/lithammer/shortuuid)   | <sub>BrjHNqCxbHkwBRLhkSRz78<br>HDHjMwEHjnoLYjodMvDALg<br>guPA33LzknFSQYNbKJtm4R</sub>                                                    |    <sub>179659</sub> |    6304 |  2873 |       135 |
| [satori/go](https://github.com/satori/go.uuid)                  | <sub>9ce723e4-5737-4e5b-a54a-e17086bf66d5<br>f4989905-508a-447e-9be4-3b44d5b14f1b<br>694b0805-23e6-4b70-89ee-b5735b59f90a</sub>          |   <sub>1303719</sub> |   890.5 |    64 |         2 |
| [nu7hatch/gouuid](https://github.com/nu7hatch/gouuid)           | <sub>721f0f92-13a1-4eb6-6442-a12984f0a131<br>58dcddc6-21f2-451e-4035-7098300b8398<br>b2223556-d9b4-4fe4-4134-41b55bb3a399</sub>          |    <sub>901984</sub> |    1197 |   184 |         7 |
| [google/uuid](https://github.com/google/uuid)                   | <sub>aa1fe2d6-7360-48ec-9c06-997f75664658<br>eea39318-5d61-450c-b867-e64f65fad0fa<br>d9d82f8a-7139-44a6-8702-b0c4c5424129</sub>          |   <sub>1250854</sub> |   945.8 |    64 |         2 |
| [lucsky/cuid](https://github.com/lucsky/cuid)                   | <sub>ckv5z1nw800004uaqlegq4tns<br>ckv5z25bx00005vaqyr1wqb4h<br>ckv5z2lsn00006waqwc0b6qcv</sub>                                           |   <sub>3237922</sub> |   370.5 |    55 |         4 |
| [matoous/go-nanoid/v2](https://github.com/matoous/go-nanoid/v2) | <sub>4Oxn94IiH2RO10y7Rd6p4<br>BQKZ1qhr2ij52ACLGem_s<br>cfhM5t9ykJbu1_3xx8JP8</sub>                                                       |   <sub>1000000</sub> |    1019 |   144 |         3 |
