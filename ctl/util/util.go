package util

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// LoadSolutionsDir define
func LoadSolutionsDir() ([]int, []string, int) {
	solutionIds, soNames, total := loadFile("../leetcode/")
	fmt.Printf("Read the solutions of %v questions, there are %v files (may include .DS_Store) in the current directory, and %v questions in the directory are trying\n", len(solutionIds), total, total-len(solutionIds))
	return solutionIds, soNames, total - len(solutionIds)
}

func loadFile(path string) ([]int, []string, int) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	solutionIds, soNames, solutionsMap := []int{}, []string{}, map[int]string{}
	for _, f := range files {
		if f.Name()[4] == '.' {
			tmp, err := strconv.Atoi(f.Name()[:4])
			if err != nil {
				fmt.Println(err)
			}
			solutionIds = append(solutionIds, tmp)
			solutionsMap[tmp] = f.Name()
		}
	}
	sort.Ints(solutionIds)
	for _, v := range solutionIds {
		if name, ok := solutionsMap[v]; ok {
			soNames = append(soNames, name)
		}
	}
	return solutionIds, soNames, len(files)
}

// GetAllFile define
func GetAllFile(pathname string, fileList *[]string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			// fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname+fi.Name()+"/", fileList)
		} else {
			// fmt.Println(fi.Name())
			*fileList = append(*fileList, fi.Name())
		}
	}
	return *fileList, err
}

// LoadChapterFourDir define
func LoadChapterFourDir() ([]string, []int) {
	files, err := GetAllFile("../website/content/ChapterFour/", &[]string{})
	// files, err := ioutil.ReadDir("../website/content/ChapterFour/")
	if err != nil {
		fmt.Println(err)
	}
	solutions, solutionIds, solutionsMap := []string{}, []int{}, map[int]string{}
	for _, f := range files {
		if f[4] == '.' {
			tmp, err := strconv.Atoi(f[:4])
			if err != nil {
				fmt.Println(err)
			}
			solutionIds = append(solutionIds, tmp)
			// len(f.Name())-3 = file name without .md suffix
			solutionsMap[tmp] = f[:len(f)-3]
		}
	}
	sort.Ints(solutionIds)
	fmt.Printf("Read the solutions of %v questions in Chapter 4\n", len(solutionIds))
	for _, v := range solutionIds {
		if name, ok := solutionsMap[v]; ok {
			solutions = append(solutions, name)
		}
	}
	return solutions, solutionIds
}

// WriteFile define
func WriteFile(fileName string, content []byte) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("write file successful")
}

// LoadFile define
func LoadFile(filePath string) ([]byte, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader, output := bufio.NewReader(f), []byte{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, nil
			}
			return nil, err
		}
		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}
}

// DestoryDir define
func DestoryDir(path string) {
	filepath.Walk(path, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if !fi.IsDir() {
			return nil
		}
		name := fi.Name()
		if strings.Contains(name, "temp") {
			fmt.Println("temp file name:", path)
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Println("delet dir error:", err)
			}
		}
		return nil
	})
}

// CopyFile define
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// BinarySearch define
func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// GetChpaterFourFileNum define
func GetChpaterFourFileNum(num int) string {
	if num < 100 {
		return fmt.Sprintf("%04d~%04d", (num/100)*100+1, (num/100)*100+99)
	}
	return fmt.Sprintf("%04d~%04d", (num/100)*100, (num/100)*100+99)
}
