#### Unique ID Generators benchmarked

| Package                        | Generated ID sample                     |         op    | ns/op  |   B/op   |  allocs/op  |
|--------------------------------|-----------------------------------------|--------------:|-------:|---------:|------------:|
| <sub>github.com/rs/xid</sub>              | <sub>bk3nbif75t2rio46gqpg<br>bk3nbif75t2rio46gqq0<br>bk3nbif75t2rio46gqqg</sub>                    |	<sub>10000000</sub>   |   104  |      32  |          1  |
| <sub>github.com/segmentio/ksuid</sub>     | <sub>1MlrOfKT8cD6XzaBnohu9nUTBxE<br>1MlrOcSmL9MbMokyvQcpZNCZN5C<br>1MlrOgKNJQu3gM34KOs98VLdI3J</sub>             |	 <sub>1000000</sub>   |  1132  |      32  |          1  |
| <sub>github.com/kjk/betterguid</sub>      | <sub>-Lh_80ZiiWoATBE8exs2<br>-Lh_80ZiiWoATBE8exs3<br>-Lh_80ZiiWoATBE8exs4</sub>                    |	<sub>20000000</sub>   |   110  |      32  |          1  |
| <sub>github.com/oklog/ulid</sub>          | <sub>01DDJJ869E9KZS3NVMGQC8D6DN<br>01DDJJ869EZTGPDR68HCKAS0DW<br>01DDJJ869E9BJM8N9SBASPN7Q1</sub>              |	  <sub>200000</sub>   | 10859  |    5472  |          4  |
| <sub>github.com/sony/sonyflake</sub>      | <sub>385768684000066<br>385768684010066<br>385768684020066</sub>                         |	   <sub>20000</sub>   | 82055  |   32616  |        165  |
| <sub>github.com/chilts/sid</sub>          | <sub>1560769993008531236-6649545531703618877<br>1560769993008535930-4802730258495174943<br>1560769993008537813-4507423309930203142</sub> |	 <sub>5000000</sub>   |   330  |     115  |          3  |
| <sub>github.com/lithammer/shortuuid</sub> | <sub>BrjHNqCxbHkwBRLhkSRz78<br>HDHjMwEHjnoLYjodMvDALg<br>guPA33LzknFSQYNbKJtm4R</sub>                  |	  <sub>200000</sub>   |  7839  |    2953  |        136  |
| <sub>github.com/satori/go.uuid</sub>      | <sub>9ce723e4-5737-4e5b-a54a-e17086bf66d5<br>f4989905-508a-447e-9be4-3b44d5b14f1b<br>694b0805-23e6-4b70-89ee-b5735b59f90a</sub>    |	 <sub>2000000</sub>   |   886  |      64  |          2  |
| <sub>github.com/nu7hatch/gouuid</sub>     | <sub>721f0f92-13a1-4eb6-6442-a12984f0a131<br>58dcddc6-21f2-451e-4035-7098300b8398<br>b2223556-d9b4-4fe4-4134-41b55bb3a399</sub>    |	 <sub>1000000</sub>   |  1291  |     224  |          7  |
| <sub>github.com/google/uuid</sub>         | <sub>aa1fe2d6-7360-48ec-9c06-997f75664658<br>eea39318-5d61-450c-b867-e64f65fad0fa<br>d9d82f8a-7139-44a6-8702-b0c4c5424129</sub>    |	 <sub>2000000</sub>   |   884  |      64  |          2  |


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
