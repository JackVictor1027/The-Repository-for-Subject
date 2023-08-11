package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	var err error
	router := gin.Default()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //Redis服务器地址
		Password: "",               //Redis最好设置密码，不能被任意的人访问
		DB:       0,                //Redis数据库索引，默认0
	})
	//先检查客户端和Redis的连接是否正常
	_, err = client.Ping().Result()
	if err != nil {
		log.Fatal(err)   //记录到日志中
		fmt.Println(err) //打印错误信息
	}

	//点赞功能/douyin/favorite/action/ - 赞操作（“改”）
	router.POST("/douyin/favorite/action", func(c *gin.Context) {
		var err error
		//需要用到client user_id video_id actionType
		//在Redis中将用户ID添加到视频点赞的有序集合中
		video_id := c.Request.Header.Get("VideoID")
		actionType := c.Request.Header.Get("ActionType")

		user_id := "123456" //待接收
		videoInfo := "bilibili"
		//videoInfo := model.VideoInfo{} //待接收

		jsonVideoInfo, err := json.Marshal(videoInfo) //序列化
		if err != nil {
			panic(err)
		}
		if actionType == "1" { //为视频点赞
			err = client.ZAdd(user_id, redis.Z{Score: 1, Member: jsonVideoInfo}).Err() //用户的点赞，要记录对于用户来说，ta点赞了多少个视频，有哪些视频被点赞了
			if err != nil {
				panic(err)
				return
			}
			_, err = client.ZAdd(video_id, redis.Z{Score: 1, Member: user_id}).Result() //视频的点赞，要记录对于视频来说，有多少人赞了，哪些人赞了
			if err != nil {
				panic(err)
				return
			}
			fmt.Println("您已成功为Id为" + video_id + "的视频点了一个赞！")
		} else { //给视频取消点赞
			err = client.ZRem(user_id, redis.Z{Score: 1, Member: jsonVideoInfo}).Err() //用户的点赞，要记录对于用户来说，ta点赞了多少个视频，有哪些视频被点赞了
			if err != nil {
				panic(err)
				return
			}
			_, err := client.ZRem(video_id, redis.Z{Score: 1, Member: user_id}).Result()
			if err != nil {
				panic(err)
				return
			}
			fmt.Println("您已成功给Id为" + video_id + "的视频取消点赞")
		}
		c.JSON(200, gin.H{"status": 0, "status_msg": "OK"}) //真的应该这样去处理错误数据吗
	})

	//获取Redis中用户ID为user_id的用户的所有点赞视频的有序集合/douyin/favorite/list/ - 喜欢列表（“查”）
	router.GET("/douyin/favorite/list/", func(c *gin.Context) {
		//var videolist []model.VideoInfo
		user_id := "123456" //待接收

		videos, _ := client.ZRange(user_id, 0, -1).Result()
		if err != nil {
			panic(err)
		}
		//err = json.Unmarshal([]byte(videos), &result)//只能一个个去解析了
		if err != nil {
			panic(err)
		}
		//type LikeList struct {
		//	StatusCode int               `json:"status_code"`
		//	StatusMsg  string            `json:"status_msg"`
		//	Videos     []model.VideoInfo `json:"videos"`
		//}
		type LikeList struct {
			StatusCode int      `json:"status_code"`
			StatusMsg  string   `json:"status_msg"`
			Videos     []string `json:"videos"`
		}
		likelist := LikeList{
			StatusCode: 200,
			StatusMsg:  "OK",
			Videos:     videos, //错误
		}
		c.JSON(200, likelist)
	})
	router.Run(":80")
}
