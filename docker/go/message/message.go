package message

import (
	"fmt"

	"vinet/menu"
)

func GetHello() string {
	return "ごきげんよう^_^"
}

func GetMenu() string {
	return fmt.Sprintf(`いかがなさいますか?
%d. マーケット一覧を表示して
%d. 板情報を表示して
%d. アクセスキーを登録して
%d. 資産残高を取得して
%d. 仮想通貨預入履歴を表示して
%d. 入金履歴を表示して
%d. 預入用アドレスを表示して
%d. 新規注文を出して
%d. 特に用はないよ
%s`,
		menu.ShowMarkets,
		menu.ShowBoard,
		menu.RegisterAccessKey,
		menu.ShowBalance,
		menu.ShowCoinIn,
		menu.ShowDeposit,
		menu.ShowAddress,
		menu.SendChildOrder,
		menu.Exit,
		GetInputLine(),
	)
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
