package poem

import (
	"context"
	"fmt"
	"org.mutantcat.rhyme/exec"
	"org.mutantcat.rhyme/fetcher"
	"path"
	"strconv"
	"strings"
	"time"
)

// fromAllows 按 shell 风格 glob 判断 name 是否匹配 patterns 中的任一规则。
// from 字段使用 glob（如 check-*、install-*），空列表表示不做限制。
func fromAllows(patterns []string, name string) bool {
	if len(patterns) == 0 {
		return true
	}
	for _, p := range patterns {
		if ok, err := path.Match(p, name); err == nil && ok {
			return true
		}
	}
	return false
}

// runLine 执行一条命令行并附带超时。返回输出、是否超时、以及错误。
// 超时通过 context 判定，cancel 在返回时立即执行，避免在循环中泄漏。
func runLine(parent context.Context, line string, su bool, waitTime int) ([]byte, bool, error) {
	ctx, cancel := context.WithTimeout(parent, time.Duration(waitTime)*time.Second)
	defer cancel()
	line = strings.TrimPrefix(line, "su ")
	cmd := exec.Run(ctx, line, su)
	output, err := cmd.Output()
	if ctx.Err() == context.DeadlineExceeded {
		return output, true, err
	}
	return output, false, err
}

// baseName 取文件路径去目录、去后缀后的名称。
func baseName(file string) string {
	file = strings.ReplaceAll(file, "\\", "/")
	parts := strings.Split(file, "/")
	name := parts[len(parts)-1]
	if idx := strings.LastIndex(name, "."); idx > 0 {
		name = name[:idx]
	}
	return name
}

func RunPoem(poem Poem, su bool, depth int) bool {
	maxDepth := 5
	if d, err := strconv.Atoi(poem.Args["NEED_DEPTH"]); err == nil {
		maxDepth = d
	} else {
		fmt.Println("韵 - 读取最大联调深度错误:", err)
	}
	if depth > maxDepth {
		fmt.Println("诗 -", "超过最大联调深度:", depth, "(将自动结束所有执行的依赖项)")
		return false
	}
	fmt.Println("诗 -", "开始执行诗句:", poem.Title)
	fmt.Println()

	// 递归执行前置依赖。
	for _, need := range poem.Need {
		if need == "" {
			continue
		}
		fmt.Println("诗 -", "需要的前置验证:", need)
		file := fetcher.CheckAndGetPoemFile(poem.Path, need+".poem")
		if file == "" {
			fmt.Println("韵 - 文件夹内没有对应诗歌文件")
			continue
		}
		sub := ParsePoem(fetcher.ReadFile(file), poem.Args)
		sub.Path = file
		sub.Name = baseName(file)
		if !fromAllows(sub.From, poem.Name) {
			fmt.Println("韵 - 前置规则匹配失败:", poem.Name)
			return false
		}
		if RunPoem(*sub, su, depth+1) {
			fmt.Println("诗 -", "前置验证成功:", need)
			fmt.Println()
		} else {
			fmt.Println()
			fmt.Println("诗 -", "前置诗歌验证失败:", need)
			fmt.Println("诗 -", "前置诗歌执行结果:", sub.Bad)
			return false
		}
	}

	success := true
	line_number := 1
	// 正文按 (命令行, 等待行) 成对消费，每次前进两行。
	for index := 0; index < len(poem.Poem)-1; index += 2 {
		cmdLine := poem.Poem[index]
		waitLine := poem.Poem[index+1]
		if !strings.HasPrefix(waitLine, "-p") {
			fmt.Println("诗 -", "无效等待时间:", waitLine)
			fmt.Println("诗 -", "执行结果:", poem.Bad)
			return false
		}

		waitCmd := strings.Split(strings.TrimPrefix(waitLine, "-p"), " ")
		waitTime, err := strconv.Atoi(waitCmd[0])
		if err != nil {
			fmt.Println("解析等待时间错误:", err)
			fmt.Println("诗 -", "执行结果:", poem.Bad)
			return false
		}

		bad_saver := ""
		if len(waitCmd) > 1 && strings.HasPrefix(waitCmd[1], "bad=") {
			bad_saver = strings.TrimPrefix(waitCmd[1], "bad=")
		}

		check_flag := ""
		check_flag_data := ""
		if len(waitCmd) > 2 {
			parts := strings.SplitN(waitCmd[2], "=", 2)
			check_flag = parts[0]
			if len(parts) > 1 {
				check_flag_data = parts[1]
			}
		}

		output, timedOut, err := runLine(context.Background(), cmdLine, su, waitTime)
		if timedOut {
			fmt.Println("诗 -", "命令超时取消", poem.Name, line_number, ":", cmdLine)
			fmt.Println()
			fmt.Println("诗 -", "执行结果:", poem.Bad)
			return false
		}

		if err != nil {
			fmt.Println("诗 -", "执行命令错误", poem.Name, line_number, ":", err)
			if !runFix(poem, bad_saver, cmdLine, su, depth, waitTime, line_number) {
				fmt.Println("诗 -", "执行结果:", poem.Bad)
				return false
			}
			// 修复失败；runFix 内部未成功即已返回 false 到达此处。
			return false
		}

		fmt.Println()
		fmt.Println("诗 -", "命令执行成功", poem.Name, line_number, ":", string(output))
		success = true

		// 仅在等待行携带判定标志时，按输出决定本条成败。
		if check_flag != "" {
			success = applyCheck(check_flag, check_flag_data, string(output))
		}

		line_number++
	}

	if poem.Good == "" {
		poem.Good = "没有设置成功执行结果(成功)"
	}
	if poem.Bad == "" {
		poem.Bad = "没有设置失败执行结果(失败)"
	}
	if success {
		fmt.Println("诗 -", "执行结果:", poem.Good)
	} else {
		fmt.Println("诗 -", "执行结果:", poem.Bad)
	}
	return success
}

