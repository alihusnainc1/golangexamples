package golangexamples
import "github.com/ehteshamz/greetings"
func ConcatSlice(sliceToConcat []byte) string{
	temp := ""
	for i:=0 ; i<len(sliceToConcat);i++ {
		temp+=string(sliceToConcat[i])
		if(i!=len(sliceToConcat)-1){
		temp+="-"}
	}
	return temp
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int){
	for i:=0 ; i<len(sliceToEncrypt);i++ {
		sliceToEncrypt[i]=sliceToEncrypt[i]+byte(ceaserCount)
	}
}

func EZGreetings(name string) string{
	temp:=""
	temp=greetings.PrintGreetings(name)
	return temp
	
	}
