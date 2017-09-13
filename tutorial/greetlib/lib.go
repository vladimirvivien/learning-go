package greetlib

var greetings = map[string]string{
	"English": "Hello, World!",
	"French":  "Salut Monde",
	"Chinese": "世界您好",
	"Klingon": "qo' vIvan",
	"Hindi":   "हैलो वर्ल्ड",
	"Korean":  "안녕하세요",
	"Russian": "привет мир",
	"Swahili": "Wapendwa Dunia",
	"Spanish": "Hola Mundo",
	"Turkish": "Merhaba Dünya",
}

// GreetIn returns a greeting in specified lang
func GreetIn(lang string) string {
	if greeting, ok := greetings[lang]; ok {
		return greeting
	}
	return greetings["English"]
}
