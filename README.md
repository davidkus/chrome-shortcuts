# Chrome Shortcuts

Shortcuts for Chrome address bar - e.g. "gh" -> "https://github.com"

## How it works

1. Copy `config.example.json` to `config.json`
2. Define shortcuts in `config.json`
3. Add "http://localhost:59438/search?q=%s" as your default search provider in Chrome
4. `make build deploy load`
