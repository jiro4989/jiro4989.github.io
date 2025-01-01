---
layout: default
title: "Web アクセシビリティの視点で自分の Web アプリを見つめ直す"
date: 2025-01-01 21:43:41 +0900
categories: tech
---

# Web アクセシビリティの視点で自分の Web アプリを見つめ直す

仕事で Web アクセシビリティについて意識する機会が増えたので、
個人で開発してる Web アプリを Web アクセシビリティの視点で見つめ直してみた。

といっても、自分が Web フロントの改修をしているわけではなく、
同じ部署でアクセシビリティ活動が進行しているので意識することが多いだけ。

* Table of contents
{:toc}

## Web アクセシビリティのソース

Web アクセシビリティについては WCAG（Web Content Accessibility Guidelines）のサイトで説明がある。
日本語訳サイトの[Web Content Accessibility Guidelines (WCAG) 2.2 (日本語訳)](https://waic.jp/translations/WCAG22/)があるので、
僕はそのサイトをとりあえず見ている。
また、バージョンが古いが[WCAG 2.1 達成方法集](https://waic.jp/translations/WCAG21/Techniques/)というページがあるので、ここも参考にしている。

項目がめちゃくちゃ多いんで、これら全部目をとおして適用しようとするとキツイ。

## 具体的に何を想定するか

アクセシビリティを意識するってのは、
具体的には以下の問題を抱えた方々の利用を想定するということ。

1. 目が見えない、あるいは見えにくい
1. 難聴
1. 身体を動かしづらい
1. 発話困難
1. 光感受性発作
1. 学習障害
1. 認知障害
1. 上記の要素の組み合わせなどの様々な障害のある人

色盲の方だと色の区別がしづらい。
弱視や老眼では、小さい文字が読みづらい。
腕に障害があってマウス操作が難しければ、キーボードだけで操作をする。
また、色のコントラストが小さいと普通の人でも見づらい、といった感じだろうか。

## とりあえず自分のアプリを見る

これらを踏まえて、自分の Web アプリの[bloodborne-build-simulator](https://jiro4989.github.io/bloodborne-build-simulator/)を
アクセシビリティの視点で確認した。

![改修前のトップページ画像](https://github.com/jiro4989/bloodborne-build-simulator/raw/v1.3.3/docs/toppage.png)

Windows 11 に標準で搭載されているナレーターを起動する。
ナレーターの起動は Win + Ctrl + Enter で起動する。

目が見えない人になったつもりで、目を閉じてサイトを操作する。
マウスは使えないのでキーボードだけで操作する。
基本は TAB キーでフォーカスを移動する。

ナレーターはフォーカスのあたった要素が何なのか読み上げてくれるので、それを頼りに操作しようとした。
操作したところ、以下の問題に気づいた。

1. サイト表示時にどこにもフォーカスがあたっていない
1. それぞれの input 要素に説明がないため、何を操作するものなのか分からない
   1. 例えば体力の値を 10 減算するボタンは「-10」と「ボタン」であることしか分からない
1. ボタンを押した結果、何が変化したのか分からない
1. ボタンを押すとフォーカスが外れるため、押すたびに自分が何にフォーカスがあたっているか分からなくなる

他にも問題はあるはずだが、少なくともこれらの問題に気づいた。

自分が作ったサイトだから勝手がわかってるが、
初めて触る人にはまったく理解できないアプリだろう、と感じた。

## 改修内容

前述の気づきのもと、以下の改修をした。

### input に placeholder を付与

ビルド名はただの text input だが、placeholder が未設定だった。

[HTML属性: placeholder - MDN](https://developer.mozilla.org/ja/docs/Web/HTML/Attributes/placeholder)によると、
label の代わりに placeholder を使うのは NG。
期待するデータの種類のヒントになる単語や短いフレーズが適当とのこと。

これを踏まえて「例：上質ビルド」といったテキスト例を placeholder に付与した。
placeholder はスクリーンリーダーで読み上げてくれるので、フォーカス時に例が読み上げられることを確認した。
これで少なくとも何を記述するべきかが明確になった。

### label を付与できない要素に aria-label を付与

前述のビルド名や、各種ステータスの input の説明をどうするべきかは悩んだ。
いわゆる入力フォームであれば label とセットで書くだけで良い。

```html
<label>
ユーザ名を入力してください。
<input type="text" placeholder="例：田中太郎" />
</label>
```

が、今回のアプリではレイアウトの都合で label を付与するのが難しかった。
そのため代用手段として [aria-label](https://developer.mozilla.org/ja/docs/Web/Accessibility/ARIA/Attributes/aria-label)
を付与した。
aria-label は適切なテキストがなかった場合に使用する属性だ。

> メモ: aria-labelは、DOMにラベルとして参照する適切なテキストがない場合に、
> 対話型要素、または他の ARIA 宣言によって対話するように作られた要素に使用することができます。

前述の「-10」ボタンなどは、最終的に以下のコンポーネントになった。

```tsx
const IncreaseAndDecreaseButton = ({statusName, currentValue, additionalValue, value, max, text, setValue}: {statusName: string, currentValue: number, additionalValue: number, value: number,  max: number, text: string, setValue: Dispatch<SetStateAction<number>>}) => {
  const num = Math.abs(value)
  const op = 0 < value ? '加算' : '減算'
  return (
    <button
      type="button"
      className={buttonClass}
      aria-label={`${statusName}の値を ${num} ${op}します。現在の値は ${currentValue} です`}
      onClick={e => setValueWithValidation(additionalValue + value, setValue, 0, max)}
      >{text}</button>
  )
}
```

例えば「-10」ボタンの場合、フォーカスがあたるとスクリーンリーダーは
「体力の値を 10 減算します。現在の値は 10 です」といった具合に読み上げる。
加算と減算両方のボタンで使用する共通コンポーネントであるため、
操作する値が正負どちらかによって、加算・減算の文章を切り替えるようにした。
これにより、そのボタンを操作すると、何の要素がどう変化するか分かるようになった。

### フォーカスが外れる不具合の修正

このシミュレータは普段マウスでしか操作してなかった。
そのため、キーボードで操作するまで勝手にフォーカスが外れることに気付けなかった。
勝手にフォーカスが外れるのは意図しない振る舞いだったので、明確な実装ミスだった。

不具合の理由は単純で、UI 全体のコンポーネント内で、コンポーネントを定義していたから。

```tsx
export default function Simulator() {
  // ↓ コンポーネント内でコンポーネントを定義している
  const ButtonComponent = () => {
    return (
      <button>てきすと</button>
    )
  }

  return (
    <main>
      ほにゃらら
      <ButtonComponent />
    </main>
  )
}
```

これは全く同じ事例の記事があったため、それを参考に解消した。

* [[React]フォーム入力の度にフォーカスが外れてしまうときに確認すべきこと2選 - Qiita](https://qiita.com/shunexe/items/5d88e255f18280d6941d)

コンポーネントをコンポーネント関数の外に移動することで解消した。

```tsx
const ButtonComponent = () => {
  return (
    <button>てきすと</button>
  )
}

export default function Simulator() {
  return (
    <main>
      ほにゃらら
      <ButtonComponent />
    </main>
  )
}
```

### 共有用リンクに label を付与

もともとは共有用 URL を折りたたんだ a タグにしていた。
が、折りたたみ a タグではなく text input にすれば折りたたみが不要だったので、そちらに変更した。
text input に変更したことで label とセットで input の役割を説明できるようになったため、label を付与した。

## まとめ

普段なんとなくで実装していた自分の Web アプリも、
視点を変えるだけでいろんな課題があったことに気づいた。

今回の改修は、自分が気づいた範囲のみの改修であり、全盲の方にはまだ操作しづらいと思う。
しかしながら、改修前よりはずっと親切な作りになったと思う。

今回のアプリは Bloodborne プレイヤーが主なユーザと想定しているため、
目の不自由な方が使うことはほとんどないだろう。
しかし、Web アクセシビリティの勉強のための題材として、
少しずつ改修を続けようと思う。
