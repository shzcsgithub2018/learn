package set

import (
	"github.com/redis/go-redis/v9"
	"github.com/shzgithub2018/learn/redis/dal"
	"log"
)

type Set struct {
	key string
	cli *redis.Client
}

func NewSet(key string) *Set {
	dal.GetRedisCli(ctx).Del(ctx, key)
	return &Set{
		key: key,
		cli: dal.GetRedisCli(ctx),
	}
}

func (s *Set) SAdd(members ...string) {
	cmd := s.cli.SAdd(ctx, s.key, members)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	}
}

func (s Set) SRem(members ...string) {
	cmd := s.cli.SRem(ctx, s.key, members)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	}
}

func (s Set) SMove(key, element string) {
	cmd := s.cli.SMove(ctx, s.key, key, element)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	}
}

func (s Set) SMembers() {
	cmd := s.cli.SMembers(ctx, s.key)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (s Set) SCard() {
	cmd := s.cli.SCard(ctx, s.key)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (s Set) SIsMember(member string) {
	cmd := s.cli.SIsMember(ctx, s.key, member)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (s Set) SRandMember() {
	cmd := s.cli.SRandMember(ctx, s.key)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (s Set) SRandMemberN(count int64) {
	cmd := s.cli.SRandMemberN(ctx, s.key, count)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (s Set) SPop() {
	cmd := s.cli.SPop(ctx, s.key)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (s Set) SPopN(count int64) {
	cmd := s.cli.SPopN(ctx, s.key, count)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (s Set) SInter(keys ...string) {
	keys = append(keys, s.key)
	cmd := s.cli.SInter(ctx, keys...)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (s Set) SInterStore(des string, keys ...string) {
	keys = append(keys, s.key)
	cmd := s.cli.SInterStore(ctx, des, keys...)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(s.cli.SMembers(ctx, des).Val())
	}
}

func (s Set) SUnion(keys ...string) {
	keys = append(keys, s.key)
	cmd := s.cli.SUnion(ctx, keys...)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (s Set) SUnionStore(des string, keys ...string) {
	keys = append(keys, s.key)
	cmd := s.cli.SUnionStore(ctx, des, keys...)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(s.cli.SMembers(ctx, des).Val())
	}
}

func (s Set) SDiff(keys ...string) {
	keys = append([]string{s.key}, keys...)
	cmd := s.cli.SDiff(ctx, keys...)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(cmd.Val())
	}
}

func (s Set) SDiffStore(des string, keys ...string) {
	keys = append([]string{s.key}, keys...)
	cmd := s.cli.SDiffStore(ctx, des, keys...)
	if err := cmd.Err(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(s.cli.SMembers(ctx, des).Val())
	}
}
