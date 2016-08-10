# enocean
Tools for EnOcean protocol 



## References

* ERP: EnOcean Radio Protol : https://www.**enocean**.com/erp2/
* ESP: EnOcean Serial Protocol : https://www.enocean.com/esp
* EEP: EnOcean Equipment Prifile : http://www.enocean-alliance.org/eep/

Copyright 2015 Kazutaka Ogaki <ocaokgbu@gmail.com>

Licensed under the MIT License


## Samples

### mqttworker.go

```
go run mqttworker.go loop --host broker.hostname --sub 'prefix/enoceangateway/topic' --pub 'prefix/worker/enocean/publish'
```

sample outputs

```
INFO[0000] Broker URI: tcp://broker.hostname          
INFO[0000] Sub Topic: prefix/enoceangateway/topic                     
INFO[0000] Pub Topic: prefix/worker/enocean/publish     
INFO[0000] connecting...                                
INFO[0000] client connected                             
INFO[0038] topic:prefix/enoceangateway/topic msg:U

?"OB?=? 
{"RawData": "{165 0 0 0 false [4 0 79 66] [] [0 0 50 8] 61}", "TeachIn": false, "manufactuererId": "MANUFACTURER_RESERVED", "temperature": 32}
INFO[0038] published: %v, to topic: %v{"RawData": "{165 0 0 0 false [4 0 79 66] [] [0 0 50 8] 61}", "TeachIn": false, "manufactuererId": "MANUFACTURER_RESERVED", "temperature": 32}prefix/worker/enocean/publish 
^Csignal: interrupt
```

subscriber output

```
$ mosqutto_sub -t "#" -d
Client mosqsub/hostname received PUBLISH (d0, q0, r0, m0, 'prefix/worker/enocean/publish', ... (142 bytes))

{"RawData": "{165 0 0 0 false [4 0 79 66] [] [0 0 50 8] 61}", "TeachIn": false, "manufactuererId": "MANUFACTURER_RESERVED", "temperature": 32}
```
