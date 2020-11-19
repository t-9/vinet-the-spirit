package message

func GetHello() string {
	return "ごきげんよう^_^"
}

func GetMenu() string {
	return `いかがなさいますか?
1. マーケット一覧を表示して
2. 板情報を表示して
0. 特に用はないよ
> `
}

func GetBye() string {
	return "それでは，またお会いしましょう！"
}
