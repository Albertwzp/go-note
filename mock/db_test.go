package main

import (
    "testing"
    "errors"
    "github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    m := NewMockDB(ctrl)
    m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exsist"))

    if v := getFromDB(m, "Tom"); v != -1 {
        t.Fatal("expect -1, but got ", v)
    }
}
