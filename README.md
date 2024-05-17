# minimum_fts
This is a minimal full text search engine program implemented in Go. It allows the user to understand the simplified structure of the fts.

![スクリーンショット 2024-05-17 9 21 23](https://github.com/shiga0/minimum_fts/assets/13078565/f35a4b4f-d252-40ac-9b79-c018ed6f5f27)

## Dependencies
-  [snowball](https://github.com/kljensen/snowball?tab=readme-ov-file): Stemming language and algorithms.

## Usage

```
minimum_fts % go run main.go
========================
🚀    minimum_fts    🚀
========================
📕 Loaded 10
🛠 Indexed 10
🤓 search_word : deep reason
🔎 Search found 2
    2 It wasn't for any deep reason.
    7 Why did I do such a reckless thing, some might ask. It wasn't for any deep reason.
========================
🎊🥳🎉 Congrats!  🎊🥳🎉
========================
```
