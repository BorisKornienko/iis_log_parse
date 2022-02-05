package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	for {
		log_writer("test.log")
		rand.Seed(time.Now().UnixNano())
		min_rand := 9
		max_rand := 999
		log_rand := rand.Intn(max_rand-min_rand+1) + min_rand
		time.Sleep(time.Duration(log_rand) * time.Millisecond)
		if log_rand > 400 {
			fmt.Println(log_rotator("test.log"))
		}
	}
}

func log_writer(file_name string) {
	// this func write random lines by tamplate to log file
	micro_time := time.Now() //micro_time.String()
	log_body := [7]string{"With", "way", "package", "excludes", "slice", "dynamically", "indices"}
	log_code := [7]string{"200", "206", "400", "404", "403", "407", "500"}

	rand.Seed(time.Now().UnixNano())
	min_rand := 9
	max_rand := 999
	log_rand := rand.Intn(max_rand-min_rand+1) + min_rand

	min_slice := 0
	max_slice := 6
	slice_rand := rand.Intn(max_slice-min_slice+1) + min_slice
	slice_rand_2 := rand.Intn(max_slice-min_slice+1) + min_slice
	slice_rand_3 := rand.Intn(max_slice-min_slice+1) + min_slice

	log_string := micro_time.String() + "   " + log_body[slice_rand] + " " + log_body[slice_rand_2] + " " + log_body[slice_rand_3] + "   " + strconv.Itoa(log_rand) + "   " + log_code[slice_rand] + "\n"
	fmt.Println(log_string)

	// os.WriteFile(file_name, []byte(log_string), 0666)
	f, err := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(log_string); err != nil {
		log.Println(err)
	}
}

func log_rotator(file_name string) int {
	// this func check file size if I forget to kill process
	info, err := os.Stat(file_name)
	if err != nil {
		log.Println(err)
	}
	if info.Size() > 1000000 {
		os.WriteFile(file_name, []byte(" "), 0666)
	}
	return int(info.Size())
}
