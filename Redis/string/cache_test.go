package String

import (
	"github.com/shzgithub2018/learn/Redis/dal"
	"testing"
)

func TestCache(t *testing.T) {
	key, val := "《面朝大海，春暖花开》", "\n从明天起，做一个幸福的人\n喂马、劈柴，周游世界\n从明天起，关心粮食和蔬菜\n我有一所房子，面朝大海，春暖花开\n从明天起，和每一个亲人通信\n告诉他们我的幸福\n那幸福的闪电告诉我的\n我将告诉每一个人\n给每一条河每一座山取一个温暖的名字\n陌生人，我也为你祝福\n愿你有一个灿烂的前程\n愿你有情人终成眷属\n愿你在尘世获得幸福\n我只愿面朝大海，春暖花开"
	cache := NewCache(ctx, dal.GetRedisCli(ctx))
	cache.Set(key, val)
	t.Log(key, cache.Get(key))
}
