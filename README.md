## Golang WASM And Web UI Form

## Why WebAssembly?

WebAssembly was created to solve several fundamental challenges in web development:

1. **Performance Limitations**: JavaScript, while versatile, has performance limitations for computationally intensive tasks. WebAssembly provides near-native performance for web applications.

2. **Language Diversity**: Before WebAssembly, JavaScript was the only language that could run natively in browsers. WebAssembly allows developers to use various programming languages for web development.

3. **Code Reuse**: WebAssembly enables the reuse of existing codebases written in languages like C, C++, or Rust in web applications, without the need for complete rewrites in JavaScript.

4. **Security**: WebAssembly runs in a sandboxed environment, providing an additional layer of security compared to native plugins or extensions.

5. **Predictable Performance**: Unlike JavaScript, which is interpreted or JIT-compiled, WebAssembly is a low-level language that can be executed consistently across different browsers and devices.

6. **Efficient Binary Format**: WebAssembly's binary format is designed to be compact and efficiently parsed, leading to faster loading times compared to equivalent JavaScript.

7. **Complementing JavaScript**: WebAssembly is designed to work alongside JavaScript, not replace it, allowing developers to use the best tool for each part of their application.

### WebAssembly with Go

This project demonstrates the use of WebAssembly (Wasm) with Go, showcasing a web application that runs Go code directly in the browser.

## Advantages of Using WebAssembly with Go

1. **Performance**: WebAssembly runs at near-native speed, making it possible to run computationally intensive tasks in the browser with performance comparable to native applications.

2. **Code Reuse**: You can reuse existing Go code in web applications, reducing the need to rewrite logic in JavaScript.

3. **Type Safety**: Go's strong type system helps catch errors at compile-time, leading to more robust web applications.

4. **Concurrency**: Go's excellent concurrency model with goroutines and channels can be leveraged in browser-based applications.

5. **Rich Ecosystem**: Access Go's extensive standard library and third-party packages in your web applications.

6. **Security**: WebAssembly runs in a sandboxed environment, providing an additional layer of security.

7. **Cross-Platform Compatibility**: WebAssembly is supported by all major browsers, ensuring wide compatibility.

8. **Smaller Payload**: Compared to JavaScript, WebAssembly binaries are often smaller, leading to faster load times.

9. **Language Interoperability**: WebAssembly allows Go to interact with JavaScript and the DOM, enabling the creation of full-featured web applications.

10. **Offline Capabilities**: WebAssembly modules can be cached and run offline, improving application availability.

### Prerequisites

- Go 1.11 or later
- A modern web browser that supports WebAssembly

## Project Structure

- `main.go`: The main Go file containing both server-side and client-side code.
- `web/`: Directory containing static files and the WebAssembly binary.
  - `app.wasm`: Compiled WebAssembly binary.
  - `wasm_exec.js`: JavaScript support file for WebAssembly.
  - `index.html`: HTML entry point for the web application.

> Documentation Below Screenshot

![My Image](./webui.jpg)

> Please See The Contents Of >
- `main.go` (Front End)
- `server.go` (Backend Server)
- `index.html` (Web Form)

```bash
go get -u github.com/maxence-charriere/go-app/v10/pkg/app
```

> Current Working Directory >

```bash
~/git/goworkspace/src/wasm_v1
```

```bash
cat $GOROOT/misc/wasm/wasm_exec.js > ~/git/goworkspace/src/wasm_v1/web/wasm_exec.js
```

```bash
go mod init && go mod tidy
```

> Run The Server `server.go`

```bash
go build -o server server.go
```

```bash
./server
Server is running on :8080
```

```bash
 ls -l
total 13936
-rw-r--r--  1 user1  staff      382 Oct 17 09:59 README.md
-rw-r--r--@ 1 user1  staff      186 Oct 17 08:31 go.mod
-rw-r--r--@ 1 user1  staff     1448 Oct 17 08:29 go.sum
-rw-r--r--  1 user1  staff     7442 Oct 17 09:54 main.go
-rwxr-xr-x@ 1 user1  staff  7107634 Oct 17 08:58 server --------> Binary
-rw-r--r--  1 user1  staff     1871 Oct 17 08:58 server.go
drwxr-xr-x  5 user1  staff      160 Oct 17 09:54 web -----------> Directory
```

```bash
ls -l web
total 33776
-rwxr-xr-x@ 1 user1  staff  17265271 Oct 17 09:54 app.wasm -----> Generated
-rw-r--r--  1 user1  staff      1091 Oct 17 09:48 index.html
-rw-r--r--  1 user1  staff     16687 Oct 17 08:30 wasm_exec.js
```

> Build Front-End UI `main.go`

```bash
GOARCH=wasm GOOS=js go build -o web/app.wasm main.go
```

```bash
go run main.go
Server running on http://localhost:8000
```

> Access The Web UI

```
http://localhost:8000/
```