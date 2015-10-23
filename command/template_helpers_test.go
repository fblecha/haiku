package command

import (
	//	"fmt"
	//"html/template"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConvertTemplateName(t *testing.T) {
	appDir, err := filepath.Abs("../example")
	if err != nil {
		t.Fatal(err)
	}
	path := "../example/views/dogs/layout.html"
	absPath, _ := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}
	expectedName := "dogs/layout.html"
	actualName := ConvertTemplateName(appDir, absPath)
	if actualName != expectedName {
		t.Fatalf("Expected '%s' but actually got '%s' ", expectedName, actualName)
	}
}

func TestFindPartialTemplates(t *testing.T) {
	wd, _ := os.Getwd()
	os.Chdir("../example")
	expectations := []string{
		"dogs/menu2.partial.html", //dog specific menu
		"menu.partial.html",       //site menu
	}
	dir, err := filepath.Abs("../example")
	if err != nil {
		t.Fatal(err)
	}
	results := FindPartialTemplates(dir)

	for i, expected := range expectations {
		found := false
		for j := range results {
			if strings.ContainsAny(string(expected[i]), results[j]) {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("unable to find %s in %q ", expected, results)
		}
	}
	os.Chdir(wd)
}

// func TestLoadPartialTemplates(t *testing.T) {
// 	//tmpl := template.New("root")
// 	partials = FindPartialTemplates("../example")
// 	fmt.Printf("found partials = %q \n", partials)
//
// 	//appDir, _ := filepath.Abs("../example") //HACK err not used
// 	//tmpl = LoadPartialTemplates(appDir, partials, tmpl)
// 	// var results map[string]bool
// 	//fmt.Printf("len(tmpl.Templates) = %s \n", len(tmpl.Templates()))
// 	// for _, tp := range tmpl.Templates() {
// 	// 	results[tp.Name()] = true
// 	// }
// 	// expectations := []string{
// 	// 	"views/dogs/menu.partial.html", //dog specific menu
// 	// 	"views/menu.partial.html",      //site menu
// 	// }
// 	// for _, expected := range expectations {
// 	// 	if _, ok := results[expected]; ok != true {
// 	// 		t.Fatalf("expected %s in %q but it wasn't there", expected, results)
// 	// 	}
// 	// }
// }