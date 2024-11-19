package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/antongoncharik/csv-gen-adobe-stock/internal/handlers"
)

func main() {
	// http.HandleFunc("/", handlers.UploadTemplateHandler)
	// http.HandleFunc("/upload", handlers.UploadFileHandler)
	// http.HandleFunc("/table", handlers.TableTamplateHandler)
	// http.HandleFunc("/download", handlers.DownloadCSVHandler)

	// templatesDir := filepath.Join("templates")
	// handlers.LoadTemplates(templatesDir)

	// log.Println("Starting server at :8080")
	// err := http.ListenAndServe(":8080", nil)
	// log.Fatal(err)

	spinner := []rune{'|', '/', '-', '\\'}
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				for _, r := range spinner {
					fmt.Printf("\rLoading... %c", r)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()

	aa := `Valentine's Day Kiss Amidst a Red Heart Shower
	Nighttime Fireworks Over City and Ferris Wheel
	Gingerbread Men in a Snowy Christmas Scene
	Elegant Christmas Macarons with Festive Frosting
	Festive Green Christmas Cocktail with Candy
	A Plate of Delicious Homemade Gingerbread Men
	Five Cheerful Gingerbread Men on a Gray Surface
	Festive Christmas Donut with Holly and Berries
	Magical Christmas Eve Scene with Santa and Trees
	Classic Candy Cane with a Red Bow on White
	Romantic Wooden Heart Frame with Flowers and Hearts
	Festive Christmas Macarons with Candy Toppers
	Branch of Ripe Apricots on White Background
	Apricot Branch on Light Blue Background
	Bunch of Bananas on Purple Background
	Two Shiny Cherries with Red Drops
	Three Cherries Splashing in Cherry Juice
	Single Cherry Splashing in Cherry Juice
	Two Cherries Splashing in Cherry Juice on Purple
	Single Cherry Splashing in Cherry Juice on White
	Two Cherries Splashing Juice on Pink Background
	Double Cherry Splash on White
	Single Cherry Splash on White Background
	Two Cherries and Red Juice Splatter
	Cherries with Juice Splatter on Peach Background
	Two Cherries with Minimal Splash
	Single Cherry with Red Splatter
	Two Cherries with Juice Splash on Pink Background
	Single Cherry with Pink Splatter
	Two Cherries with Red Juice Splatter
	Two Cherries Colliding with a Juice Splash
	Three Cherries and Cherry Juice Splash on Purple
	Single Cherry Splash on White Background
	Single Cherry Splash Against Pink Background
	Two Cherries and a Cherry Juice Splash
	Two Cherries in Cherry Juice
	Three Cherries with Juice Drops on Pink Background
	Single Banana on Aqua Background
	Chocolate Layer Cake with Five Candles
	Pink Birthday Cake with Two Candles
	Vanilla Cake with Three Lit Candles
	Chocolate Cake with Candles
	Birthday Cake with Rainbow Sprinkles and Candles
	Double Layer Chocolate Cake with Three Candles
	White Cake with "Happy Birthday" Candles
	Simple White Cake with Three Candles
	Pink Birthday Cake with Two Candles and Balloons
	Vanilla Cake with Three Candles and Pink Background
	Layered Cake with Three Candles and Balloons
	White Cake with "Happy Birthday" and Candles
	Festive Birthday Cake with Three Candles and Sprinkles
	Two-Layer Birthday Cake with Three Candles and Cupcake
	White Birthday Cake with Four Candles and Sprinkles
	Birthday Cake with Four Candles on Pink Background
	Birthday Cake with Seven Candles and Sprinkles
	Birthday Cake with Seven Candles on Pink Background
	Birthday Cake with Eight Candles on Purple Background
	Birthday Cake with Seven Candles on Pink Background
	Birthday Cake with Six Candles on Pink Background
	Christmas Wreath with Red Bow and Silver Ornaments
	Classic Christmas Wreath with Red Bow on White
	Christmas Wreath with Red Berries on Orange Background
	Christmas Wreath with Red Berries on Beige Background
	Simple Christmas Wreath on Red Background
	Christmas Wreath with Red and Gold Ornaments
	Christmas Wreath with Red Bow on Pink Background
	Frosted Christmas Wreath with Red Ornaments on Red
	Half Christmas Wreath with Red Bows on Purple
	Christmas Wreath with Two Bows on Yellow Background
	Christmas Wreath with Gold Star on Red Background
	Red and Silver Christmas Wreath with Snowflakes
	Simple Christmas Wreath with Red Berries on Yellow
	Christmas Wreath with Red Berries and Red Bow on Teal
	Partial Christmas Garland with Gold Ornaments on Pink
	Christmas Wreath with Gold Ornaments and Flower on Pink
	Elegant Gold Christmas Wreath on Light Blue Background
	Gold Christmas Wreath Detail on Pink Background
	Christmas Wreath with Gold Ornaments and Bow on Orange
	Gold Christmas Garland on Pink Background
	Gold Christmas Wreath with Poinsettia on Orange
	Gold Christmas Wreath with Red Bow on Red
	Gold Christmas Ornaments and Pine Branches on Pink
	Gold Christmas Wreath with Beige Bow on Pink
	Broken Christmas Ornaments in the Snow
	Broken Christmas Ornaments with Snow
	Green Christmas Ornament in the Snow
	Christmas Toys in a Cardboard Box
	Teddy Bear with Christmas Decorations
	Snowy Teddy Bear with Christmas Decorations
	Snowy Teddy Bear with Holly and Gifts
	Teddy Bear Under the Christmas Tree
	Teddy Bear with Christmas Lights and Gifts
	Teddy Bear in the Snow by a Christmas Tree
	White Teddy Bear by Christmas Tree with Presents
	Christmas Teddy Bear in Winter Wonderland
	Teddy Bear in Santa Hat by Christmas Tree
	Teddy Bear with Christmas Lights and Gifts
	Christmas Teddy Bear with Red Bow
	Teddy Bear Under the Christmas Tree
	Teddy Bear in Santa Hat with Christmas Presents`

	titles := handlers.SplitLines(string(aa))

	const requestLimit = 3
	const interval = time.Minute / requestLimit

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	keywords := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, title := range titles {
		<-ticker.C
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			kwrds := handlers.GetKeywords(t)
			mu.Lock()
			keywords[t] = kwrds
			mu.Unlock()
		}(title)
	}
	wg.Wait()

	done <- true
}
