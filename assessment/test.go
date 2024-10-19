package assessment

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisClient struct {
	ctx context.Context
	rdb *redis.Client
}

func NewRedisClient(addr, password string, db, poolSize, minIdleConns int, poolTimeout time.Duration) *RedisClient {
	return &RedisClient{
		ctx: context.Background(),
		rdb: redis.NewClient(&redis.Options{
			Addr:         addr,
			Password:     password,
			DB:           db,
			PoolSize:     poolSize,
			MinIdleConns: minIdleConns,
			PoolTimeout:  poolTimeout,
		}),
	}
}

func (r *RedisClient) Close() {
	err := r.rdb.Close()
	if err != nil {
		log.Printf("failed to close redis: %v", err)
	}
}

func (r *RedisClient) StringExample(key string, value string) {
	// 设置字符串值
	err := r.rdb.Set(r.ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("SET 错误: %v", err)
	}

	// 获取字符串值
	val, err := r.rdb.Get(r.ctx, key).Result()
	if err != nil {
		log.Fatalf("GET 错误: %v", err)
	}
	fmt.Printf("GET %s: %s\n", key, val)
}

func (r *RedisClient) IncrementExample(key string) {
	// 自增
	val, err := r.rdb.Incr(r.ctx, key).Result()
	if err != nil {
		log.Fatalf("INCR 错误: %v", err)
	}
	fmt.Printf("INCR %s: %d\n", key, val)

	// 自减
	val, err = r.rdb.Decr(r.ctx, key).Result()
	if err != nil {
		log.Fatalf("DECR 错误: %v", err)
	}
	fmt.Printf("DECR %s: %d\n", key, val)
}

func (r *RedisClient) ListExample(key string, values ...string) {
	// 推入列表
	err := r.rdb.RPush(r.ctx, key, values).Err()
	if err != nil {
		log.Fatalf("RPUSH 错误: %v", err)
	}

	// 弹出列表第一个元素
	val, err := r.rdb.LPop(r.ctx, key).Result()
	if err != nil {
		log.Fatalf("LPOP 错误: %v", err)
	}
	fmt.Printf("LPOP %s: %s\n", key, val)
}

func (r *RedisClient) SetExample(key string, values ...string) {
	// 添加到集合
	err := r.rdb.SAdd(r.ctx, key, values).Err()
	if err != nil {
		log.Fatalf("SADD 错误: %v", err)
	}

	// 检查成员是否在集合中
	isMember, err := r.rdb.SIsMember(r.ctx, key, values[0]).Result()
	if err != nil {
		log.Fatalf("SISMEMBER 错误: %v", err)
	}
	fmt.Printf("SISMEMBER %s contains %s: %v\n", key, values[0], isMember)
}

func (r *RedisClient) SortedSetExample(key string) {
	// 添加元素到有序集合
	err := r.rdb.ZAdd(r.ctx, key, redis.Z{
		Score:  100.0,
		Member: "Alice",
	}).Err()
	if err != nil {
		log.Fatalf("ZADD 错误: %v", err)
	}

	// 获取有序集合的排名
	rank, err := r.rdb.ZRank(r.ctx, key, "Alice").Result()
	if err != nil {
		log.Fatalf("ZRANK 错误: %v", err)
	}
	fmt.Printf("ZRANK %s: Alice is ranked at %d\n", key, rank)
}

func (r *RedisClient) HashExample() {
	// 设置哈希字段
	err := r.rdb.HSet(r.ctx, "user:1000", "name", "Alice", "age", "30").Err()
	if err != nil {
		log.Fatalf("HSET 错误: %v", err)
	}

	// 获取哈希字段
	name, err := r.rdb.HGet(r.ctx, "user:1000", "name").Result()
	if err != nil {
		log.Fatalf("HGET 错误: %v", err)
	}
	fmt.Printf("user:1000's name: %s\n", name)

	age, err := r.rdb.HGet(r.ctx, "user:1000", "age").Result()
	if err != nil {
		log.Fatalf("HGET 错误: %v", err)
	}
	fmt.Printf("user:1000's age: %s\n", age)
}

func (r *RedisClient) HyperLogLogExample(key string, elements ...string) {
	// 添加元素到 HyperLogLog
	interfaceElements := make([]interface{}, len(elements))
	for i, v := range elements {
		interfaceElements[i] = v
	}

	err := r.rdb.PFAdd(r.ctx, key, interfaceElements...).Err()
	if err != nil {
		log.Fatalf("PFADD 错误: %v", err)
	}

	// 获取 HyperLogLog 的基数估计
	count, err := r.rdb.PFCount(r.ctx, key).Result()
	if err != nil {
		log.Fatalf("PFCOUNT 错误: %v", err)
	}
	fmt.Printf("PFCOUNT %s: %d unique elements\n", key, count)
}

