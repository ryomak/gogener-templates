package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/briandowns/spinner"
	deep "github.com/patrikeh/go-deep"
	"github.com/patrikeh/go-deep/training"
	util "github.com/ryomak/go-deep-util"
	iclass "github.com/ryomak/go-deep-util/iclassifier"
)

func init(){
  appName := "[[.AppName]]"
  fmt.Println(appName,"start")
}

func main() {

	loading := spinner.New(spinner.CharSets[9], 100*time.Millisecond)

	labels := []string{
		"tida",
		"yuna",
		"lulu",
	}

	var i util.IBrainUtil

	i = iclass.Init(
		labels,
		"dataset",
		30,
		30,
	)
	rand.Seed(time.Now().UnixNano())

	data, err := i.MakePattern()
	if err != nil {
		panic(err)
	}

	if len(data) == 0 {
		fmt.Println("no data")
		os.Exit(1)
	}

	//shuffle
	ex := training.Examples(util.DatsetToDataSets(data))
	ex.Shuffle()

	neural := deep.NewNeural(&deep.Config{
		Inputs:     len(data[0].Input),
		Layout:     append([]int{1000, 100}, len(data[0].Response)),
		Activation: deep.ActivationSoftmax,
		Mode:       deep.ModeMultiClass,
		Weight:     deep.NewNormal(1, 0),
		Bias:       true,
	})

	fmt.Printf("train start[testcase:%d] ...\n", len(data))
	loading.Start()
	trainer := training.NewBatchTrainer(training.NewAdam(0.001, 0, 0, 0), 40, len(ex)/2, 12)
	training, heldout := ex.Split(0.8)
	trainer.Train(neural, training, heldout, 60)
	loading.Stop()

	inputFile := "input.jpg"
	gazou, _ := i.Decode(inputFile)
	out, _ := i.Encode(neural.Predict(gazou))

	fmt.Printf("[Result]\n%s is maybe : %v \n", inputFile, out)

	doTest(neural, ex, i)
}

func doTest(neural *deep.Neural, ex training.Examples, i util.IBrainUtil) {

	fmt.Println("Test start with learned Model")
	sum := float64(len(ex))
	correct := 0.0

	for _, p := range ex {
		actual, _ := i.Encode(neural.Predict(p.Input))
		except, _ := i.Encode(p.Response)
		if actual == except {
			correct++
		} else {
			fmt.Printf("miss:except: %s,but actual: %s\n", except, actual)
		}
	}
	fmt.Printf("[Test Result]\ncorrect:%v, sum:%v  %0.1f％\n", correct, sum, 100*correct/sum)
}
