# GSS: 動的 CSS ジェネレーターとウォッチャー

GSS (Go Style Sheets) は、HTML ファイル内のクラス名に基づいて動的に CSS ファイルを生成および監視するための便利なツールです。 HTML ファイルに変更が加えられると、GSS はそれらの変更を検出し、対応する CSS ファイルを自動的に更新します。

## 特徴
1. **動的な CSS 生成**: HTML ファイル内のクラス名を検出し、それらのクラスに対して CSS ルールを動的に生成します。
2. **監視モード**: HTML ファイルの変更を監視し、変更が検出されたときに CSS ファイルを自動的に更新します。
3. **カスタム CSS テンプレート**: ニーズに合わせてカスタマイズできる一連の CSS テンプレートを提供します。
4. **応答性**: さまざまな画面サイズに対応するレスポンシブ デザイン用のメディア クエリをサポートします。
5. **アニメーションと遷移**: ボタンやその他の要素にアニメーション効果を追加します。
6. **フォーム スタイリング**: 入力フィールドやエラーメッセージなど、フォーム要素のスタイリングを強化します。
7. **グリッド システム**: グリッド レイアウトを使用して、応答性の高いレイアウトを作成します。
8. **カラー スキーム**: プロジェクト全体で使用するカラー パレットを定義します。

## 開始するには
GSS を使用するには、次の手順に従います。

1. **インストール**: GSS は Go モジュールとして配布されます。プロジェクトでまだ行っていない場合は、次のコマンドを実行してインストールします。

   ```shell
   go get github.com/username/gss
   ```

2. **HTML フォルダーを準備する**: HTML ファイルを含むフォルダーを作成します。 GSS はこのフォルダー内の HTML ファイルを処理します。

3. **CSS テンプレートをカスタマイズする**: `css_templates.go` ファイルで提供される CSS テンプレートをカスタマイズして、独自のデザインと要件に合わせて調整できます。

4. **コマンドラインから実行する**: 次のコマンドを使用して GSS を実行します。

   ```shell
   go run main.go <HTML_FOLDER_PATH> [--watch]
   ```

   `<HTML_FOLDER_PATH>` を HTML ファイルを含むフォルダーのパスに置き換えます。オプションの `--watch` フラグを含めると、GSS は HTML ファイルの変更を監視し、自動的に CSS ファイルを更新します。

5. **HTML ファイルを編集する**: HTML ファイルにクラス名を追加または編集します。 GSS はこれらの変更を検出し、対応する CSS ルールを生成します。

## CSS テンプレートと使用法
GSS は、さまざまなデザイン要件に対応する幅広い CSS テンプレートを提供します。

### コンテナー

**クラス名**: `container`

**使用法**: このクラスは、コンテンツを中央に配置し、最大幅を画面幅の 100% に制限するコンテナーを作成します。これは、Web サイトのメイン コンテンツ領域をラップするのに便利です。

```html
<div class="container">
    <!-- コンテンツここに -->
</div>
```

### フレックス

**クラス名**: `flex`

**使用法**: このクラスは、要素をフレックスボックス コンテナーとして表示し、子要素を水平方向に並べて配置します。これは、親要素内で項目を柔軟にレイアウトする場合に便利です。

```html
<div class="flex">
    <div>フレックスアイテム 1</div>
    <div>フレックスアイテム 2</div>
    <div>フレックスアイテム 3</div>
</div>
```

### 青色背景

**クラス名**: `blue`

