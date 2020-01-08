OpenWRT administration HTTP api written in go  (Golang MIPS/ARM/x86 cross compilation)
----

This project was created to be used for a mobile APP to display system info and network status. OpenWRT configuartion changes and other features could also be implemented with little effort.

* Golang net/http and [httprouter](https://github.com/julienschmidt/httprouter) is used for HTTP API
* HTTP API uses JSON data to comunicate with ubus [(OpenWrt micro bus architecture)](https://openwrt.org/docs/techref/ubus)


Implemented API endpoints:

```
METHOD: GET

/
```
```
METHOD: GET

/system
```
```
METHOD: GET

/system/board
```
```
METHOD: GET

/system/services
```
```
METHOD: GET

/network
```
```
METHOD: GET

/network/interfaces
```
```
METHOD: GET

/network/interfaces/wan
```
```
METHOD: GET

/network/interfaces/lan
```
```
METHOD: GET

/wlan
```
```
METHOD: GET

/wlan/clients
```
