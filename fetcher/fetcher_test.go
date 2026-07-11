package fetcher

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// writePoem 向 dir 下写入一个仅含 title 的最小 .poem 文件，用于测试 fixture。
func writePoem(t *testing.T, dir, name string) {
	t.Helper()
	content := "title: " + name + "\ninfo: fixture\nfrom: check-*\nneed:\n\npoem:\necho x\n-p10 bad= check-have-any\n\ngood: ok\nbad: fail\n"
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
		t.Fatalf("写入 fixture 失败: %v", err)
	}
}

// TestSearchFolderFindsInSubdir 验证：关键词只出现在子目录名里时也能被找到。
// 原来的实现把 "递归进目录" 写在 else if 里，导致目录名已含关键词时会被跳过。
func TestSearchFolderFindsInSubdir(t *testing.T) {
	root := t.TempDir()
	sub := filepath.Join(root, "check")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatalf("创建子目录失败: %v", err)
	}
	writePoem(t, sub, "check-uname-a.poem")

	// 搜索词 "check" 只与子目录名和文件名都有关，但根目录名不含 "check"。
	got := SearchFolder(root, "check")
	if len(got) != 1 {
		t.Fatalf("期望找到 1 个，实际找到 %d 个: %v", len(got), got)
	}
}

// TestSearchFolderKeywordInRootDir 验证：关键词在根目录的文件名里也能被找到。
func TestSearchFolderKeywordInRootDir(t *testing.T) {
	root := t.TempDir()
	writePoem(t, root, "check-uname-a.poem")
	writePoem(t, root, "whoami.poem")

	got := SearchFolder(root, "check")
	sort.Strings(got)
	if len(got) != 1 || filepath.Base(got[0]) != "check-uname-a.poem" {
		t.Fatalf("期望找到 [check-uname-a.poem]，实际: %v", got)
	}
}

// TestSearchFolderFromFile 验证：传入文件路径时从该文件所在目录开始搜索。
func TestSearchFolderFromFile(t *testing.T) {
	root := t.TempDir()
	writePoem(t, root, "check-uname-a.poem")

	file := filepath.Join(root, "check-uname-a.poem")
	got := SearchFolder(file, "check")
	if len(got) != 1 {
		t.Fatalf("期望 1，实际 %d: %v", len(got), got)
	}
}
