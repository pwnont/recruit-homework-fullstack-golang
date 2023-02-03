package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"sort"
	"sync"
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
	resultMap := make(map[string]bool)
	for _, name := range a {
		resultMap[name] = true
	}
	for _, name := range b {
		resultMap[name] = true
	}

	result = make([]string, 0, len(resultMap))
	for name := range resultMap {
		result = append(result, name)
	}

	sort.Strings(result)
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

	//user, _ := GetUserBasicInfo(id)
	//GetUserAddress(id)
	//return user, nil

	/*user, err := GetUserBasicInfo(id)
	if err != nil {
		return UserInfo{}, err
	}

	address, err := GetUserAddress(id)
	if err != nil {
		return UserInfo{}, err
	}

	user.Address = address
	return user, nil
	*/
	var user UserInfo
	var address UserAddress
	var err error

	var wg = sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		user, err = GetUserBasicInfo(id)
	}()

	go func() {
		defer wg.Done()
		address, err = GetUserAddress(id)
	}()

	wg.Wait()
	if err != nil {
		return UserInfo{}, err
	}

	user.Address = address
	return user, nil
}

/*
EX3. improve this function
*/
func Ex3FileOperation() error {
	/*OLD
	fmt.Println("Opening a file ex3.txt")
	file, _ := os.OpenFile("../assets/ex3.txt", os.O_RDWR|os.O_APPEND, 0644)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
	}

	file.WriteString("\r\nAppend this text to a file in Go")

	defer file.Close()
	return nil
	*/
	filePath := "../assets/ex3.txt"
	fmt.Println("Opening a file", filePath)

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", filePath, err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan file %s: %v", filePath, err)
	}

	text := "\r\nAppend this text to a file in Go"
	if _, err := file.WriteString(text); err != nil {
		return fmt.Errorf("failed to write to file %s: %v", filePath, err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("failed to close file %s: %v", filePath, err)
	}

	return nil
	/*improved
	1. if an error occurs, the function returns an error message indicating that the file could not be opened.
	2. if an error occurs, the function returns an error message indicating that the file could not be scanned.
	3. if an error occurs, the function returns an error message indicating that the file could not be written to.
	4. if an error occurs, the function returns an error message indicating that the file could not be closed.
	*/
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
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
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
