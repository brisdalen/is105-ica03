package frequence

import (
	"testing"
)

func BenchmarkFrequencepg100(b *testing.B) {
	benchFrequencePG100(1,b)

}
func BenchmarkFrequenceBible(b *testing.B) {
	benchFrequenceBible(1,b)
}

func BenchmarkFrequenceBible2(b *testing.B) {
	benchBible2(1,b)
}

func benchFrequencePG100(i int, b *testing.B) {
	for n := 0; n < b.N; n++{
		b.StopTimer()
		b.StartTimer()
		Frequence("../files/pg100.txt")
	}
}
func benchFrequenceBible(i int, b*testing.B){
	for n := 0; n < b.N; n++{
		b.StopTimer()
		b.StartTimer()
		Frequence("../files/bible.txt")
	}
}
func benchmarkFrequenceBible2(i int,b *testing.B) {
	for n := 0; n < b.N; n++{
		b.StopTimer()
		b.StartTimer()
		Frequence("../files/bible2.txt")
	}

}