# golang-study

一个用于学习和实践 Go 的示例工程，已采用现代化目录结构，便于扩展与维护。

## 工程目录

```text
.
├── api/                # OpenAPI/接口协议定义（预留）
├── build/              # 构建相关脚本与资源（预留）
├── cmd/
│   └── app/            # 应用程序入口
├── configs/            # 配置文件（预留）
├── deployments/        # 部署清单（预留）
├── docs/               # 项目文档（预留）
├── internal/
│   └── app/            # 仅项目内部使用的业务代码
├── pkg/
│   └── version/        # 可对外复用的公共库代码
├── scripts/            # 自动化脚本（预留）
└── test/               # 集成/端到端测试（预留）
```

## 快速开始

```bash
go test ./...
go run ./cmd/app
```

## Commit Message Generation Skill

This repository integrates a commit message generation skill located at `.github/skills/commit-message-generation-skill.md`. This skill enables automatic generation of Conventional Commits style commit messages using Copilot/ChatGPT or compatible CLI tools.

### Usage
- Summarize your staged changes or provide a list of changed files.
- Call the commit message generation skill (see `.github/skills/commit-message-generation-skill.md` for details).
- Use the output as your commit message, e.g.:

```bash
git diff --cached --name-only > /tmp/commit_changes.txt
CHANGES=$(cat /tmp/commit_changes.txt)
git commit -m "$(copilot-commit-skill --input "$CHANGES")"
```

- For integration with Git hooks, see `.github/skills/prepare-commit-msg.example`.
- For full style rules and prompt templates, see `.github/skills/copilot-commit-skill.md`.