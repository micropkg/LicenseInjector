package injector

import "github.com/micropkg/gstore"

//Inject License into binary
func Inject(text string) {
	s := gstore.Open("_license")
	s.Append(text)
}

//Get Injected Licenses
func Get(text string) []string {
	s := gstore.Open("_license")
	return s.Get()
}
