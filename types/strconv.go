package types

import(
	"strconv"
)

func IntToString(i int)(string){
	
	return strconv.Itoa(i)
}

func  Int64ToString(i int64)(string){

	return strconv.FormatInt(i,10)
}

func StringToInt(s string)(int,error){

	return strconv.Atoi(s)
}

func StringToint64(s string)(int64,error){

	return  strconv.ParseInt(s, 10, 64)
}

func Hex2dec(hexstr string) string{

    i, _ := strconv.ParseInt(hexstr, 16, 0)
    return strconv.FormatInt(i, 10)
}
