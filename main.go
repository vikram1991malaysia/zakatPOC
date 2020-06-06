package main

import (
	"log"
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"html/template"
	"strconv"
	"math"
	"encoding/json"
)

//Global variables
	var tpl *template.Template
	
	func init(){
	
		tpl = template.Must(template.New("main").ParseGlob("assets/*.html"))
	}

	func main(){

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)

		

		http.HandleFunc("/", index)
		http.HandleFunc("/zakaat", zakaat)
		http.HandleFunc("/hajj", hajj)

		http.HandleFunc("/css/", serveResource)
    	http.HandleFunc("/vendor/", serveResource)
    	http.HandleFunc("/icons/", serveResource)
    	http.HandleFunc("/js/", serveResource)
    	http.HandleFunc("/fonts/", serveResource)
    	http.HandleFunc("/images/", serveResource)

		http.ListenAndServe(":8090", nil)
		
	}

	func serveResource(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Connected")
	    path := "assets"+r.URL.Path
	    // fmt.Println(path)
	    var contentType string
	    if strings.HasSuffix(path, ".css") {
	        contentType = "text/css"
	    } else if strings.HasSuffix(path, ".js") {
	        contentType = "text/javascript"
	    } else if strings.HasSuffix(path, ".png") {
	        contentType = "image/png"
	    } else {
	        contentType = "text/plain"
	    }

	    f, err := os.Open(path)

	    if err == nil {
	        defer f.Close()
	        w.Header().Add("Content-Type", contentType)

	        br := bufio.NewReader(f)
	        br.WriteTo(w)
	    } else {
	        w.WriteHeader(404)
	    }
	}
	func index(w http.ResponseWriter, r *http.Request) {
		
		// w.Header().Set("Content-Type", "text/html")
		tpl.ExecuteTemplate(w, "index.html", nil)
         
	}

	func hajj(w http.ResponseWriter, r *http.Request){
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		return
		} 

		var (
			amt string
			asset string
		)

		amt = r.FormValue("amt")
		asset = r.FormValue("asset")
		
		fmt.Println("amount = "+amt)
		fmt.Println("asset = "+asset)

		fmt.Println(price(asset))

		balance, err := strconv.ParseFloat(amt, 64)
		asset_price := price(asset)

		if err != nil {
			log.Fatal(err)
		}

		final := balance * asset_price

		final = math.Round(final*10)/10

		//final balance in USD
		fmt.Println("final balance in USD")
		fmt.Println(final)
		
		// USD
		usd := final //USD value of crypto

		var nisaab_in_USD float64 = 14000

		hajj_in_USD := usd - nisaab_in_USD


		mapD := map[string]float64{"hajj_in_USD": hajj_in_USD}
    	mapB, _ := json.Marshal(mapD)
    	fmt.Println(string(mapB))
		fmt.Fprint(w, string(mapB))
	}


	func zakaat(w http.ResponseWriter, r *http.Request){
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		return
		} 

		var (
			amt string
			asset string
		)

		amt = r.FormValue("amt")
		asset = r.FormValue("asset")
		
		fmt.Println("amount = "+amt)
		fmt.Println("asset = "+asset)

		fmt.Println(price(asset))

		balance, err := strconv.ParseFloat(amt, 64)
		asset_price := price(asset)

		if err != nil {
			log.Fatal(err)
		}

		final := balance * asset_price

		

		final = math.Round(final*10)/10

		//final balance in USD
		fmt.Println("final balance in USD")
		fmt.Println(final)
		

		// USD
		usd := final //USD value of crypto

		nisaab_in_USD := 75 * 56.35

		OnZakaat_USD := usd - nisaab_in_USD

		zakaat_in_USD := OnZakaat_USD * 0.025
		fmt.Println("zakaat in USD")
		fmt.Println(math.Round(zakaat_in_USD*1000)/1000)


		//Bitcoin
		zakaat_in_bitcoin := zakaat_in_USD/asset_price
		fmt.Println("zakaat in Bitcoin")
		fmt.Println(math.Round(zakaat_in_bitcoin*1000)/1000)
		// fmt.Fprint(w, " ", zakaat_in_bitcoin)
		//Gold
		zakaat_in_gold := zakaat_in_USD/56.35
		fmt.Println("zakaat in Gold")
		fmt.Println(math.Round(zakaat_in_gold*1000)/1000)
		// zakaat_in_gold = math.Round(zakaat_in_gold*1000)/1000
		// fmt.Fprint(w, " ",zakaat_in_gold)


		// Malaysian Ringgit
		rm := final * 4.30 // 1 usd to rm value is 4.30

		nisaab_rm := 75 * 240.15 // 75 gms * per 1 gram of gold in RM

		OnZakaat_rm := rm - nisaab_rm // The value on which zakaat comes

		zakaatinRinggit := OnZakaat_rm * 0.025 // zakaat in RM

		// fmt.Fprint(w, zakaat)
		fmt.Println("total crypto value in RM")
		fmt.Println(rm)
		fmt.Println("nisaab in RM")
		fmt.Println(nisaab_rm)
		fmt.Println("on zakaat")
		fmt.Println(OnZakaat_rm)
		fmt.Println("zakaat in RM")
		fmt.Println(math.Round(zakaatinRinggit*1000)/1000)


		type detials struct {
			Asset string `json:"asset"`
			Zakaat_in_USD float64 `json:"zakaat_in_USD"`
			Zakaat_in_gold float64 `json:"zakaat_in_gold"`
			Zakaat_in_bitcoin float64 `json:"zakaat_in_bitcoin"`
			Zakaat_in_rm float64 `json:"zakaat_in_rm"`
		}
		mapD := &detials {
			Asset:asset, 
			Zakaat_in_USD: zakaat_in_USD, 
			Zakaat_in_gold: zakaat_in_gold, 
			Zakaat_in_bitcoin: zakaat_in_bitcoin, 
			Zakaat_in_rm: zakaatinRinggit,
		}
    	mapB, _ := json.Marshal(mapD)
    	fmt.Println(string(mapB))
		fmt.Fprint(w, string(mapB))
	}