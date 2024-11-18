# Flash

<div align="center">

[![Go](https://img.shields.io/badge/Go-1.19-blue.svg?style=flat&logo=go)](https://golang.org)  
[![Open Source](https://img.shields.io/badge/Open%20Source-%F0%9F%92%9A-brightgreen?style=flat)](https://opensource.org)

</div>

**Flash** is an AI-powered code vulnerability scanner designed to help developers identify security vulnerabilities in their code. By leveraging open-source AI models like **LLaMA**, Flash automates the review process for various coding languages and provides detailed reports with potential vulnerabilities, proof of concepts, and recommended fixes. Flash can generate reports in Markdown format, making it easy for developers to integrate security analysis into their workflow.

## Features

- **AI-Powered Code Analysis**: Uses LLaMA models to analyze code and detect potential security vulnerabilities.
- **Multi-Platform Support**: Works across various platforms and languages, making it a flexible solution for code review.
- **Detailed Vulnerability Reports**: Generates reports with detailed descriptions of identified vulnerabilities, proof of concepts, and recommended fixes.
- **Supports Multiple Languages**: Works with PHP, Python, JavaScript, and more.
- **Markdown Report Generation**: Outputs security analysis in Markdown format for easy integration with GitHub and other platforms.

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
- `-config`: Path to the configuration file (default is `config.json`).

### Example:

```bash
go run main.go -file example.php -save /home/reports -config config.json
```

### Configuration File (config.json)

The `config.json` file contains the LLaMA API endpoint. Configure it as shown below:

```json
{
  "llama": {
    "api_url": "http://localhost:8000/generate"
  }
}
```

This endpoint should point to a running instance of the LLaMA model API. Update the `api_url` to match your LLaMA setup.

## Contribution

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-branch-name`.
3. Make your changes and commit them: `git commit -m 'Add new feature'`.
4. Push to the branch: `git push origin feature-branch-name`.
5. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, feel free to reach out to the repository owner.

---
Developed with ❤️ by Mohammed Fathy @Secfathy

Modified by @Moaz-Awad
