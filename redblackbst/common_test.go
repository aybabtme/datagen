package redblackbst

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/browser"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"
)

func loadWordlist(filename string, fallback []string) (out []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("couldn't open %q: %v", filename, err)
		return fallback
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		out = append(out, scan.Text())
	}

	if scan.Err() != nil {
		log.Printf("scanning %q: %v", filename, scan.Err())
		return fallback
	}

	return
}

var cachekv map[K]string

func makeKV() map[K]string {
	if cachekv != nil {
		return cachekv
	}

	cachekv := map[K]string{}
	for _, word := range loadWordlist(
		"/usr/share/dict/web2",
		append(web2fallback, web2afallback...),
	) {
		cachekv[K("key:"+word)] = "value:" + word
	}
	return cachekv
}

var (
	small = []string{"abasement", "abaser", "Abasgi", "abash", "abashed", "abashedly", "abashedness", "abashless", "abashlessly", "abashment", "abasia", "abasic", "abask", "Abassin", "abastardize", "abatable", "abate", "abatement", "abater", "abatis", "abatised", "abaton", "abator", "abattoir", "Abatua", "abature", "abave", "abaxial", "abaxile", "abaze", "abb", "Abba", "abbacomes", "abbacy", "Abbadide"}

	web2fallback  = []string{"A", "a", "aa", "aal", "aalii", "aam", "Aani", "aardvark", "aardwolf", "Aaron", "Aaronic", "Aaronical", "Aaronite", "Aaronitic", "Aaru", "Ab", "aba", "Ababdeh", "Ababua", "abac", "abaca", "abacate", "abacay", "abacinate", "abacination", "abaciscus", "abacist", "aback", "abactinal", "abactinally", "abaction", "abactor", "abaculus", "abacus", "Abadite", "abaff", "abaft", "abaisance", "abaiser", "abaissed", "abalienate", "abalienation", "abalone", "Abama", "abampere", "abandon", "abandonable", "abandoned", "abandonedly", "abandonee", "abandoner", "abandonment", "Abanic", "Abantes", "abaptiston", "Abarambo", "Abaris", "abarthrosis", "abarticular", "abarticulation", "abas", "abase", "abased", "abasedly", "abasedness", "abasement", "abaser", "Abasgi", "abash", "abashed", "abashedly", "abashedness", "abashless", "abashlessly", "abashment", "abasia", "abasic", "abask", "Abassin", "abastardize", "abatable", "abate", "abatement", "abater", "abatis", "abatised", "abaton", "abator", "abattoir", "Abatua", "abature", "abave", "abaxial", "abaxile", "abaze", "abb", "Abba", "abbacomes", "abbacy", "Abbadide"}
	web2afallback = []string{"A", "acid", "abacus", "major", "abacus", "pythagoricus", "A", "battery", "abbey", "counter", "abbey", "laird", "abbey", "lands", "abbey", "lubber", "abbot", "cloth", "Abbott", "papyrus", "abb", "wool", "A-b-c", "book", "A-b-c", "method", "abdomino-uterotomy", "Abdul-baha", "a-be", "aberrant", "duct", "aberration", "constant", "abiding", "place", "able-bodied", "able-bodiedness", "able-minded", "able-mindedness", "able", "seaman", "aboli", "fruit", "A", "bond", "Abor-miri", "a-borning", "about-face", "about", "ship", "about-sledge", "above-cited", "above-found", "above-given", "above-mentioned", "above-named", "above-quoted", "above-reported", "above-said", "above-water", "above-written", "Abraham-man", "abraum", "salts", "abraxas", "stone", "Abri", "audit", "culture", "abruptly", "acuminate", "abruptly", "pinnate", "absciss", "layer", "absence", "state", "absentee", "voting", "absent-minded", "absent-mindedly", "absent-mindedness", "absent", "treatment", "absent", "voter", "Absent", "voting", "absinthe", "green", "absinthe", "oil", "absorption", "bands", "absorption", "circuit", "absorption", "coefficient", "absorption", "current", "absorption", "dynamometer", "absorption", "factor", "absorption", "lines", "absorption", "pipette", "absorption", "screen", "absorption", "spectrum", "absorption", "system", "A", "b", "station", "abstinence", "theory", "abstract", "group", "Abt", "system", "abundance", "declaree", "aburachan", "seed", "abutment", "arch", "abutment", "pier", "abutting", "joint", "acacia", "veld", "academy", "blue", "academy", "board", "academy", "figure", "acajou", "balsam", "acanthosis", "nigricans", "acanthus", "family", "acanthus", "leaf", "acaroid", "resin", "Acca", "larentia", "acceleration", "note", "accelerator", "nerve", "accent", "mark", "acceptance", "bill", "acceptance", "house", "acceptance", "supra", "protest", "acceptor", "supra", "protest", "accession", "book", "accession", "number", "accession", "service", "access", "road", "accident", "insurance"}
)

// helps print stats about a tree

type nodeSpec struct {
	HasLeft  bool
	HasRight bool
	IsRed    bool
}

func (n nodeSpec) String() string {
	if n.IsRed {
		return fmt.Sprintf("Red(%v, %v)", n.HasLeft, n.HasRight)
	}
	return fmt.Sprintf("Black(%v, %v)", n.HasLeft, n.HasRight)
}

type nodeSpecs []nodeSpec

func (n nodeSpecs) Len() int      { return len(n) }
func (n nodeSpecs) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n nodeSpecs) Less(i, j int) bool {
	in, jn := n[i], n[j]
	if in.IsRed && !jn.IsRed {
		return true
	}
	if jn.IsRed && !in.IsRed {
		return false
	}
	if in.HasLeft && !jn.HasLeft {
		return true
	}
	if jn.HasLeft && !in.HasLeft {
		return false
	}
	if in.HasRight && !jn.HasRight {
		return true
	}
	if jn.HasRight && !in.HasRight {
		return false
	}
	return false
}

func openDot(r io.Reader) {
	cmd := exec.Command("dot", "-Tsvg")
	cmd.Stdin = r
	data, err := cmd.CombinedOutput()
	if err != nil {
		panic(string(data))
	}
	svg := bytes.NewBuffer(data)
	browser.OpenReader(svg)
}

func printTreeStats(tree *RedBlack, filename string) {

	dot := bytes.NewBuffer(nil)
	dotGraph(tree.root, dot, filename)
	openDot(dot)

	specmap := map[nodeSpec]int{}

	visit := func(x *node) bool {

		s := nodeSpec{
			HasLeft:  x.left == nil,
			HasRight: x.right == nil,
			IsRed:    isRed(x),
		}

		specmap[s] = specmap[s] + 1

		return true
	}

	min, _, _ := tree.Min()
	max, _, _ := tree.Max()
	nodes(tree.root, visit, min, max)

	var sorted nodeSpecs
	for k := range specmap {
		sorted = append(sorted, k)
	}
	sort.Sort(sorted)

	for _, spec := range sorted {
		fmt.Printf("%d x %v\n", specmap[spec], spec)
	}
}