**使用法**: このクラスは、要素の背景色を青色 (#5980F6) に設定します。このクラスは、背景色をすばやく追加するのに役立ちます。

```html
<div class="blue">青色の背景を持つ要素</div>
```

### グリッド

**クラス名**: `grid`、`grid-cols-1`、`grid-cols-2`、`grid-cols-3`、`grid-cols-4`

**使用法**: これらのクラスは、グリッド レイアウトを作成するために使用されます。 `grid` クラスはグリッド コンテナーを定義し、`grid-cols-N` クラスはグリッド内の列の数を指定します (`N` は 1 から 4 までの数)。

```html
<div class="grid grid-cols-2">
    <div>グリッド アイテム 1</div>
    <div>グリッド アイテム 2</div>
    <div>グリッド アイテム 3</div>
</div>
```

### テキスト配置

**クラス名**: `text-center`、`text-left`、`text-right`

**使用法**: これらのクラスは、テキストの配置を制御します。 `text-center` はテキストを中央揃えに、`text-left` はテキストを左揃えに、`text-right` はテキストを右揃えにします。

```html
<h1 class="text-center">中央揃えの見出し</h1>
<p class="text-left">左揃えの段落</p>
```

### フォントの太さとサイズ

**クラス名**: `font-bold`、`text-xl`、`text-2xl`、`text-3xl`、`text-4xl`、`text-5xl`

**使用法**: これらのクラスは、フォントの太さとサイズを変更します。 `font-bold` はフォントを太字にします。 `text-xl`、`text-2xl` などでは、フォント サイズが増加します ( `text-2xl` は `text-xl` より大きくなります)。

```html
<h1 class="font-bold text-2xl">太字で大きな見出し</h1>
```

### ボタン

**クラス名**: `btn`、`btn-primary`、`btn-secondary`

**使用法**: これらのクラスは、ボタンのスタイルを定義します。 `btn` クラスは基本的なボタンを作成し、`btn-primary` と `btn-secondary` は異なる色のボタンを作成します。

```html
<button class="btn btn-primary">プライマリ ボタン</button>
<button class="btn btn-secondary">セカンダリ ボタン</button>
```

### レスポンシブ デザイン

GSS には、レスポンシブ デザイン用のメディア クエリが組み込まれています。特定の画面サイズに対して異なるスタイルを適用できます。

```css
@media (max-width: 768px) {
    /* 768px 以下の画面幅に対するスタイル */
}

@media (min-width: 768px) and (max-width: 1024px) {
    /* 768px から 1024px の画面幅に対するスタイル */
}
```

### アニメーションと遷移

GSS は、ボタンのような要素にアニメーション効果を追加できます。

```css
.btn {
    /* ボタンのアニメーション */
    transition: background-color 0.3s, transform 0.2s;
}

.btn:hover {
    /* ホバーエフェクト */
    background-color: #4FD1C5;
    transform: translateY(-2px);
}
```

### フォーム スタイリング

GSS は、入力フィールドやエラーメッセージなど、フォーム要素のスタイリングを強化します。

```css
.input-field {
    /* 入力フィールドのスタイル */
    border: 1px solid #D1D5DB;
    border-radius: 0.375rem;
    padding: 0.5rem 0.75rem;
    transition: border-color 0.3s;
}

.input-field:focus {
    /* フォーカス時のスタイル */
    border-color: #4FD1C5;
    outline: none;
    box-shadow: 0 0 0 3px rgba(79, 209, 197, 0.2);
}

.error-message {
    /* エラーメッセージのスタイル */
    color: #EF4444;
    font-size: 0.875rem;
}
```

### グリッド システム

GSS は、グリッド レイアウトを使用して応答性の高いデザインを作成するグリッド システムを提供します。

```css
.container {
    /* コンテナのスタイル */
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 1rem;
}

.column {
    /* カラムのスタイル */
    padding: 1rem;
    background-color: #F3F4F6;
    border-radius: 0.5rem;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}
```

### カラー スキーム

GSS を使用すると、プロジェクト全体で使用するカラー パレットを定義できます。

```css
.primary-color {
    /* プライマリカラー */
    color: #4FD1C5;
}

.secondary-color {
    /* セカンダリカラー */
    color: #6B7280;
}
```

## 貢献
プルリクエストは歓迎です！ GSS をより良いものにするために、バグ修正、機能強化、または追加の CSS テンプレートの貢献を歓迎します。

## ライセンス
このプロジェクトは MIT ライセンスの下でライセンスされています。詳細については、「LICENSE」ファイルを参照してください。