# grpc-cleanarch-vue-go-example
BE: grpc + clearn architecture + Go
FE: grpc-web + Vuejs

## simple gps map (WIP)
display member's coordinate of same room
<img width="727" alt="スクリーンショット 2020-02-22 17 48 15" src="https://user-images.githubusercontent.com/21288308/75089409-d3cde780-559b-11ea-9bf1-f63eafa62f04.png">



### I/F
```
/radar/:roomname/:nickname
ex)/radar/roomA/ryoma
```

## Version
- Node:v13.2
- Go:v1.12

## USAGE
1. vue proto build
``` $ cd front && sh protobuild.sh && cd ..  ```

2. vue npm install
``` $ cd front && npm install && cd ..```

2. go proto build
``` cd server && make proto-build && cd ..```

3. go db setting
```$ make migrate/init && make migrate/up ```

4. running server
```$ cd server && make run ```

5. running vue-server
```$ npm run serve ```

display member point of same room

## TODO
- [ ] to set docker mysql

## LICENSE
MIT
