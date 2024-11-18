# Flash-LLaMA

<div align="center">

[![Go](https://img.shields.io/badge/Go-1.19-blue.svg?style=flat&logo=go)](https://golang.org)  
[![Open Source](https://img.shields.io/badge/Open%20Source-%F0%9F%92%9A-brightgreen?style=flat)](https://opensource.org)

</div>

**Flash-LLaMA** is an AI-powered code vulnerability scanner designed to help developers and cybersecurity professionals identify security vulnerabilities in their code. By leveraging **Groq's advanced LLaMA models**, Flash-LLaMA automates the code review process for various programming languages, providing detailed reports that include detected vulnerabilities, proof of concepts, and recommended fixes. It seamlessly integrates security analysis into your development workflow with Markdown-based reporting.

## Features

- **AI-Powered Code Analysis**: Leverages Groq's LLaMA models to identify vulnerabilities like SQL injection, XSS, CSRF, and insecure file handling.
- **Multi-Language Support**: Scans code written in PHP, Python, JavaScript, and more.
- **Markdown Report Generation**: Outputs security analysis in Markdown format for easy integration with GitHub and other platforms.
- **Customizable Scanning**: Analyze individual files or entire directories for vulnerabilities.
- **Graceful Shutdown**: Supports signal handling for secure termination during scans.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Moaz-Awad/Flash-LLaMA.git
    ```

2. Navigate to the project directory:
    ```bash
    cd flash
    ```

3. Build the application:
    ```bash
    go build
    ```

4. Run the application:
    ```bash
    go run main.go -file <codefile> -save <outputdir> -config config.json
    ```

## Usage

Flash scans code files for vulnerabilities by sending code snippets to the LLaMA model, which then returns a detailed analysis of the vulnerabilities. The results can be saved as Markdown reports.

### Command-line Options:

- `-file`: Path to the code file to be analyzed.
- `-dir`: Path to the directory of files to be analyzed.
- `-save`: Directory to save the results (default is current directory).
  
### Example:

```bash
go run main.go -file example.php -save /home/reports -config config.json
```

### Setup YOUR API Key Form groq

The `main.go` file contains the LLaMA API endpoint. Configure it as shown below:

```
// Initialize the Groq client with the API key from the config
client := groq.NewClient(groq.WithAPIKey("YOUR_API_KEY")) // Replace YOUR_API_KEY with your Groq API key
```

Replace YOUR_API_KEY with your Groq API key. Ensure the file is correctly configured before running the tool.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, feel free to reach out to the repository owner.

---
Developed by @Secfathy

Modified by @Moaz-Awad
