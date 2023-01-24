package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

/*
EX1. Edit this function to pass Unit test

When passed two slices of names,
it will return a slice containing the names that appear in either or both slices.
The returned slice should have no duplicates.
*/
func Ex1UniqueNames(a, b []string) []string {
	var result []string

	//
	// YOUR CODE HERE
	//

	return result
}

/*
EX2. Edit this function to pass Unit test
*/
func Ex2GetUserInfo(id int) (UserInfo, error) {

	// ฟังก์ชั่นนี้สามารถแก้โค๊ดได้ทั้งหมด

	//
	// YOUR CODE HERE
	//

	user, _ := GetUserBasicInfo(id)
	GetUserAddress(id)
	return user, nil
}

/*
EX3. improve this function
*/
func Ex3FileOperation() error {
	fmt.Println("Opening a file ex3.txt")
	file, _ := os.OpenFile("../assets/ex3.txt", os.O_RDWR|os.O_APPEND, 0644)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
	}

	file.WriteString("\r\nAppend this text to a file in Go")

	defer file.Close()
	return nil
}

/*
EX4. เขียนฟังก์ชั่นนี้ให้ทำงานเร็วที่สุด และ ผ่าน Unit test

หาผลรวมของจำนวนที่เป็นเท่าของ (multiples of) 3 หรือ 5 ที่น้อยกว่าจำนวนที่กำหนดให้
ตัวอย่างเช่น จำนวนที่กำหนดให้คือ 10
จะได้จำนวนที่ที่เป็นเท่าของ 3 หรือ 5 คือ 3, 5, 6 และ 9. ผลรวมคือ 23
*/
func Ex4MultiplesOf3and5(number int) int {
	sum := 0

	//
	// YOUR CODE HERE
	//

	return sum
}

type UserInfo struct {
	ID       int
	Name     string
	Company  string
	Position string
	Address  UserAddress
}
type UserAddress struct {
	AddressNo   string
	SubDistrict string
	District    string
	Province    string
	Postcode    int
}

func GetUserBasicInfo(id int) (UserInfo, error) {
	// Simulate database slow query
	time.Sleep(500 * time.Millisecond)
	user := UserInfo{
		ID:       id,
		Name:     "John Ginger",
		Company:  "JLC Group",
		Position: "Developer",
	}
	return user, nil
}
func GetUserAddress(id int) (UserAddress, error) {
	// Simulate database slow query
	time.Sleep(800 * time.Millisecond)
	userAddress := UserAddress{
		AddressNo:   "62 ซ.นาคนิวาส 6",
		SubDistrict: "ลาดพร้าว",
		District:    "ลาดพร้าว",
		Province:    "กทม",
		Postcode:    10230,
	}
	return userAddress, nil
}

func main() {

}
