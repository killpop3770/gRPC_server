package main

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	//"google.golang.org/protobuf/encoding/protowire"
)
type GRPCServer struct{}

func (server *GRPCServer) A(ctx context.Context, cons *Consignment) (*Response, error){
	//F
	cons.F.InputVar = map[string]string{"1":"123","2":"123","3":"123","4":"123"}

	//S
	cons.S.X = "124"
	cons.S.Y = "123"
	cons.S.Z = "123"
	cons.S.D = "123"

	//func
	var count int
	v := reflect.ValueOf(cons.S)
	w := reflect.ValueOf(cons.F)
	fmt.Println(v.MapKeys())
	//fmt.Println(reflect.Indirect(v).FieldByName(cons.S.X))

	//Решение №0
	//protowire.ConsumeBytes()
	//protowire.ConsumeField(cons.S)

	//Решение №1
	//for i, j := range cons.F.InputVar{
	//
	//}

	//Решение №3
	//fmt.Println(strings.Contains(cons.F.InputVar[strconv.Itoa(1)], cons.S.Y))


	for _, e:= range reflect.Indirect(v).MapKeys(){
		ptrV := v.MapIndex(e)
		//println(reflect.Indirect(v).MapIndex(e).String())
		println(reflect.TypeOf(ptrV).String())
	}

	//for i:=0; i<reflect.Indirect(w).NumField(); i++{
	//	for _, j:= range reflect.Indirect(v).MapKeys(){
			//v1 := reflect.Indirect(v).Field(i).String()
			//w1 := reflect.Indirect(w).MapIndex(j).
			//t := reflect.Indirect(v).Interface()

			//println(t)
			//println(v1)
			//println(w1.String())

			//if strings.Contains(v1, w1) {
			//	count +=1
			//}

	//	}
	//}

	if count == 4{
		return &Response{Result: true}, nil
	}else{
		return &Response{Result: false}, nil
	}
}