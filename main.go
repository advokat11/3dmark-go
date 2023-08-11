package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

var benchmarks = []struct {
	name string
	def  string
}{
	{"Night raid", "nightraid.3dmdef"},
	{"Time Spy", "timespy.3dmdef"},
	{"Fire Strike", "firestrike.3dmdef"},
	{"Sky Diver", "skydiver.3dmdef"},
	{"Cloud Gate", "cloudgate.3dmdef"},
	{"Ice Storm", "icestorm.3dmdef"},
}

func main() {
	fmt.Println("Внимание! Бенчмаркинг может занять продолжительное время.")
	fmt.Println("Не выключайте программу или компьютер во время тестирования.")

	fmt.Println("\nДоступные бенчмарки:")
	for i, b := range benchmarks {
		fmt.Printf("%d. %s\n", i+1, b.name)
	}

	fmt.Print("\nВыберите бенчмарки для запуска (например, 1 2 3) или нажмите Enter для запуска всех: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	var selected []int
	if text != "" {
		nums := strings.Split(text, " ")
		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err == nil && i > 0 && i <= len(benchmarks) {
				selected = append(selected, i-1)
			}
		}
	}

	if len(selected) == 0 {
		for i := range benchmarks {
			selected = append(selected, i)
		}
	}

	// Progress bar
	p := mpb.New(mpb.WithWidth(64))
	total := len(selected)
	bar := p.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.Name("Прогресс тестирования: "),
			decor.CountersKibiByte("% .2f / % .2f"),
		),
		mpb.AppendDecorators(
			decor.Percentage(),
		),
	)

	// Run selected benchmarks
	for _, idx := range selected {
		b := benchmarks[idx]
		runBenchmark(p, bar, b.name, b.def)
	}

	// Wait for progress bar to finish
	p.Wait()

	fmt.Println("\nТестирование окончено!")
}

// Остальной код без изменений

func runBenchmark(p *mpb.Progress, bar *mpb.Bar, name, def string) {
	cmd := exec.Command("C:\\Program Files\\Futuremark\\3Dmark\\3DMarkCmd.exe",
		"--definition="+def,
		"--systeminfo", "on",
		"--out="+os.Getenv("HOMEPATH")+"\\Desktop\\"+name+".3dmark-result",
		"--trace",
		"--export="+os.Getenv("HOMEPATH")+"\\Desktop\\"+name+".xml")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Ошибка запуска %s %v\n", name, err)
	} else {
		bar.Increment() // Correct way to increment the progress bar
	}
}
