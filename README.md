# PQL

> Player Quick Launcher —— 一个使用 Wails、Go 与 Vue 3 构建的桌面快捷启动应用。

> 当前项目仍处于开发阶段，部分页面或功能可能尚未完善。

## 功能特性

- 游戏管理与快捷启动
- 直播
- 视频与动漫内容浏览、搜索和播放
- 可导入、导出及管理解析源
- 下载管理、应用设置与主题跟随

## 环境要求

开始前请安装：

- Go 1.25+
- Node.js 22+
- pnpm
- Wails v3 CLI

安装 Wails v3 CLI：

```bash
go install github.com/wailsapp/wails/v3/cmd/wails3@latest
```

不同操作系统所需的 WebView、编译器等依赖，请参考 [Wails 官方文档](https://v3alpha.wails.io/getting-started/installation/)。

## 快速开始

克隆项目并进入目录：

```bash
git clone https://github.com/yx1126/PQL
```

```bash
cd PQL
```

安装前端依赖：

```bash
cd frontend
pnpm install
cd ..
```

启动开发环境：

```bash
wails3 dev
```

首次执行时，构建任务也会自动安装前端依赖并生成前后端绑定代码。

## 构建与打包

构建当前平台的可执行文件：

```bash
wails3 build
```

构建产物默认输出到 `bin/` 目录。


## 项目结构

```text
PQL/
├── DB/                 # SQLite 模型及数据访问层
├── build/              # Wails 多平台构建与打包配置
├── frontend/           # Vue 3 前端项目
│   ├── bindings/       # Wails 自动生成的 TypeScript 绑定
│   ├── public/         # 静态资源
│   └── src/            # 页面、组件、状态管理与业务逻辑
├── pkg/
│   ├── parse/          # 视频及直播解析逻辑
│   ├── service/        # 暴露给前端的后端服务
│   └── utils/          # 通用工具
├── main.go             # 桌面应用入口
├── systray.go          # 系统托盘逻辑
└── Taskfile.yml        # 项目任务入口
```

## 数据与配置

- 开发模式下，本地数据库保存在 `bin/.PQL/dbs/`。
- 生产模式下，本地数据库保存在应用工作目录的 `.PQL/dbs/`。

## 许可证

本项目基于 [MIT License](./LICENSE) 开源。

## 免责声明

本项目仅用于技术学习与交流。使用者应遵守所在地法律法规以及相关平台的服务条款，并对使用本项目产生的行为和后果自行负责。

# 鸣谢

- 感谢 [xuanxuan123xuan/dart_simple_live](https://github.com/xuanxuan123xuan/dart_simple_live)
- 感谢 [wbt5/real-url](https://github.com/wbt5/real-url)