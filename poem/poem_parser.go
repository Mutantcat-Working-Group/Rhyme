package poem

import "strings"

// ParsePoem 解析 .poem 文本为结构体。
//
// 解析基于行前缀（title:/info:/from:/need:/poem:/good:/bad:），不依赖行的绝对位置，
// 因此字段可以省略、重排，字段之间也可以穿插空行。
//
// 遇到 poem: 进入正文段，其后所有非空行（已 trim）都视为诗句，直到遇到 good:/bad: 结束。
// 正文里命令行应与 -p 等待行成对出现，该不变量由 RunPoem 在运行时校验。
func ParsePoem(origin string, args map[string]string) *Poem {
	// 将原文中的$&{abc} 替换为args["abc"]
	for key, value := range args {
		origin = strings.ReplaceAll(origin, "$&{"+key+"}", value)
	}

	p := &Poem{}
	p.Origin = origin
	p.Args = args
	p.Poem = make([]string, 0)

	lines := strings.Split(origin, "\n")
	for i := range lines {
		lines[i] = strings.ReplaceAll(lines[i], "\r", "")
	}

	inPoem := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// poem: 之后进入正文，good:/bad: 会终止正文段。
		if inPoem {
			if strings.HasPrefix(trimmed, "good:") {
				p.Good = strings.TrimSpace(strings.TrimPrefix(trimmed, "good:"))
				inPoem = false
				continue
			}
			if strings.HasPrefix(trimmed, "bad:") {
				p.Bad = strings.TrimSpace(strings.TrimPrefix(trimmed, "bad:"))
				inPoem = false
				continue
			}
			// 跳过空行（诗集里空行仅作分隔，不构成命令）；其余行 trim 后保留。
			if trimmed == "" {
				continue
			}
			p.Poem = append(p.Poem, trimmed)
			continue
		}

		switch {
		case strings.HasPrefix(trimmed, "title:"):
			p.Title = strings.TrimSpace(strings.TrimPrefix(trimmed, "title:"))
		case strings.HasPrefix(trimmed, "info:"):
			p.Info = strings.TrimSpace(strings.TrimPrefix(trimmed, "info:"))
		case strings.HasPrefix(trimmed, "from:"):
			p.From = splitFields(strings.TrimPrefix(trimmed, "from:"))
		case strings.HasPrefix(trimmed, "need:"):
			p.Need = splitFields(strings.TrimPrefix(trimmed, "need:"))
		case strings.HasPrefix(trimmed, "poem:"):
			inPoem = true
		case strings.HasPrefix(trimmed, "good:"):
			p.Good = strings.TrimSpace(strings.TrimPrefix(trimmed, "good:"))
		case strings.HasPrefix(trimmed, "bad:"):
			p.Bad = strings.TrimSpace(strings.TrimPrefix(trimmed, "bad:"))
		}
	}

	return p
}

// splitFields 按空白切分并去掉空串，避免 strings.Split 产生空元素。
func splitFields(s string) []string {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return nil
	}
	return fields
}
