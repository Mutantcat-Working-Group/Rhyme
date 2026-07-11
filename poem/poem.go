package poem

type Poem struct {
	Name   string            // 诗歌名称
	Origin string            // 原始文本
	Title  string            // 标题
	Info   string            // 信息
	From   []string          // 允许的来源（glob，用于被 need 与 bad 引用时校验）
	Need   []string          // 需要的前置验证
	Poem   []string          // 诗句（命令行与 -p 等待行交替）
	Good   string            // 成功时输出内容
	Bad    string            // 失败时输出内容
	Args   map[string]string // 参数
	Path   string            // 文件路径
}
