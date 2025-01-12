# 🛡️ Flash-LLaMA

<div align="center">

<img src="https://img.shields.io/badge/FLASH--LLAMA-AI%20POWERED-FF6B6B?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCI+PHBhdGggZmlsbD0iI0ZGRiIgZD0iTTEyIDJMMSAxMmgzdjhoNnYtNmgydjZoNnYtOGgzTDEyIDJ6Ii8+PC9zdmc+" alt="Flash-LLaMA" />

[![Go Version](https://img.shields.io/badge/Go-1.23.2-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![Groq AI](https://img.shields.io/badge/Groq-LLaMA--3.1--70B-FF4B4B?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCI+PHBhdGggZmlsbD0iI0ZGRiIgZD0iTTEyIDJMMSAxMmgzdjhoNnYtNmgydjZoNnYtOGgzTDEyIDJ6Ii8+PC9zdmc+)](https://groq.com)
[![Security Scanner](https://img.shields.io/badge/Security-Scanner-2ea44f?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCI+PHBhdGggZmlsbD0iI0ZGRiIgZD0iTTEyIDJMMSAxMmgzdjhoNnYtNmgydjZoNnYtOGgzTDEyIDJ6Ii8+PC9zdmc+)](https://github.com/AmazingMoaaz/Flash-LLaMA)

[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCI+PHBhdGggZmlsbD0iI0ZGRiIgZD0iTTEyIDJMMSAxMmgzdjhoNnYtNmgydjZoNnYtOGgzTDEyIDJ6Ii8+PC9zdmc+)](LICENSE)
[![Maintained](https://img.shields.io/badge/Maintained-YES-brightgreen.svg?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCI+PHBhdGggZmlsbD0iI0ZGRiIgZD0iTTEyIDJMMSAxMmgzdjhoNnYtNmgydjZoNnYtOGgzTDEyIDJ6Ii8+PC9zdmc+)](https://github.com/AmazingMoaaz/Flash-LLaMA)

<br>

[📚 Documentation](#-documentation) •
[🚀 Quick Start](#-quick-start) •
[💡 Examples](#-example-commands) •
[🤝 Contributing](#-contributing)

<br>

```ascii
                        _____ _           _           _     _          __  __    _    
                       |  ___| | __ _ ___| |__       | |   | |    __ _|  \/  |  / \   
                       | |_  | |/ _' / __| '_ \ _____| |   | |   / _' | |\/| | / _ \  
                       |  _| | | (_| \__ \ | | |_____| |___| |__| (_| | |  | |/ ___ \ 
                       |_|   |_|\__,_|___/_| |_|     |_____|_____\__,_|_|  |_/_/   \_\
```

</div>

<div align="center">
<i>🔒 Next-Generation AI-Powered Code Security Scanner</i>
</div>

<br>

<p align="center">
  <b>Flash-LLaMA</b> is a revolutionary code security scanner that harnesses the power of <b>Groq's LLaMA-3.1-70B</b> model to detect vulnerabilities with unprecedented accuracy. Transform your security workflow with AI-driven insights and automated fixes.
</p>

---

## ⚡ Highlights

<table align="center">
<tr>
<td align="center">🧠</td>
<td><b>Advanced AI Analysis</b><br>Powered by LLaMA-3.1-70B</td>
<td align="center">🔍</td>
<td><b>Deep Scanning</b><br>File & Directory Analysis</td>
</tr>
<tr>
<td align="center">📊</td>
<td><b>Rich Reports</b><br>Detailed Markdown Output</td>
<td align="center">🎨</td>
<td><b>Beautiful CLI</b><br>Color-Coded Interface</td>
</tr>
</table>

## 🎯 Key Features

<table>
<tr>
<td width="50%">

### 🔒 Security Features
- 🛡️ Vulnerability Detection
- 🔍 Code Pattern Analysis
- ⚠️ Risk Assessment
- 🛠️ Auto-Fix Suggestions

</td>
<td width="50%">

### 💻 Developer Experience
- 📄 Single File Scanning
- 📁 Directory Analysis
- 🎨 Interactive CLI
- 📊 Markdown Reports

</td>
</tr>
</table>

## 🚀 Quick Start

### 📦 Prerequisites

```bash
🔷 Go 1.23.2+
🔑 Groq API Key
🔄 Git
```

### ⚡ One-Line Installation

```bash
git clone https://github.com/AmazingMoaaz/Flash-LLaMA.git && cd Flash-LLaMA && go build
```

### 🔧 Configuration

<details>
<summary>📄 Using .env file</summary>

```env
GROQ_API_KEY=your_groq_api_key_here
```
</details>

## 💫 Usage

### 🎮 Command Palette

```bash
./flash-llama [options] <target>
```

<table>
<tr>
<th>Option</th>
<th>Description</th>
<th>Example</th>
</tr>
<tr>
<td><code>-file</code></td>
<td>Scan single file</td>
<td><code>-file app.go</code></td>
</tr>
<tr>
<td><code>-dir</code></td>
<td>Scan directory</td>
<td><code>-dir ./src</code></td>
</tr>
<tr>
<td><code>-save</code></td>
<td>Report location</td>
<td><code>-save ./reports</code></td>
</tr>
</table>

### 🌟 Example Commands

```bash
# Scan Single File
./flash-llama -file ./src/app.go -save ./security-reports

# Scan Directory
./flash-llama -dir ./src -save ./vulnerability-reports
```

## 📊 Report Structure

<table>
<tr>
<td width="25%" align="center">🎯<br><b>Vulnerabilities</b></td>
<td width="25%" align="center">🔬<br><b>Analysis</b></td>
<td width="25%" align="center">⚠️<br><b>Risk Levels</b></td>
<td width="25%" align="center">🛠️<br><b>Solutions</b></td>
</tr>
</table>

## 🏗️ Project Structure

```
Flash-LLaMA/
├── 📂 config/
│   ├── 🎨 ascii.go        # CLI Styling
│   └── ⚙️ config.go       # Configuration
├── 📂 internal/
│   ├── 📝 markdown/       # Reports
│   └── 🔍 scanner/        # Analysis
├── 🚀 main.go             # Entry Point
└── ⚙️ .env                # Settings
```

## 🛠️ Development

<table>
<tr>
<td>

### 📦 Dependencies
```go
require (
    github.com/jpoz/groq
    github.com/joho/godotenv
)
```

</td>
<td>

### 🔨 Build Steps
```bash
go mod download
go build
go test ./...
```

</td>
</tr>
</table>

## 🤝 Contributing

<table>
<tr>
<td>1. 🔀 Fork</td>
<td>2. 🌿 Branch</td>
<td>3. 💻 Code</td>
<td>4. 🔍 Test</td>
<td>5. 📤 PR</td>
</tr>
</table>

## 🔐 Security Best Practices

- 🔒 Secure API Key Storage
- ⚠️ No Credentials in VCS
- 👀 Regular Report Reviews

## 📜 License & Credits

<table>
<tr>
<td width="50%">

### 📄 License
MIT License - See [LICENSE](LICENSE)

</td>
<td width="50%">

### 👥 Made by
- 👨‍💻 @AmazingMoaaz

</td>
</tr>
</table>

---

<div align="center">

### 🛡️ Secure Your Code with Flash-LLaMA

<br>

[![Star](https://img.shields.io/github/stars/AmazingMoaaz/Flash-LLaMA?style=for-the-badge&color=yellow&logo=github)](https://github.com/AmazingMoaaz/Flash-LLaMA)
[![Follow](https://img.shields.io/github/followers/AmazingMoaaz?style=for-the-badge&color=blue&logo=github)](https://github.com/AmazingMoaaz)

<br>

<sub>Made with ❤️ by the Security Community</sub>

</div>
