# go-p2pchat
## Usage 

### User1
set user1 user name and open port
```go run chat.go -name "user1" -p "1111" ```

### User2
set user2 and introduce user2 to user1
```go run chat.go -name "user2" -p "1112" -user "user1@(IP of user1):(Port of user1)" ```

### User3
set user3 and introduce user3 to user1 or user2
connecting other node automatically
```go run chat.go -name "user2" -p "1112" -user "user[1,2]@(IP of user[1,2]):(Port of user[1,2])" ```

## TODO
- [ ] public mode
- [ ] node list => all
- [ ] control

## PULL REQUEST
https://github.com/ryomak/go-p2pchat
