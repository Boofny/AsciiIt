#! /usr/bin/env bash

prodUrl="https://brandonr-portfolio.pages.dev/homepage"

cat >> README.md <<EOF
# This is my Portfolio 

## Hosted on cloud flare pages at 
```bash
${prodUrl}
```
EOF
