package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Features struct {
	// omitempty 在序列化时,忽略空值
	Id    primitive.ObjectID 		`json:"_id,omitempty" bson:"_id,omitempty"`
	Value string					`json:"value,omitempty" bson:"value,omitempty" `
	Title string					`json:"title,omitempty" bson:"title,omitempty"`
	//Num   int						`json:"num,omitempty"`
	//// 案件来源
	//Analysis_source string 			`json:"analysis_source,omitempty"`
	//// 案件编号
	//Case_id string					`json:"case_id,omitempty"`
	//// 相关案件编号
	//Analysis_ids []string			`json:"analysis_ids,omitempty"`
	//// 分类
	//Category      string   			`json:"category,omitempty"`
	//Category_code string			`json:"category_code,omitempty"`
	//Part_category string			`json:"part_category,omitempty"`
	//// 框架名称
	//Framework string				`json:"framework,omitempty"`
	//Level     string				`json:"level,omitempty"`
	//Status    int					`json:"status,omitempty"`
	//Push      int					`json:"push,omitempty"`
	//// 行为
	//Behavior string					`json:"behavior,omitempty"`
	//// 格式
	//Format string					`json:"format,omitempty"`
	//
	//Is_regex bool					`json:"is_regex,omitempty"`
	//// 分析师，实时修改
	//Creator string					`json:"creator,omitempty"`
	//// 上线操作, 审核人
	//Auditor string					`json:"auditor,omitempty"`
	//// 审核时间，包括审核成功审核失败
	//Audited_at time.Time			`json:"audited_at,omitempty"`
	//// 布控时间
	//Deployed_at time.Time			`json:"deployed_at,omitempty"`
	//// 待回扫时间
	//Rescan_at time.Time				`json:"rescan_at,omitempty"`
	//// 验证时间
	//Check_at time.Time				`json:"check_at,omitempty"`
	//// 停用时间
	//Stop_at time.Time				`json:"stop_at,omitempty"`
	//// 登录状态
	//Login_status string				`json:"login_status,omitempty"`
	//// 备注
	//Remarks string					`json:"remarks,omitempty"`
	//// 特征审核备注
	//Audited_remark string			`json:"audited_remark,omitempty"`
	//Is_deleted     bool				`json:"is_deleted,omitempty"`
	//// 末次命中时间
	//Last_hit_at time.Time			`json:"last_hit_at,omitempty"`
	//// 中标累计次数
	//Hits_num int					`json:"hits_num,omitempty"`

	C_at *time.Time		`json:"c_at,omitempty" bson:"c_at,omitempty"`
	U_at *time.Time		`json:"u_at,omitempty" bson:"u_at,omitempty"`
}
//type Post struct {
//	Value string `json:"value,omitempty"`
//
//	num int `json:"int,omitempty"`
//}

func main() {
	// 设置客户端链接配置
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.110.45:27017"))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//err = client.Connect(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// 设置客户端链接配置
	clientOptions := options.Client().ApplyURI("mongodb://192.168.110.45:27017")

	// 链接mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 检查链接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	// 指定获取要操作的数据集
	db := client.Database("liu")
	col_feature := db.Collection("features")
	//col_feature := db.Collection("feature")
	//var feature Features
	//filtter := bson.M{"value": "/images/dy/cp_logo.png"}
	//err = col_feature.FindOne(context.TODO(), filtter).Decode(&feature)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(feature)
	//fmt.Printf("%+v\n", *feature)
	//fmt.Printf("Found a single document: %+v\n", feature)

	//post := Post{"title", "body"}
	//
	//
	//
	//insertResult, err := col_feature.InsertOne(context.TODO(), post)
	//
	//
	//
	//if err != nil {
	//
	//	log.Fatal(err)
	//
	//}
	//
	//
	//
	//fmt.Println("Inserted post with ID:", insertResult.InsertedID)
	filter := bson.D{}

	var post Features

	err = col_feature.FindOne(context.TODO(), filter).Decode(&post)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("Found post with  ", post.Id)
	Id := post.Id.Hex()
	fmt.Printf("%T\n", Id)
	fmt.Printf("%v\n", Id)
	fmt.Printf("Found a single document: %+v\n", post)
}
