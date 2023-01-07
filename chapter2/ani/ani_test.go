// 测试代码
package main

import "testing"

func Test_Sleeping(t *testing.T) {
	c := NewCat("Sinmigo", "white", 20)
	c.Sleeping()
	c.Print()
}

func Test_Eating(t *testing.T) {
	c := NewCat("Sinmigo", "white", 20)
	c.Eating()

	if c.Color == "White" {
		t.Log("Eating测试通过")
	} else {
		t.Error("Eating测试不通过")
	}

}

func BenchmarkBigLen(b *testing.B) {
	//c :=NewCat("Sinmigo","white",20)
	for i := 0; i < b.N; i++ {
		//c.Sleeping()
	}
}
