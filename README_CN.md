[中文版](README_CN.md) | [English version](README.md)

[![Release](https://github.com/bynow2code/urlcheck/actions/workflows/release.yml/badge.svg)](https://github.com/bynow2code/urlcheck/actions/workflows/release.yml)

# urlcheck：并发URL可用性检测工具

快速检测批量URL的可用性，支持并发控制、超时设置、结果导出CSV。

## 功能
- 从命令行或文件读取URL列表
- 自定义并发数和超时时间
- 输出状态码、响应耗时、失败原因
- 支持结果导出为CSV文件

## 使用示例
### 检测单个URL
```
./urlcheck https://baidu.com
```

### 从文件读取，并发20，超时3秒，导出结果
```
./urlcheck -f urls.txt -c 20 -t 3 -o result.csv
```

## 快速安装

### 方法一：一键安装脚本（推荐）

适用于 Linux 和 macOS 系统：

```
curl -sfL https://raw.githubusercontent.com/bynow2code/urlcheck/main/install.sh | bash
```

### 方法二：手动下载

访问 [GitHub Releases](https://github.com/bynow2code/urlcheck/releases/latest) 页面下载适合您系统的预编译版本：

| 系统 | 架构 | 文件名 |
|------|------|--------|
| Windows | AMD64 | `urlcheck-windows-amd64.exe` |
| macOS | AMD64 | `urlcheck-darwin-amd64` |
| macOS | ARM64 | `urlcheck-darwin-arm64` |
| Linux | AMD64 | `urlcheck-linux-amd64` |
| Linux | ARM64 | `urlcheck-linux-arm64` |

下载后，Windows 用户需要将文件重命名为 `urlcheck.exe`，其他系统请根据需要添加执行权限。

### 验证安装

安装完成后，运行以下命令验证是否安装成功：

```
urlcheck -v
```

## 支持

如果你觉得这个工具对你有帮助，请给项目点个star！✨
