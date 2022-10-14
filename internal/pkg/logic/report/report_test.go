package report

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestGetReportDetail(t *testing.T) {
	//client, _ := elastic.NewClient(
	//	elastic.SetURL("http://172.17.101.191:9200"),
	//	elastic.SetSniff(false),
	//	elastic.SetBasicAuth("elastic", "ZSrfx4R6ICa3skGBpCdf"),
	//	elastic.SetErrorLog(log.New(os.Stdout, "APP", log.Lshortfile)),
	//	elastic.SetHealthcheckInterval(30*time.Second),
	//)
	//_, _, err := client.Ping("http://172.17.101.191:9200").Do(context.Background())
	//if err != nil {
	//	panic(fmt.Sprintf("es连接失败: %s", err))
	//}
	//var report rao.GetReportReq
	//report.ReportID = 762
	//err, result := GetReportDetail(context.Background(), report, "http://172.17.101.191:9200", "elastic", "ZSrfx4R6ICa3skGBpCdf")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//res, _ := json.Marshal(result)
	//log.Println(string(res))
	ctx := new(gin.Context)
	conf := fmt.Sprintf("mongodb://%s:%s@%s/%s", "kunpeng", "kYjJpU8BYvb4EJ9x", "172.17.18.255:27017", "kunpeng")

	clientOptions := options.Client().ApplyURI(conf)
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return
	}

	collection := mongoClient.Database("kunpeng").Collection("api_debug")

	filter := bson.D{{"uuid", "ecbbabcc-443c-46c4-abce-6170f19ca06b"}}
	fmt.Println("lllllll        ", collection)
	cur := collection.FindOne(context.TODO(), filter)
	result, err := cur.DecodeBytes()
	list, err := result.Elements()
	for index, value := range list {
		fmt.Println("index         ", index, " value:           ", string(value.Value().Value))
	}
	fmt.Println("1111111", result, " errr:           ", err)
	//if cur == nil {
	//	debug := bson.D{{fmt.Sprintf("%d", 123), "All"}}
	//	_, err = collection.InsertOne(ctx, debug)
	//	if err != nil {
	//		response.ErrorWithMsg(ctx, errno.ErrRedisFailed, err.Error())
	//		return
	//	}
	//} else {
	//	debug := bson.D{{fmt.Sprintf("%d", 123), "all"}}
	//	_, err = collection.UpdateMany(ctx, filter, debug)
	//	if err != nil {
	//		response.ErrorWithMsg(ctx, errno.ErrRedisFailed, err.Error())
	//		return
	//	}
	//}
}
