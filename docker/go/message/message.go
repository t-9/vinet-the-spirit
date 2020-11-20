package message

import "fmt"

func GetHello() string {
	return "ごきげんよう^_^"
}

func GetMenu() string {
	return fmt.Sprintf(`いかがなさいますか?
1. マーケット一覧を表示して
2. 板情報を表示して
3. アクセスキーを登録して
0. 特に用はないよ
%s`, GetInputLine())
}

func GetBye() string {
	return "それでは，またお会いしましょう！"
}

func GetWhichBoard() string {
	return "どちらの板情報を表示致しますか?"
}

func GetWrongChoice() string {
	return "選択肢の番号を入力してくださいませ"
}

func GetInputLine() string {
	return "> "
}

func GetAPIKey() string {
	return "API Key"
}

func GetAPISecret() string {
	return "API Secret"
}
