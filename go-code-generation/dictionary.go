package main

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "KeyType=string,int ValueType=string,int"

type KeyType generic.Type
type ValueType generic.Type

type KeyTypeValueTypeDictionary struct {
	items map[KeyType]ValueType
}
