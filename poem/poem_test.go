package poem

import (
	"os"
	"path/filepath"
	"testing"
)

// TestParseExamplePoems 用仓库里真实的 .poem 文件验证解析器，
// 确保字段对应到正确的位置而不依赖行索引。
func TestParseExamplePoems(t *testing.T) {
	files, err := filepath.Glob("../examples/poems/**/*.poem")
	if err != nil {
		t.Fatalf("glob 失败: %v", err)
	}
	if len(files) == 0 {
		t.Fatalf("没有找到任何示例 poem 文件")
	}

	for _, f := range files {
		t.Run(filepath.Base(f), func(t *testing.T) {
			content, err := os.ReadFile(f)
			if err != nil {
				t.Fatalf("读取失败 %s: %v", f, err)
			}
			p := ParsePoem(string(content), map[string]string{"OS_TYPE": "linux"})
			if p.Title == "" {
				t.Errorf("%s: 标题为空", f)
			}
			if len(p.Poem) == 0 {
				t.Errorf("%s: 正文为空", f)
			}
			// 正文应由 (命令行, -p 等待行) 成对组成。
			if len(p.Poem)%2 != 0 {
				t.Errorf("%s: 正文行数 %d 不是偶数（应与 -p 行成对）", f, len(p.Poem))
			}
			for i := 1; i < len(p.Poem); i += 2 {
				if !startsWith(p.Poem[i], "-p") {
					t.Errorf("%s: 第 %d 行 %q 应为 -p 等待行", f, i, p.Poem[i])
				}
			}
		})
	}
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// TestParsePoemFields 直接验证各字段不求诸行索引：调整字段顺序或省略可选字段，
// 仍应正确解析。
func TestParsePoemFields(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		wantTitle string
		wantFrom  []string
		wantNeed  []string
		wantPoem  []string
		wantGood  string
		wantBad   string
	}{
		{
			name: "标准顺序",
			input: `title: 标题
info: 介绍
from: check-* install-*
need: pre1 pre2

poem:
uname -a
-p10 bad= check-have-any

good: 成功
bad: 失败`,
			wantTitle: "标题",
			wantFrom:  []string{"check-*", "install-*"},
			wantNeed:  []string{"pre1", "pre2"},
			wantPoem:  []string{"uname -a", "-p10 bad= check-have-any"},
			wantGood:  "成功",
			wantBad:   "失败",
		},
		{
			name: "调整字段顺序且省略可选字段 info/need",
			input: `title: 仅标题
from: install-*

poem:
echo hi
-p10 bad= check-have-any

good: 好
bad: 坏`,
			wantTitle: "仅标题",
			wantFrom:  []string{"install-*"},
			wantPoem:  []string{"echo hi", "-p10 bad= check-have-any"},
			wantGood:  "好",
			wantBad:   "坏",
		},
		{
			name: "字段之间夹有空行",
			input: `title: 有空行

info: 介绍

from: check-*

need: pre

poem:
echo x
-p10 bad= check-have-any

good: 成
bad: 败`,
			wantTitle: "有空行",
			wantFrom:  []string{"check-*"},
			wantNeed:  []string{"pre"},
			wantPoem:  []string{"echo x", "-p10 bad= check-have-any"},
			wantGood:  "成",
			wantBad:   "败",
		},
		{
			name: "替换参数占位符",
			input: `title: 安装 $&{app}
info: 介绍
from: install-*
need:

poem:
echo $&{app}
-p10 bad= check-have-any

good: 完成
bad: 出错`,
			wantTitle: "安装 nginx",
			wantFrom:  []string{"install-*"},
			wantPoem:  []string{"echo nginx", "-p10 bad= check-have-any"},
			wantGood:  "完成",
			wantBad:   "出错",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			args := map[string]string{"app": "nginx"}
			p := ParsePoem(tc.input, args)
			if p.Title != tc.wantTitle {
				t.Errorf("Title = %q, want %q", p.Title, tc.wantTitle)
			}
			if !equalSlice(p.From, tc.wantFrom) {
				t.Errorf("From = %v, want %v", p.From, tc.wantFrom)
			}
			if !equalSlice(p.Need, tc.wantNeed) {
				t.Errorf("Need = %v, want %v", p.Need, tc.wantNeed)
			}
			if !equalSlice(p.Poem, tc.wantPoem) {
				t.Errorf("Poem = %v, want %v", p.Poem, tc.wantPoem)
			}
			if p.Good != tc.wantGood {
				t.Errorf("Good = %q, want %q", p.Good, tc.wantGood)
			}
			if p.Bad != tc.wantBad {
				t.Errorf("Bad = %q, want %q", p.Bad, tc.wantBad)
			}
		})
	}
}

func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TestFromAllows 验证 glob 匹配逻辑。
func TestFromAllows(t *testing.T) {
	cases := []struct {
		patterns []string
		name     string
		want     bool
	}{
		{[]string{"check-*", "install-*"}, "install-nginx", true},
		{[]string{"check-*", "install-*"}, "check-uname-a", true},
		{[]string{"check-*", "install-*"}, "fix-epel", false},
		{[]string{"*"}, "anything", true},
		{nil, "anything", true},
		{[]string{}, "anything", true},
		{[]string{"install-*"}, "Install-CaseSensitive", false},
	}
	for _, tc := range cases {
		if got := fromAllows(tc.patterns, tc.name); got != tc.want {
			t.Errorf("fromAllows(%v, %q) = %v, want %v", tc.patterns, tc.name, got, tc.want)
		}
	}
}

// TestApplyCheck 逐分支验证输出判定函数。
func TestApplyCheck(t *testing.T) {
	cases := []struct {
		flag string
		data string
		out  string
		want bool
	}{
		{"check-have-any", "", "something", true},
		{"check-have-any", "", "", false},
		{"check-have-any-r", "", "something", false},
		{"check-have-any-r", "", "", true},
		{"check-have-none", "", "", true},
		{"check-have-none", "", "x", false},
		{"check-have-none-r", "", "", false},
		{"check-have-none-r", "", "x", true},
		{"check-have-all", "abc", "abc", true},
		{"check-have-all", "abc", "abcd", false},
		{"check-have-all-r", "abc", "abc", false},
		{"check-have-all-r", "abc", "abcd", true},
		{"check-have", "cat", "a cat sits", true},
		{"check-have", "cat", "a dog sits", false},
		{"check-have-r", "cat", "a cat sits", false},
		{"check-have-r", "cat", "a dog sits", true},
		{"unknown-flag", "", "whatever", false},
	}
	for _, tc := range cases {
		if got := applyCheck(tc.flag, tc.data, tc.out); got != tc.want {
			t.Errorf("applyCheck(%q, %q, %q) = %v, want %v", tc.flag, tc.data, tc.out, got, tc.want)
		}
	}
}

// TestBaseName 验证文件名提取。
func TestBaseName(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"./examples/poems/check/check-uname-a.poem", "check-uname-a"},
		{"fix-epel-release-linux-centos7-amd64.poem", "fix-epel-release-linux-centos7-amd64"},
		{`C:\a\b\whoami-windows.poem`, "whoami-windows"},
	}
	for _, tc := range cases {
		if got := baseName(tc.in); got != tc.want {
			t.Errorf("baseName(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
