# Copilot: Commit Message Skill

目的
- 指导 Copilot/ChatGPT 在生成提交信息时遵循本仓库的提交规范（Conventional Commits 风格）。

规范
- 首行格式：type(scope?): short summary
- 允许的 type：build, chore, ci, docs, feat, fix, perf, refactor, revert, style, test
- 首行长度 <= 50 字符；正文每行 <= 72 字符
- 校验正则（示例）：^(build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test)(\([a-z0-9\-_]+\))?: .{1,50}$

示例
- 合格：feat(auth): add JWT refresh endpoint
- 不合格：Fixed login bug

AI 提示模板（直接用于 Copilot / ChatGPT）
- 简短生成（首选）:
  > 你是一个提交信息助手。请根据下面的“改动说明”生成一个符合本仓库规范（Conventional Commits）的提交消息。  
  > 要求：首行为 `type(scope?): short summary` 且不超过 50 字；之后可选地给出一段不超过 72 字换行的详细描述。  
  > 不要输出任何解释或多余内容，只返回最终提交消息文本。  
  > 改动说明：<在此粘贴改动摘要或已暂存文件列表>

- 带 scope 建议:
  > 同上，但同时自动建议最合适的 scope（小写、连字符），如果不确定可省略 scope。

使用方法
- 在 Copilot Chat 中粘贴改动摘要，调用上面的提示模板；复制输出的提交文本到 `git commit` 编辑器或 `git commit -m`.
- 将该文件加入仓库（`.github/`）后，Copilot 在建议提交信息时会参考此规则与提示。

可选：把下面 hook / 模板 / CI 一并加入仓库以实现本地/远端强制。
