package checkspaces

import (
	"go/ast"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "checkspaces is a checker for spaces between // and directives."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "checkspaces",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func init() {
	Analyzer.Flags.StringVar(&configPath, "configPath", "", "config file path(abs)")
}

type Directive struct{
	Content []string`yaml:"directive"`
}

var configPath string

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func readDirectiveConfig() *Directive {
	var cp string
	// if configPath flag is not set
	if configPath ==""{
		curDir, _ := os.Getwd()
		for !fileExists(curDir+"/checkspaces.yml"){
			// Search up to the root
			if curDir == filepath.Dir(curDir) || curDir == ""{
				// If checkspaces.yml is not found
				return nil
			}
			curDir = filepath.Dir(curDir)
		}
		cp = curDir+"/checkspaces.yml"
	}else{
		cp = configPath
	}

	buf, err := ioutil.ReadFile(cp)
	if err != nil {
		return nil
	}
	directivesFromConfig := Directive{}
	err = yaml.Unmarshal([]byte(buf), &directivesFromConfig)
	if err != nil {
		log.Fatalf("yml parse error:%v", err)
	}
	return &directivesFromConfig
}

func run(pass *analysis.Pass) (interface{}, error) {
	directives := readDirectiveConfig()
	if directives == nil {
		directives = &Directive{}
	}
	directives.Content = append(directives.Content, "go:embed")
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Comment)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Comment:
			for _, directive := range directives.Content{
				if strings.Contains(n.Text, "// "+directive) {
					pass.Reportf(n.Pos(), "There is a space between slash and the directive: %s", directive)
				}
			}
		}
	})

	return nil, nil
}