// runFix 尝试调用 bad= 指向的修复诗，并在修复成功后重试原命令。
// 仅在修复成功且重试命令也成功时返回 true。
func runFix(poem Poem, bad_saver, cmdLine string, su bool, depth, waitTime, line_number int) bool {
	if bad_saver == "" {
		return false
	}
	fmt.Println("诗 -", "尝试修复问题:", bad_saver)
	file := fetcher.CheckAndGetPoemFile(poem.Path, bad_saver+".poem")
	if file == "" {
		fmt.Println("韵 - 文件夹内没有对应诗歌文件")
		return false
	}
	sub := ParsePoem(fetcher.ReadFile(file), poem.Args)
	sub.Path = file
	sub.Name = baseName(file)
	if !fromAllows(sub.From, poem.Name) {
		fmt.Println("韵 - 修复规则匹配失败:", poem.Name)
		return false
	}
	if !RunPoem(*sub, su, depth+1) {
		fmt.Println()
		fmt.Println("诗 -", "修复诗歌验证失败:", bad_saver)
		fmt.Println("诗 -", "修复诗歌执行结果:", sub.Bad)
		return false
	}
	fmt.Println("诗 -", "修复验证成功:", bad_saver)

	// 修复成功，重试原命令。
	output, timedOut, err := runLine(context.Background(), cmdLine, su, waitTime)
	if timedOut {
		fmt.Println("诗 -", "修复后命令超时取消", poem.Name, line_number, ":", cmdLine)
		return false
	}
	if err != nil {
		fmt.Println("诗 -", "修复后命令仍错误", poem.Name, line_number, ":", err)
		return false
	}
	fmt.Println("诗 -", "修复后命令执行成功", poem.Name, line_number, ":", string(output))
	return true
}

func applyCheck(flag, data, output string) bool {
	switch flag {
	case "check-have-any":
		return len(output) > 0
	case "check-have-any-r":
		return len(output) == 0
	case "check-have-none":
		return len(output) == 0
	case "check-have-none-r":
		return len(output) > 0
	case "check-have-all":
		return output == data
	case "check-have-all-r":
		return output != data
	case "check-have":
		return strings.Contains(output, data)
	case "check-have-r":
		return !strings.Contains(output, data)
	}
	return false
}
