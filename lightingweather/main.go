 package main

 func main() {
	w, err := owm.NewCurrent("F","EN",os.Getenv("OWM_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByZip(os.Getenv("ZIP_CODE"), os.Getenv("COUNTRY_CODE"))
	fmt.Println(w.Main.Temp)
 }