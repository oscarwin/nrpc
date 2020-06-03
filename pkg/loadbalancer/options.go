package loadbalancer

import (
	"github.com/spaolacci/murmur3"
)

type Option func(*Options)

type Options struct {
	Key      string
	HashFunc func(string) uint32
}

func newOptions() *Options {
	return &Options{
		Key: "",
		HashFunc: func(key string) uint32 {
			return murmur3.Sum32([]byte(key))
		},
	}
}

func WithKey(key string) Option {
	return func(o *Options) {
		o.Key = key
	}
}

func WithHashFunc(hashFunc func(string) uint32) Option {
	return func(o *Options) {
		o.HashFunc = hashFunc
	}
}
