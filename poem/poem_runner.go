package poem

import (
	"context"
	"fmt"
	"org.mutantcat.rhyme/exec"
	"org.mutantcat.rhyme/fetcher"
	exe "os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func RunPoem(poem Poem, su bool, depth int) bool {
	atoi, err := strconv.Atoi(poem.Args["NEED_DEPTH"])
	bad_saver := ""
	check_flag := ""
	check_flag_data := ""
	if err != nil {
		fmt.Println("韵 - 读取最大联调深度错误:", err)
		atoi = 5
	}
	if depth > atoi {
		fmt.Println("诗 -", "超过最大联调深度:", depth, "(将自动结束所有执行的依赖项)")
		return false
	}
	fmt.Println("诗 -", "开始执行诗句:", poem.Title)
	fmt.Println()
	success := true
	// 读取和检查诗句的need
	for _, need := range poem.Need {
		if need != "" {
			fmt.Println("诗 -", "需要的前置验证:", need)
			file := fetcher.CheckAndGetPoemFile(poem.Path, need+".poem")
			if file == "" {
				fmt.Println("韵 - 文件夹内没有对应诗歌文件")
			} else {
				parsePoem := ParsePoem(fetcher.ReadFile(file), poem.Args)
				parsePoem.Path = file
				file = strings.ReplaceAll(file, "\\", "/")
				poemName := strings.Split(file, "/")
				poemName = strings.Split(poemName[len(poemName)-1], ".")
				parsePoem.Name = poemName[0]
				// 先验证当前要运行的诗歌是否满足前置诗歌中from中的正则

				can_do_pre := false
				for _, from := range parsePoem.From {
					// 编译正则表达式
					re, err := regexp.Compile(from)
					if err != nil {
						fmt.Println("韵 - 前置中正则表达式错误:", err)
					}
					if re.MatchString(poem.Name) {
						can_do_pre = true
					}
				}
				if !can_do_pre {
					success = false
					fmt.Println("韵 - 前置规则匹配失败:", poem.Name)
					return success
				}

				if RunPoem(*parsePoem, su, depth+1) {
					fmt.Println("诗 -", "前置验证成功:", need)
					fmt.Println()
				} else {
					fmt.Println()
					fmt.Println("诗 -", "前置诗歌验证失败:", need)
					success = false
					fmt.Println("诗 -", "前置诗歌执行结果:", parsePoem.Bad)
					return success
				}
			}
		}
	}

	cmd := &exe.Cmd{}
	line_number := 1
	for index := 0; index < len(poem.Poem)-1; index++ {
		// 读取命令行和等待时间行
		cmdLine := poem.Poem[index]
		waitLine := poem.Poem[index+1]
		if !strings.HasPrefix(waitLine, "-p") {
			fmt.Println("诗 -", "无效等待时间:", waitLine)
			success = false
			fmt.Println("诗 -", "执行结果:", poem.Bad)
			return success
		}

		// 解析等待时间
		waitCmd := strings.Split(strings.TrimPrefix(waitLine, "-p"), " ")
		waitTimeStr := waitCmd[0] // 等待时间
		waitTime, err := strconv.Atoi(waitTimeStr)
		if err != nil {
			fmt.Println("解析等待时间错误:", err)
			success = false
			fmt.Println("诗 -", "执行结果:", poem.Bad)
			return success
		}

		if strings.HasPrefix(waitCmd[1], "bad=") {
			bad_saver = strings.TrimPrefix(waitCmd[1], "bad=")
		}

		if len(waitCmd) > 2 {
			check_flags := strings.Split(waitCmd[2], "=")
			check_flag = check_flags[0]
			if len(check_flags) > 1 {
				check_flag_data = check_flags[1]
			}
		}

		// 执行命令并设置超时
		fmt.Println("诗 -", "正在执行命令", poem.Name, line_number, ":", cmdLine, "(最长", waitTime, "秒)")
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(waitTime)*time.Second)
		defer cancel()
		cmdLine = strings.TrimPrefix(cmdLine, "su ") // 如果原本有一个su加权 则去掉 之后由系统主动指定
		cmd := exec.Run(ctx, cmd, cmdLine, su)       // 假设是Linux命令
		output, err := cmd.Output()

		// 处理结果
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("诗 -", "命令超时取消", poem.Name, line_number, ":", cmdLine)
			success = false
			fmt.Println()
			fmt.Println("诗 -", "执行结果:", poem.Bad)
			return success
		} else if err != nil {
			fmt.Println("诗 -", "执行命令错误", poem.Name, line_number, ":", err)
			if bad_saver != "" {
				fmt.Println("诗 -", "尝试修复问题:", bad_saver)
				file := fetcher.CheckAndGetPoemFile(poem.Path, bad_saver+".poem")
				if file == "" {
					fmt.Println("韵 - 文件夹内没有对应诗歌文件")
				} else {
					parsePoem := ParsePoem(fetcher.ReadFile(file), poem.Args)
					parsePoem.Path = file
					file = strings.ReplaceAll(file, "\\", "/")
					poemName := strings.Split(file, "/")
					poemName = strings.Split(poemName[len(poemName)-1], ".")
					parsePoem.Name = poemName[0]
					// 先验证当前要运行的诗歌是否满足前置诗歌中from中的正则

					can_do_pre := false
					for _, from := range parsePoem.From {
						// 编译正则表达式
						re, err := regexp.Compile(from)
						if err != nil {
							fmt.Println("韵 - 修复中正则表达式错误:", err)
						}
						if re.MatchString(poem.Name) {
							can_do_pre = true
						}
					}
					if !can_do_pre {
						success = false
						fmt.Println("韵 - 修复规则匹配失败:", poem.Name)
						return success
					}

					if RunPoem(*parsePoem, su, depth+1) {
						fmt.Println("诗 -", "修复验证成功:", bad_saver)
					} else {
						fmt.Println()
						fmt.Println("诗 -", "修复诗歌验证失败:", bad_saver)
						success = false
						fmt.Println("诗 -", "修复诗歌执行结果:", parsePoem.Bad)
						return success
					}
				}
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(waitTime)*time.Second)
				defer cancel()
				cmd := exec.Run(ctx, cmd, cmdLine, su)
				output, err := cmd.Output()
				if err != nil {
					fmt.Println("诗 -", "修复后命令仍错误", poem.Name, line_number, ":", err)
					success = false
					fmt.Println("诗 -", "执行结果:", poem.Bad)
					return success
				} else {
					fmt.Println("诗 -", "修复后命令执行成功", poem.Name, line_number, ":", string(output))
					success = true
				}
			}
			return success
		} else {
			fmt.Println()
			fmt.Println("诗 -", "命令执行成功", poem.Name, line_number, ":", string(output))
			success = true
		}

		switch check_flag {
		case "check-have-any":
			if len(string(output)) > 0 {
				success = true
			} else {
				success = false
			}
			break
		case "check-have-any-r":
			if len(string(output)) > 0 {
				success = false
			} else {
				success = true
			}
			break
		case "check-have-none":
			if len(string(output)) == 0 {
				success = true
			} else {
				success = false
			}
			break
		case "check-have-none-r":
			if len(string(output)) == 0 {
				success = false
			} else {
				success = true
			}
			break
		case "check-have-all":
			if string(output) == check_flag_data {
				success = true
			} else {
				success = false
			}
			break
		case "check-have-all-r":
			if string(output) == check_flag_data {
				success = false
			} else {
				success = true
			}
			break
		case "check-have":
			if strings.Contains(string(output), check_flag_data) {
				success = true
			} else {
				success = false
			}
			break
		case "check-have-r":
			if strings.Contains(string(output), check_flag_data) {
				success = false
			} else {
				success = true
			}
			break
		}

		// 跳过命令行（因为已处理）
		index = index + 1
		line_number++
	}

	if success {
		fmt.Println("诗 -", "执行结果:", poem.Good)
	} else {
		fmt.Println("诗 -", "执行结果:", poem.Bad)
	}
	return success
}
