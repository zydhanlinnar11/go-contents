# 📁 Embedding File

This project demonstrates different approaches to embed files in Go applications using the `embed` package. It compares various methods for including external files directly into your Go binary.

## 🎯 Purpose

The main goal is to compare different methods of embedding files in Go applications:

- **String embedding** - Embedding text files as strings
- **Byte embedding** - Embedding binary files as byte slices  
- **FS embedding** - Embedding entire directories as filesystems
- **Cross-package embedding** - Embedding files from different packages
- **Traditional file reading** - Reading files at runtime (for comparison)

## 🚀 How to Run

### 1️⃣ String Embedding
```bash
cd cmd/embed_string
go run main.go
```
**Output**: Prints the content of `file-to-import.txt` as a string

### 2️⃣ Byte Embedding  
```bash
cd cmd/embed_byte
go run main.go
```
**Output**: Prints the size of the embedded image file in bytes

### 3️⃣ Filesystem Embedding
```bash
cd cmd/embed_fs
go run main.go
```
**Output**: Starts a web server on port 8080 serving embedded static files

### 4️⃣ Traditional File Reading
```bash
cd cmd/open
go run main.go
```
**Output**: Reads and prints file content at runtime (requires file to exist)

### 5️⃣ Cross-Package Embedding
```bash
cd cmd/outside_folder_embed
go run main.go
```
**Output**: Prints content from a file embedded in a different package

## 🔧 Methodology

### General Workflow

1. **Import the embed package**: Use `import _ "embed"` to import the embed functionality
2. **Define embed directive**: Use `//go:embed` comment to specify which files to embed
3. **Declare variables**: Create variables to hold the embedded content
4. **Access embedded data**: Use the variables in your application logic

### Key Differences

- **String embedding**: Best for text files, provides direct string access
- **Byte embedding**: Ideal for binary files, gives raw byte access
- **FS embedding**: Perfect for serving static web assets, provides filesystem interface
- **Cross-package**: Allows embedding files from different packages for better organization

## 🌐 Results

### Web Server Demo (embed_fs)
When you run the filesystem embedding example, you can:

1. **Start the server**: `go run main.go`
2. **Open your browser**: Navigate to `http://localhost:8080`
3. **View the page**: You'll see a styled "Hello World" page served from embedded files
4. **Explore the structure**: The server serves all files from the embedded `static/` directory

The web server demonstrates how embedded files can be served directly from memory without needing external file dependencies.

## ✅ Conclusion

Go's `embed` package provides an excellent solution for:

- **📦 Single binary deployment** - No external file dependencies
- **⚡ Fast startup** - Files are loaded into memory at compile time
- **🔒 Security** - Files are embedded in the binary, not accessible externally
- **🎯 Flexibility** - Multiple embedding methods for different use cases
- **🌍 Web applications** - Perfect for serving static assets in web servers

The `embed` package is a powerful tool that eliminates the need for external file management while providing excellent performance and deployment simplicity. 