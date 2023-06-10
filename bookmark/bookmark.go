package bookmark

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"time"
)

// bookmarkWalker 缓存书签目录栈
type bookmarkWalker struct {
	dirStack []string
}

// Item 书签条目
type Item struct {
	Name string   `json:"name"`
	Dirs []string `json:"dirs"`
	Href string   `json:"href"`
	Time string   `json:"time"`
}

// itemHanlder 条目处理器
type itemHanlder func(*Item) error

var (
	dirRe   = regexp.MustCompile(`^\s*<DT><H3 .*>(.*)</H3>$`)
	itemRe  = regexp.MustCompile(`^\s*<DT><A HREF="(.*)" ADD_DATE="([0-9]+)".*>(.*)</A>$`)
	listEnd = regexp.MustCompile(`^\s*</DL><p>$`)
)

func Scan(path string, hanlder itemHanlder) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	bw := &bookmarkWalker{}
	return bw.walk(file, hanlder)
}

func (bm *bookmarkWalker) walk(file *os.File, handler itemHanlder) error {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		// 条目
		item, err := bm.matchItem(line)
		if err != nil {
			return err
		}
		if item != nil {
			if err = handler(item); err != nil {
				return err
			}
			continue
		}

		// 进入目录
		matchDir := dirRe.FindStringSubmatch(line)
		if len(matchDir) > 0 {
			bm.dirStack = append(bm.dirStack, matchDir[1])
			continue
		}

		// 离开目录
		if listEnd.MatchString(line) && len(bm.dirStack) > 1 {
			bm.dirStack = bm.dirStack[:len(bm.dirStack)-1]
			continue
		}
	}

	return nil
}

func (bm *bookmarkWalker) matchItem(line string) (*Item, error) {
	m := itemRe.FindStringSubmatch(line)
	if len(m) == 0 {
		return nil, nil
	}

	seconds, err := strconv.ParseInt(m[2], 0, 64)
	if err != nil {
		return nil, err
	}
	timeStr := time.Unix(seconds, 0).Format(time.RFC3339)

	dirs := make([]string, len(bm.dirStack))
	copy(dirs, bm.dirStack)

	i := &Item{
		Name: m[3],
		Dirs: dirs,
		Href: m[1],
		Time: timeStr,
	}

	return i, nil
}
