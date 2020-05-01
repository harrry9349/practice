package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"reflect"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	start := time.Now()

	//サンプル量
	n := 750
	//中心
	x1, y1 := 50.0, 50.0

	// 点の生成
	DATA := [][]float64{}
	DATA = make([][]float64, n)
	for i := range DATA {
		DATA[i] = []float64{random(x1), random(y1)}
	}

	// k-means法
	km := Kmeans{}
	km.fit(DATA, 3)

	fmt.Println(km.data)
	fmt.Println(km.labels)
	fmt.Println(km.representatives)

	// 図の生成
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// 図のラベル
	p.Title.Text = "K-Means"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// 補助線
	p.Add(plotter.NewGrid())

	// 散布図の作成
	plot1, err := plotter.NewScatter(Points(0, n, &km))
	if err != nil {
		panic(err)
	}

	plot2, err := plotter.NewScatter(Points(1, n, &km))
	if err != nil {
		panic(err)
	}

	plot3, err := plotter.NewScatter(Points(2, n, &km))
	if err != nil {
		panic(err)
	}

	plot1.GlyphStyle.Color = color.RGBA{R: 255, B: 55, A: 0}
	plot2.GlyphStyle.Color = color.RGBA{R: 0, B: 255, A: 55}
	plot3.GlyphStyle.Color = color.RGBA{R: 55, B: 0, A: 255}

	//くプロットを図に追加
	p.Add(plot1)
	p.Add(plot2)
	p.Add(plot3)

	//ラベル
	p.Legend.Add("class1", plot1)
	p.Legend.Add("class2", plot2)
	p.Legend.Add("class3", plot3)

	// 座標範囲
	p.X.Min = 0
	p.X.Max = 100
	p.Y.Min = 0
	p.Y.Max = 100

	// result.pngに保存
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "result.png"); err != nil {
		panic(err)
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

type Kmeans struct {
	data            [][]float64 // 座標
	labels          []int       // 所属クラス
	representatives [][]float64 // クラス重心
}

//データの転置（各座標のx座標値とy座標値のスライスを作成する）
func Transpose(source [][]float64) [][]float64 {
	var dest [][]float64
	for i := 0; i < len(source[0]); i++ {
		var temp []float64
		for j := 0; j < len(source); j++ {
			temp = append(temp, 0.0)
		}
		dest = append(dest, temp)
	}

	for i := 0; i < len(source); i++ {
		for j := 0; j < len(source[0]); j++ {
			dest[j][i] = source[i][j]
		}
	}
	return dest
}

//スライスに格納されている値が最も小さくなるときにそのインデックスを返す
func ArgMin(target []float64) int {
	var (
		index int
		base  float64
	)
	for i, d := range target {
		if i == 0 {
			index = i
			base = d
		} else {
			if d < base {
				index = i
				base = d
			}
		}

	}
	return index
}

// k-means法
func (km *Kmeans) fit(X [][]float64, k int) []int {

	//乱数のタネ
	rand.Seed(time.Now().UnixNano())
	//データを構造体に格納
	km.data = X

	//乱数で与える初期値の範囲を定める
	//データを座標値ごとに走査して、各座標値の最小値と最大値を獲得する
	transposedData := Transpose(km.data)
	var minMax [][]float64
	for _, d := range transposedData {
		var (
			min float64
			max float64
		)
		for _, n := range d {
			min = math.Min(min, n)
			max = math.Max(max, n)
		}
		minMax = append(minMax, []float64{min, max})
	}

	//各座標値の最小値から最大値までの範囲に収まる乱数を生成し初期値とする
	for i := 0; i < k; i++ {
		km.representatives = append(km.representatives, make([]float64, len(minMax)))
	}
	for i := 0; i < k; i++ {
		for j := 0; j < len(minMax); j++ {
			km.representatives[i][j] = rand.Float64()*(minMax[j][1]-minMax[j][0]) + minMax[j][0]
		}
	}

	//ラベルの初期化
	//各データ点と代表座標との距離を計算
	//最も近い代表座標のインデックスをデータ点のラベルとする
	for _, d := range km.data {
		var distance []float64
		for _, r := range km.representatives {
			distance = append(distance, Dist(d, r))
		}
		km.labels = append(km.labels, ArgMin(distance))
	}

	for {
		//代表座標の更新
		//ラベルに属するデータ点の平均値を代表座標として更新する
		var tempRep [][]float64
		for i, _ := range km.representatives {
			var grouped [][]float64
			for j, d := range km.data {
				if km.labels[j] == i {
					grouped = append(grouped, d)
				}
			}
			if len(grouped) != 0 {

				transposedGroup := Transpose(grouped)
				updated := []float64{}
				for _, vectors := range transposedGroup {

					value := 0.0
					for _, v := range vectors {
						value += v
					}
					//座標値の平均
					updated = append(updated, value/float64(len(vectors)))
				}
				tempRep = append(tempRep, updated)
			}
		}
		km.representatives = tempRep

		//ラベル更新
		//データのラベルが更新により変化しなかった場合に終了
		tempLabel := []int{}
		for _, d := range km.data {
			var distance []float64
			for _, r := range km.representatives {
				distance = append(distance, Dist(d, r))
			}
			tempLabel = append(tempLabel, ArgMin(distance))
		}
		if reflect.DeepEqual(km.labels, tempLabel) {
			break
		} else {
			km.labels = tempLabel
		}
	}
	return km.labels
}

//二点のユークリッド距離
func Dist(source, dest []float64) float64 {
	var dist float64
	for i, _ := range source {
		dist += math.Pow(source[i]-dest[i], 2)
	}
	return math.Sqrt(dist)
}

//ガウス分布
func random(axis float64) float64 {
	//分散
	dispersion := 20.0
	rand.Seed(time.Now().UnixNano())
	// 1ミリ秒間隔で乱数を発生させているため1ミリ秒のウエイトを入れる
	time.Sleep(time.Millisecond * 1)
	return rand.NormFloat64()*dispersion + axis
}

//点の生成
func Points(index int, n int, km *Kmeans) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i, _ := range pts {
		if km.labels[i] == index {
			pts[i].X = km.data[i][0]
			pts[i].Y = km.data[i][1]
		}
	}
	return pts
}
