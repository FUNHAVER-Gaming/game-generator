package service

const (
	JoviPcUserId = "910342510600151050"
)

var (
	isProduction = false
)

func init() {
	//file, err := os.Open("../../dev.env")
	//if err != nil {
	//	file, err = os.Open("../../prod.env")
	//	if err != nil {
	//		fmt.Println("Failed to find any .env file, failing")
	//		os.Exit(-1)
	//		return
	//	}
	//	fmt.Println("Did not find dev.env, using production server values")
	//	isProduction = true
	//}
	//
	//if file == nil {
	//	fmt.Println("No .env file found, failing")
	//	os.Exit(-1)
	//	return
	//}
	//
	//err = godotenv.Load(file.Name())
	//
	//if file == nil {
	//	fmt.Println(fmt.Sprintf("Failed to load env file %v", err.Error()))
	//	os.Exit(-1)
	//	return
	//}

}
