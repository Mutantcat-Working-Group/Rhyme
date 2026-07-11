# 版本变更日志（CHANGELOG）

所有对本项目的显著变更都会记录在此文件。

格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)，
版本号遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

每个版本的变更分两组：
- **本次已合并**：已落地、经端到端验证的改动；
- **后续待做**：已规划但尚未实现的部分，标注 `[待修]`，完成后移入改组并更新日期。

---

## [1.0.20260711] - 2026-07-11

### 本次已合并（稳定性与正确性修复）

#### 修复
- **执行器跨平台 panic**：原 `exec.Run` 在 macOS / BSD 等非 Linux/Windows 系统返回 `nil`，
  调用方 `cmd.Output()` 触发 nil 指针解引用崩溃。重写为 Windows 用 `cmd /C`，
  其余所有系统统一 `bash -c` / `sudo bash -c`，不再返回 nil。
- **`from:` 来源匹配 nil 崩溃**：原实现将 `from:` 当正则编译，示例诗中的 `check-*`、`install-*`
  作为正则非法，返回 nil `*regexp.Regexp`，随后 `re.MatchString` 崩溃 — 所有跨诗
  `need` / `bad` 链路均不可用。改为 `path.Match` glob 匹配，与 README 示例语义一致。
- **正则编译错误被静默吞掉后仍继续解引用 nil**：随上条一并根治。
- **命令失败后漏报成功**：原流程中当命令执行出错且无 `bad=` 兜底时，`success` 未被置
  `false`，最终返回 `true`。现已统一走失败路径返回 `false`。
- **搜索跳过含关键词的子目录**：原 `else if file.IsDir` 导致目录名含关键词（如 `check/`）
  时既不被收集也不被递归，搜索结果恒为空。改为目录一律递归。
- **脆弱的行索引解析器**：原解析器依赖"title 在第 0 行、from 在第 2 行"这类绝对位置，
  缺少任一字段即全体错位。重写为基于行前缀的解析，支持字段省略、重排、穿插空行。
- **`defer cancel()` 在 for 循环中泄漏**：提取 `runLine` 辅助函数，每次调用即时
  `cancel()`，不再延迟到函数返回。
- **搜索输出拼写错误**：`rayme` → `rhyme`。

#### 清理
- 移除 `Poem` 结构体中从未读写的 `Url` / `CloudStep` 死字段。
- 移除 `cmd := &exe.Cmd{}` 外层的死变量（该参数已被内层覆盖）。
- 删除三个从未被任何代码引用的空壳文件：`fetch_from_web.go`、`fetch_from_ai.go`、
  `fetch_from_cloudstep.go`（README 原称"正在开发中"，现已改为本地文件 + 目录检索为主）。
- `go mod tidy` 移除五个从未被 import 的间接依赖（progressbar、colorstring、uniseg、
  golang.org/x/sys、golang.org/x/term）。

#### 测试
- 新增 `poem/poem_test.go`：覆盖真实示例诗解析、解析器字段顺序无关性、from 的 glob
  匹配、全部 8 个 check 判定分支、文件名提取。
- 新增 `fetcher/fetcher_test.go`：覆盖子目录搜索回归（即前述"目录名含关键词被跳过"的缺陷）。
- CI（`.github/workflows/ubuntu.yml`）新增 `go vet` 步骤；测试步骤保留，目前由真正用例驱动。

#### 维护
- 版本号从原来硬编码在 banner / README 多处的方式，收拢为 `main.go` 内的 `version`  常量，避免漂移。

---

### 后续待做 [待修]

> 以下条目已规划但尚未实现，完成后移入上方改组、更新版本号与日期。

#### 功能
- [ ] **`from:` 来源规则支持正则Or glob 显式切换**：当前统一为 glob。后续可引入前缀
  约定（如 `re:^check-.*$`）让用户选择正则，避免既有 poems 改写法。
- [ ] **正式落地"网络拉取 / 云阶 / AI 生成诗"入口**：当前相关 flag（`-web`、`-cloudstep`）
  与 `fetch_from_*` 文件已清理，待设计清晰后以全新实现接入，而非恢复空壳。
- [ ] **`need:` 支持参数化与失败短路**：当前 need 诗失败仅返回失败信息，可考虑暴露具体
  失败行号与最近一次命令输出，便于定位。

#### 稳定性
- [ ] **`RunPoem` 返回 `(bool, error)` 替代纯 `bool`**：把"超过最大联调深度"、
  "正则错误"等可恢复错误通过 error 返回，方便调用方区分"业务失败"与"程序异常"。
- [ ] **命令失败重试次数可配置**：固定重试一次（修复一次），可考虑通过 `-p` 行扩展
  语法支持 `retry=2`。

#### 可观测性
- [ ] **执行日志分级**（info/warn/error）与可选静默模式（`-quiet`），便于嵌入自动化脚本。
- [ ] `--version` 子命令直接输出版本号，无需靠解析 banner。

#### 工程化
- [ ] **集成测试**：构造一个无外部依赖的"假命令"（如基于 `/bin/echo` / `cmd /C echo`）
  在 CI 内跑通 `need` → 主诗 → `bad` 修复的全链路。
- [ ] **Release 产物**：GitHub Actions 在 tag 推送时自动交叉编译并生成多平台 artifact。
- [ ] **README 文档同步**：参数表中 `search` 的类型（标为"字符串"实际是 bool）、
  `examples` 目录拼写（`exapmle` → `example`）、去掉的"网络/AI 功能正在开发中"
  描述需校对更新。

---

_新增版本时，请将上一版"后续待做"中已完成的条目移入该版本"改组"，并保留未完成项。_
