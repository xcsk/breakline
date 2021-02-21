package title

var titles = map[string][]string{
	"Nao": []string{
		`MIDE-832 学生時代は親友の彼女で3人で雑魚寝もしてたただの女友達と大人になって再会してめちゃくちゃ中出ししまくった。 神宮寺ナオ`,
		`JUL-461 人妻オフィスレディの絶対領域 貞淑妻を襲う、部長の言いなり社内羞恥―。 神宮寺ナオ`,
		`MIDE-880 超高級SEXYランジェリー販売員の見せつけ誘惑セールス 神宮寺ナオ`,
		`JUL-388 高級娼婦 神宮寺ナオ 妖艶、華麗な人妻―、美しさ極める。`,
	},
}

var counter = make(map[string]int)

func Nao() string {
	name := "Nao"
	cnt := counter[name]
	counter[name] = (cnt + 1) % len(titles[name])
	return titles[name][cnt]
}
