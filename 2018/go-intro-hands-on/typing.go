package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	example = []string{
		"golang",
		"hands-on",
		"in",
		"kagawa",
	}

	countdown = 3
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var datas []string
	switch len(os.Args) {
	case 1:
		datas = example
	case 2:
		texts, err := read(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v", err)
			os.Exit(1)
		}
		datas = texts
	default:
		fmt.Println("[usage] typing")
		fmt.Println("        typing datafile")
		os.Exit(0)
	}

	for i := countdown; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	var cortms, mistms, alltms time.Duration
	var corcnt, miscnt int
	var allcnt = len(datas)

	//for i := 0; i < allcnt; i++ {
	//	v, datas := random(datas)
	for i, v := range datas {
		fmt.Printf("[%d/%d] %s\n", i+1, allcnt, v)

		stm := time.Now()
		var ans string
		fmt.Scanln(&ans)
		etm := time.Since(stm)

		alltms += etm

		var mark = "　 "
		if v == ans {
			corcnt = corcnt + 1
			cortms = cortms + etm
			mark = "○"
		} else {
			miscnt++
			mistms += etm
			mark = "×"
		}
		fmt.Println(mark, etm)
	}

	// Output results.
	fmt.Printf("[正誤] 正解:%d 誤り:%d\n", corcnt, miscnt)
	coravg := avg(cortms, corcnt)
	misavg := avg(mistms, miscnt)
	allavg := avg(alltms, allcnt)
	fmt.Printf("[平均] 正解:%s 誤り:%s 全て:%s\n", coravg, misavg, allavg)
}

// avg calcurate time/count avaragea.
func avg(sum time.Duration, cnt int) time.Duration {
	if cnt > 0 {
		return sum / time.Duration(cnt)
	} else {
		return time.Duration(0)
	}
}

// read is to read tet data file.
func read(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var texts []string
	//text := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		texts = append(texts, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return texts, nil
}

// randomText
func randomText(texts []string) (string, []string) {
	s := texts
	ln := len(s)
	i := rand.Intn(ln)
	v := s[i]

	t := s
	//s = []string{}
	s = make([]string, 0, ln)
	s = append(s, t[:i]...)
	s = append(s, t[i+1:]...)

	return v, s
}