func (r *RedisClient) BitmapExample(key string, offset int64, value int) {
	// 设置位图指定位置的值
	err := r.rdb.SetBit(r.ctx, key, offset, value).Err()
	if err != nil {
		log.Fatalf("SETBIT 错误: %v", err)
	}

	// 获取位图指定位置的值
	val, err := r.rdb.GetBit(r.ctx, key, offset).Result()
	if err != nil {
		log.Fatalf("GETBIT 错误: %v", err)
	}
	fmt.Printf("GETBIT %s at offset %d: %d\n", key, offset, val)
}

func (r *RedisClient) StreamExample(key string) {
	// 添加消息到 Stream
	err := r.rdb.XAdd(r.ctx, &redis.XAddArgs{
		Stream: key,
		Values: map[string]interface{}{"user": "Alice", "message": "Hello"},
	}).Err()
	if err != nil {
		log.Fatalf("XADD 错误: %v", err)
	}

	// 读取消息
	messages, err := r.rdb.XRead(r.ctx, &redis.XReadArgs{
		Streams: []string{key, "0"},
		Count:   1,
		Block:   0,
	}).Result()
	if err != nil {
		log.Fatalf("XREAD 错误: %v", err)
	}
	for _, msg := range messages {
		fmt.Printf("Stream %s: %v\n", key, msg.Messages)
	}
}

func (r *RedisClient) GeoExample(key string) {
	// 添加地理位置
	err := r.rdb.GeoAdd(r.ctx, key, &redis.GeoLocation{
		Longitude: 13.361389,
		Latitude:  38.115556,
		Name:      "Palermo",
	}).Err()
	if err != nil {
		log.Fatalf("GEOADD 错误: %v", err)
	}

	err = r.rdb.GeoAdd(r.ctx, key, &redis.GeoLocation{
		Longitude: 13.371389,
		Latitude:  38.125556,
		Name:      "William",
	}).Err()
	if err != nil {
		log.Fatalf("GEOADD 错误: %v", err)
	}

	// 获取地理位置的距离
	distance, err := r.rdb.GeoDist(r.ctx, key, "Palermo", "William", "km").Result()
	if err != nil {
		log.Fatalf("GEODIST 错误: %v", err)
	}
	fmt.Printf("GEODIST between Palermo and William: %f km\n", distance)

	locations, err := r.rdb.GeoSearch(r.ctx, key, &redis.GeoSearchQuery{
		Member:     "William",
		Radius:     100,
		RadiusUnit: "km",  // 单位可以是 m, km, mi, ft
		Sort:       "ASC", // 按距离排序
	}).Result()
	if err != nil {
		log.Fatalf("GEOSEARCH 错误: %v", err)
	}
	for _, location := range locations {
		fmt.Printf("Location found: %s\n", location)
	}

	rank, err := r.rdb.ZRank(r.ctx, key, "William").Result()
	if err != nil {
		log.Fatalf("ZRANK 错误: %v", err)
	}
	fmt.Printf("William is ranked at %d\n", rank)
}

func (r *RedisClient) PubSubExample(channel string) {
	// 订阅频道
	sub := r.rdb.Subscribe(r.ctx, channel)
	defer sub.Close()

	_, err := sub.Receive(r.ctx)
	if err != nil {
		panic(err)
	}

	// 发布消息到频道
	err = r.rdb.Publish(r.ctx, channel, "Hello, World!").Err()
	if err != nil {
		log.Fatalf("PUBLISH 错误: %v", err)
	}

	// 接收消息
	msg, err := sub.ReceiveMessage(r.ctx)
	if err != nil {
		log.Fatalf("ReceiveMessage 错误: %v", err)
	}
	fmt.Printf("Received message from %s: %s\n", msg.Channel, msg.Payload)
}

func RedisExample() {
	rds := NewRedisClient("localhost:6379", "", 0, 20, 5, 4*time.Second)
	defer rds.Close()

	rds.StringExample("example:string", "hello world")

	rds.ListExample("example:list", "item1", "item2", "item3")

	rds.SetExample("example:set", "member1", "member2", "member3")

	rds.HashExample()

	rds.IncrementExample("example:counter")

	rds.SortedSetExample("example:zset")

	rds.HyperLogLogExample("example:hll", "elem1", "elem2", "elem3")

	rds.BitmapExample("example:bitmap", 10, 1)

	rds.StreamExample("example:stream")

	rds.GeoExample("example:geo")

	rds.PubSubExample("example:channel")
}
