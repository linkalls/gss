package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "regexp"
    "strings"

    "github.com/fsnotify/fsnotify"
)

var cssTemplates = map[string]string{
    "container": `
.container {
  max-width: 100%;
  margin-left: auto;
  margin-right: auto;
  padding-left: 1rem;
  padding-right: 1rem;
}
`,

    "grid": `
.grid {
  display: grid;
  gap: 1rem;
}
`,

    "grid-cols-1": `
.grid-cols-1 {
  grid-template-columns: repeat(1, minmax(0, 1fr));
}
`,

    "grid-cols-2": `
.grid-cols-2 {
  grid-template-columns: repeat(2, minmax(0, 1fr));
}
`,

    "grid-cols-3": `
.grid-cols-3 {
  grid-template-columns: repeat(3, minmax(0, 1fr));
}
`,

    "grid-cols-4": `
.grid-cols-4 {
  grid-template-columns: repeat(4, minmax(0, 1fr));
}
`,

    "text-center": `
.text-center {
  text-align: center;
}
`,

    "text-left": `
.text-left {
  text-align: left;
}
`,

    "text-right": `
.text-right {
  text-align: right;
}
`,

    "font-bold": `
.font-bold {
  font-weight: bold;
}
`,

    "text-xl": `
.text-xl {
  font-size: 1.25rem;
}
`,

    "text-2xl": `
.text-2xl {
  font-size: 1.5rem;
}
`,

    "text-3xl": `
.text-3xl {
  font-size: 1.875rem;
}
`,

    "text-4xl": `
.text-4xl {
  font-size: 2.25rem;
}
`,

    "text-5xl": `
.text-5xl {
  font-size: 3rem;
}
`,

    "btn": `
.btn {
  display: inline-block;
  padding: 0.5rem 1rem;
  border-radius: 0.375rem;
  font-weight: 600;
  text-align: center;
  transition: background-color 0.3s;
}
`,

    "btn-primary": `
.btn-primary {
  background-color: #007bff;
  color: white;
}

.btn-primary:hover {
  background-color: #0056b3;
}
`,

    "btn-secondary": `
.btn-secondary {
  background-color: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background-color: #5a6268;
}
`,

    "container-responsive": `
@media (min-width: 640px) {
  .container {
    max-width: 640px;
  }
}

@media (min-width: 768px) {
  .container {
    max-width: 768px;
  }
}

@media (min-width: 1024px) {
  .container {
    max-width: 1024px;
  }
}

@media (min-width: 1280px) {
  .container {
    max-width: 1280px;
  }
}
`,
}

// CSSファイルを生成
func generateCSS(htmlDir string) {
	cssDir := filepath.Join(htmlDir, "styles")

	// HTMLフォルダ内のHTMLファイルを取得
	files, err := os.ReadDir(htmlDir)
	if err != nil {
			fmt.Println("HTMLフォルダの読み込みエラー:", err)
			return
	}

	// 使用されているクラス名を抽出
	classSet := make(map[string]bool)
	classRegex := regexp.MustCompile(`class="([^"]+)"`)

	for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".html" {
					htmlFilename := filepath.Join(htmlDir, file.Name())

					// HTMLファイルを読み込み
					htmlContent, err := os.ReadFile(htmlFilename)
					if err != nil {
							fmt.Println("HTMLファイルの読み込みエラー:", err)
							return
					}

					// クラス名を抽出してセットに追加
					matches := classRegex.FindAllStringSubmatch(string(htmlContent), -1)
					for _, match := range matches {
							classes := strings.Split(match[1], " ")
							for _, class := range classes {
									classSet[class] = true
							}
					}
			}
	}

	// CSSファイルを書き込む
	for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".html" {
					cssFilename := filepath.Join(cssDir, strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))+".css")
					cssContent := ""

					for class := range classSet {
							if template, exists := cssTemplates[class]; exists {
									cssContent += template
							}
					}

					err := os.WriteFile(cssFilename, []byte(cssContent), 0644)
					if err != nil {
							fmt.Println("CSSファイルの書き込みエラー:", err)
							return
					}
					fmt.Println("CSSファイルが生成されました:", cssFilename)
			}
	}
}


func main() {
    // 引数で指定されたフォルダ内のHTMLファイルを監視
    if len(os.Args) < 2 {
        fmt.Println("使用法: go run generate_css.go <HTMLフォルダ>")
        return
    }
    htmlDir := os.Args[1]

    // CSSディレクトリが存在しない場合は作成
    cssDir := filepath.Join(htmlDir, "styles")
    if _, err := os.Stat(cssDir); os.IsNotExist(err) {
        os.Mkdir(cssDir, os.ModePerm)
    }

    // 初回のCSS生成
    generateCSS(htmlDir)

    // HTMLフォルダ内のHTMLファイルの変更を監視
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        fmt.Println("Watcherの初期化エラー:", err)
        return
    }
    defer watcher.Close()

    done := make(chan bool)
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Op&fsnotify.Write == fsnotify.Write && filepath.Ext(event.Name) == ".html" {
                    fmt.Println("HTMLファイルが変更されました:", event.Name)
                    generateCSS(htmlDir)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                fmt.Println("Watcherエラー:", err)
            }
        }
    }()

    err = watcher.Add(htmlDir)
    if err != nil {
        fmt.Println("Watcherの追加エラー:", err)
        return
    }
    fmt.Println("監視中:", htmlDir)

    // exitコマンドが入力されるまで監視を継続
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        if scanner.Text() == "exit" {
            break
        }
    }

    <-done
}