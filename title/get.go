package title

var titles = map[string][]string{
	"Nao": []string{
		`MIDE-832 学生時代は親友の彼女で3人で雑魚寝もしてたただの女友達と大人になって再会してめちゃくちゃ中出ししまくった。 神宮寺ナオ`,
		`JUL-461 人妻オフィスレディの絶対領域 貞淑妻を襲う、部長の言いなり社内羞恥―。 神宮寺ナオ`,
		`MIDE-880 超高級SEXYランジェリー販売員の見せつけ誘惑セールス 神宮寺ナオ`,
		`JUL-388 高級娼婦 神宮寺ナオ 妖艶、華麗な人妻―、美しさ極める。`,
	},
	"Aika": []string{
		`PRED-246 射精まで乳首を責めちゃう年上彼女と乳首こねくり同棲生活 山岸逢花`,
		`PRED-262 綺麗なお姉さんのデカ尻でキ○タマ空っぽになっても痴女られる！ ド迫力。追撃尻痴女フルコース 山岸逢花`,
		`PRED-277 最高すぎた不倫生活。セックスも、日常も、全てでオレをダメにする愛人沼で溶かされて…。 山岸逢花`,
	},
	"Saeko": []string{
		`ADN-219 となり妻 背徳の昼下がり 松下紗栄子`,
		`IPX-461 出張先相部屋NTR 絶倫の部下に一晩中何度も中出しされた巨乳女上司 松下紗栄子`,
		`SSPD-150 夫のいない数日間、夫の部下に抱かれ続けた記録。 松下紗栄子`,
		`ATID-401 欲求不満な人妻は毎晩隣人に3時間抱かれています。 松下紗栄子`,
	},
}

var counter = make(map[string]int)

func getByName(name string) string {
	cnt := counter[name]
	counter[name] = (cnt + 1) % len(titles[name])
	return titles[name][cnt]
}

func Nao() string {
	return getByName("Nao")
}

func Aika() string {
	return getByName("Aika")
}

func Saeko() string {
	return getByName("Saeko")
}
