#!bin/bash

ENV_FILE=.env.test go test $(go list ./... | grep -v '/tests' | grep -v '/config') -coverprofile=coverage.out && \
if [ $? -ne 0 ]; then
  echo "Tests failed. Exiting."
  exit 1
fi
go tool cover -html=coverage.out -o coverage.html && \
if [ $? -ne 0 ]; then
  echo "Failed to generate coverage report. Exiting."
  exit 1
fi
rm coverage.out

echo "カバレッジレポートをブラウザで表示しますか？ (y/n)"
read -r answer
if [[ "$answer" == "y" || "$answer" == "Y" ]]; then
  if command -v xdg-open &> /dev/null; then
    xdg-open coverage.html
  elif command -v open &> /dev/null; then
    open coverage.html
  else
    echo "ブラウザを開くコマンドが見つかりません。ブラウザで coverage.html を手動で開いてください。"
  fi
else
  echo "カバレッジレポートは coverage.html に保存されました。"
fi
