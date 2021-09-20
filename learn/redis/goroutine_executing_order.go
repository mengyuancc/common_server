package redis

const (
	REDISKEY = "mobile:p30:n3000:1"
)
/*func main()  {
	conn,err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis fail, the err is", err)
		return
	}
	defer conn.Close()

	item,err := redis.String(conn.Do("lpop", REDISKEY))
	if err != nil {
		fmt.Println("lpop to redis fail, the err is", err)
		return
	}
	fmt.Println("redis lpop success, the value is", item)
}
*/