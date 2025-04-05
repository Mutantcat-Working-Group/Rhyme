package poem

import "strings"

func ParsePoem(origin string, args map[string]string) *Poem {
	// 将原文中的$&{abc} 替换为args["abc"]
	for key, value := range args {
		origin = strings.ReplaceAll(origin, "$&{"+key+"}", value)
	}

	// 解析诗句
	poem := &Poem{}
	// 保存原文
	poem.Origin = origin
	poem.Args = args
	poem.Poem = make([]string, 0)

	lines := strings.Split(origin, "\r\n")
	// 读取第一行 中的内容到标题 前面的title:可以省略
	mode := 0
	for index, line := range lines {
		if mode == 0 {
			if strings.HasPrefix(line, "title:") || index == 0 {
				poem.Title = strings.TrimPrefix(line, "title:")
				poem.Title = strings.TrimPrefix(poem.Title, " ")
			} else if strings.HasPrefix(line, "info:") || index == 1 {
				poem.Info = strings.TrimPrefix(line, "info:")
				poem.Info = strings.TrimPrefix(poem.Info, " ")
			} else if strings.HasPrefix(line, "from:") || index == 2 {
				poem.From = strings.Split(strings.TrimPrefix(strings.TrimPrefix(line, "from:"), " "), " ")
				// 去掉每部分的前后空格
				for i := 0; i < len(poem.From); i++ {
					poem.From[i] = strings.TrimSpace(poem.From[i])
				}
			} else if strings.HasPrefix(line, "need:") || index == 3 {
				poem.Need = strings.Split(strings.TrimPrefix(strings.TrimPrefix(line, "need:"), " "), " ")
				// 去掉每部分的前后空格
				for i := 0; i < len(poem.Need); i++ {
					poem.Need[i] = strings.TrimSpace(poem.Need[i])
				}
			} else if line == "" {
				mode = 1
			}
		} else if mode == 1 {
			if line == "poem:" || index == 5 {
				continue
			} else if line != "" {
				poem.Poem = append(poem.Poem, line)
			} else {
				mode = 2
			}
		} else if mode == 2 {
			if strings.HasPrefix(line, "good:") {
				poem.Good = strings.TrimPrefix(line, "good:")
				// 去掉前后空格
				poem.Good = strings.TrimPrefix(poem.Good, " ")
			} else if strings.HasPrefix(line, "bad:") {
				poem.Bad = strings.TrimPrefix(line, "bad:")
				// 去掉前后空格
				poem.Bad = strings.TrimPrefix(poem.Bad, " ")
			} else if line == "" {
				mode = 3
			}
		}
	}

	return poem
}
