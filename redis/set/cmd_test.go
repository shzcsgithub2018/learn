package set

import (
	"github.com/shzgithub2018/learn/redis/dal"
	"log"
	"testing"
)

var (
	key1 = "database"
	key2 = "databaseV2"

	members1 = []string{"a", "b", "c", "d"}
	members2 = []string{"1", "2", "3"}

	k1, k2 = "set1", "set2"
	set1   = []string{"1", "2", "3", "4"}
	set2   = []string{"1", "3", "5", "7"}
)

func TestSMove(t *testing.T) {
	set := NewSet(key1)
	set.SMove(key2, "a")
	set.SMembers()

	set.SAdd(members1...)

}

func SMembers(key string) {
	log.Println(dal.GetRedisCli(ctx).SMembers(ctx, key).Val())
}

func TestCollectionOp(t *testing.T) {
	s1, s2 := NewSet(k1), NewSet(k2)
	s1.SAdd(set1...)
	s2.SAdd(set2...)

	SMembers(k1)
	SMembers(k2)
	/*Output
	[1 2 3 4]
	[1 3 5 7]
	*/

	t.Log("Inter")
	s1.SInter(k2) // [1 3]

	t.Log("Union")
	s1.SUnion(k2) // [1 2 3 4 5 7]

	t.Log("Diff")
	s1.SDiff(k2) //[2 4]
}
