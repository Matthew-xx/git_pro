package persist

import (
	"../engine"
	"../model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T)  {
	expected := engine.Item{
		Url: "http://zhenai.com/u/105558",
		Type:"zhenai",
		Id:  "105558",
		Payload:model.Profile{
			Name: "小雪",
			Age: "24",
			Height: "165",
			Weight: "31",
			Marriage: "已婚",
			Income: "5-8千",
			Education: "本科",
		},
	}

	//依赖外部环境elastic search，先通过docker 运行elastic开启9200端口
	client,err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index  = "dating_profile"
	//save item
	err = Save(client,index,expected)

	if err != nil {
		panic(err)
	}



	//存一个东西再拿出来看是不是一样，一样就OK
	resp,err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())  //拿到id
	if err != nil {
		panic(err)
	}
	t.Logf("%s",resp.Source)

	var actual engine.Item
	err = json.Unmarshal([]byte(resp.Source),&actual)
	if err != nil {
		panic(err)
	}

	actualProfile,err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v",actual,expected)
	}
}
