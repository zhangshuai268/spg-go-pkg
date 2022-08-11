package iarray

import (
	logger "github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"os"
	"testing"
)

var arr1 = []int{1, 2, 3, 4, 5}
var arr2 = []int{2, 4, 6, 8, 10}
var arr3 = []int{2, 2, 6, 6, 10}

func TestMain(m *testing.M) {
	_, err := logger.InitLogger(false)
	if err != nil {
		return
	}
	os.Exit(m.Run())
}

func TestIntersect(t *testing.T) {
	intersect := Intersect(arr1, arr2)
	logger.Logger.Info(intersect)
}

func TestIsContain(t *testing.T) {
	contain, index := IsContain(arr1, 1)
	logger.Logger.Info(contain)
	logger.Logger.Info(index)
}
func TestReverse(t *testing.T) {
	Reverse(arr1)
	logger.Logger.Info(arr1)
}

func TestRemoveReplica(t *testing.T) {
	replica := RemoveReplica(arr3)
	logger.Logger.Info(replica)
}

func TestUnion(t *testing.T) {
	union := Union(arr1, arr2)
	logger.Logger.Info(union)
}
