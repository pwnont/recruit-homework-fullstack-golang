package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEx1UniqueNames(t *testing.T) {
	expectedResult := []string{"Ava", "Emma", "Olivia", "Sophia"}
	result := Ex1UniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"},
	)
	assert.Equal(t, expectedResult, result)
}

func TestEx2GetUserInfo(t *testing.T) {
	expectedUserInfo := UserInfo{
		ID:       1,
		Name:     "John Ginger",
		Company:  "JLC Group",
		Position: "Developer",
		Address: UserAddress{
			AddressNo:   "62 ซ.นาคนิวาส 6",
			SubDistrict: "ลาดพร้าว",
			District:    "ลาดพร้าว",
			Province:    "กทม",
			Postcode:    10230,
		},
	}

	start := time.Now()
	userResult, err := Ex2GetUserInfo(1)
	elapsed := time.Since(start)
	assert.LessOrEqual(t, elapsed, time.Duration(1*time.Second))

	assert.NoError(t, err)
	assert.Equal(t, expectedUserInfo, userResult)
}

func TestEx4MultiplesOf3and5(t *testing.T) {
	sum := Ex4MultiplesOf3and5(10)
	assert.Equal(t, 23, sum)
	sum = Ex4MultiplesOf3and5(49)
	assert.Equal(t, 543, sum)
	sum = Ex4MultiplesOf3and5(1000)
	assert.Equal(t, 233168, sum)
	sum = Ex4MultiplesOf3and5(8456)
	assert.Equal(t, 16687353, sum)
}

func BenchmarkEx4MultiplesOf3and5(b *testing.B) {
	// run function b.N times
	for n := 0; n < b.N; n++ {
		// Call your function
		Ex4MultiplesOf3and5(10000)
	}
}

// func BenchmarkF1(b *testing.B) {
// 	// run function b.N times
// 	for n := 0; n < b.N; n++ {
// 		// Call your function
// 	}
// }

func setup() {
}

func shutdown() {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
