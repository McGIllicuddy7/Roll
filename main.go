package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)
const UsageMessage = "usage roll xdy+z"
func main(){
	if(len(os.Args)<2){
		println(UsageMessage)
		return 
	}
	bytes := []byte(os.Args[1])
	for i := 2; i<len(os.Args); i++{
		bytes = append(bytes, []byte(os.Args[i])...)
	}
	amnt_bytes := make([]byte,0)
	die_bytes := make([]byte,0)
	add_bytes := make([]byte,0)
	add := true
	counter := 0
	l := len(bytes)
	got_d := false
	for counter <l{
		if bytes[counter] == 'd'{
			got_d = true
			counter ++
			break;
		}
		amnt_bytes = append(amnt_bytes, bytes[counter])
		counter ++;
	}
	if !got_d{
		println(UsageMessage)
		return
	}
	for counter <l{
		if bytes[counter] == '+'{
			add = true
			counter ++
			break
		}
		if bytes[counter] == '-'{
			add = false
			counter ++
			break
		}
		die_bytes = append(die_bytes, bytes[counter])
		counter ++
	}
	for counter <l{
		add_bytes = append(add_bytes, bytes[counter])
		counter ++;
	}
	amnt,err1 := strconv.ParseInt(string(amnt_bytes),10, 32)
	die,err2 := strconv.ParseInt(string(die_bytes),10,32)
	ad := int64(0)
	if(len(add_bytes)>=1){
		var err3 error;
		ad,err3 = strconv.ParseInt(string(add_bytes),10,32)
		if err3 != nil{
			println(UsageMessage)
			println("err3")
			return;
		}
	}
	if err1 != nil || err2 != nil{
		println(UsageMessage)
		println(string(amnt_bytes))
		println(string(die_bytes))
		println("err1 || err2")
		return
	}
	if int(amnt)<1 || int(die)<1{
		print(UsageMessage)
		return;
	}
	values := make([]int,0);
	total := int(ad)
	if !add {
		total *= -1
	}
	t := time.Now()
	ns := int64(t.Nanosecond())
	rand.Seed(ns)
	for i := 0; i<int(amnt); i++{
		v := int(rand.Int63()%die)+1
		values =append(values,v)
		total += v
	}
	fmt.Printf("Rolled: %d ", total)
	print("{")
	for i := 0; i<int(amnt-1); i++{
		print(values[i])
		print(", ")
	}
	print(values[amnt-1])
	if( ad != 0){
		print(", ")
		if(add){
			print("+")
		} else{
			print("-")
		}
		print(ad)
	}
	println("}")
}
