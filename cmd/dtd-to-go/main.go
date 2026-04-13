// dtd-to-go parses cXML DTD files and outputs draft Go struct definitions.
//
// Usage:
//
//	go run ./cmd/dtd-to-go <dtd-file> [files...]
//	go run ./cmd/dtd-to-go --list <dtd-file>
//
// The generated code requires manual review for idiomatic Go naming,
// pointer vs value decisions, and encoding/xml edge cases.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// --- Types ---

// AttrDecl represents a parsed DTD attribute declaration.
type AttrDecl struct {
	Name    string
	Type    string // "CDATA", "(val1|val2)", "NMTOKEN", etc.
	Default string // "#REQUIRED", "#IMPLIED", "\"default\"", "#FIXED \"val\""
}

// ChildRef represents a child element reference in a content model.
type ChildRef struct {
	Name        string
	Cardinality string // "", "?", "*", "+"
	InChoice    bool
}

// ElementInfo represents a fully parsed DTD element with its attributes.
type ElementInfo struct {
	Name     string
	Kind     string // EMPTY, ANY, PCDATA, MIXED, CHILDREN
	RawModel string
	Children []ChildRef
	Attrs    []AttrDecl
}

// --- Regexps ---

var (
	reComment = regexp.MustCompile(`(?s)<!--.*?-->`)
	reEntity  = regexp.MustCompile(`<!ENTITY\s+%\s+(\S+)\s+"([^"]*)"[^>]*>`)
	reElement = regexp.MustCompile(`(?s)<!ELEMENT\s+(\S+)\s+(.*?)>`)
	reAttlist = regexp.MustCompile(`(?s)<!ATTLIST\s+(\S+)\s+(.*?)>`)
	reEntRef  = regexp.MustCompile(`%(\w+);`)
	reChild   = regexp.MustCompile(`([A-Z][A-Za-z0-9]*)([?*+]?)`)
)

// --- Main ---

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: dtd-to-go [--list|--validate-all] <dtd-file> [files...]\n")
		os.Exit(1)
	}

	listOnly := false
	validateAll := false
	var files []string
	for _, a := range os.Args[1:] {
		if a == "--list" {
			listOnly = true
		} else if a == "--validate-all" {
			validateAll = true
		} else {
			files = append(files, a)
		}
	}
	if listOnly && validateAll {
		fmt.Fprintf(os.Stderr, "Error: --list and --validate-all are mutually exclusive\n")
		os.Exit(1)
	}

	if validateAll {
		if len(files) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: dtd-to-go --validate-all <dtd-file> <go-file>\n")
			os.Exit(1)
		}
		ok, err := validateAllStructs(files[0], files[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if !ok {
			os.Exit(2)
		}
		return
	}

	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no DTD files specified\n")
		os.Exit(1)
	}

	for _, f := range files {
		if listOnly {
			listElements(f)
		} else if err := generateStructs(f); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}

type structModel struct {
	StructName string
	Element    string
	Attrs      map[string]bool
	Children   map[string]bool
}

func validateAllStructs(dtdPath, goPath string) (bool, error) {
	data, err := os.ReadFile(dtdPath)
	if err != nil {
		return false, err
	}
	src := reComment.ReplaceAllString(string(data), "")
	ents := parseEntities(src)

	elems := map[string]*ElementInfo{}
	for _, m := range reElement.FindAllStringSubmatch(src, -1) {
		name := m[1]
		raw := strings.TrimSpace(expand(m[2], ents))
		info := &ElementInfo{Name: name, RawModel: raw}
		classify(info)
		elems[name] = info
	}
	for _, m := range reAttlist.FindAllStringSubmatch(src, -1) {
		name := m[1]
		body := expand(strings.TrimSpace(m[2]), ents)
		if info, ok := elems[name]; ok {
			info.Attrs = append(info.Attrs, parseAttrs(body)...)
		}
	}

	models, err := parseStructModels(goPath)
	if err != nil {
		return false, err
	}

	validated := 0
	failures := 0
	for _, sm := range models {
		info, ok := elems[sm.Element]
		if !ok {
			continue
		}
		validated++

		expectedAttrs := map[string]bool{}
		for _, a := range info.Attrs {
			expectedAttrs[a.Name] = true
		}
		expectedChildren := map[string]bool{}
		for _, c := range info.Children {
			expectedChildren[c.Name] = true
		}

		missingAttrs := diffSet(expectedAttrs, sm.Attrs)
		extraAttrs := diffSet(sm.Attrs, expectedAttrs)
		missingChildren := diffSet(expectedChildren, sm.Children)
		extraChildren := diffSet(sm.Children, expectedChildren)

		if len(missingAttrs) == 0 && len(extraAttrs) == 0 && len(missingChildren) == 0 && len(extraChildren) == 0 {
			fmt.Printf("OK  %s (%s)\n", sm.StructName, sm.Element)
			continue
		}

		failures++
		fmt.Printf("DIFF %s (%s)\n", sm.StructName, sm.Element)
		if len(missingAttrs) > 0 {
			fmt.Printf("  missing attrs: %s\n", strings.Join(missingAttrs, ", "))
		}
		if len(extraAttrs) > 0 {
			fmt.Printf("  extra attrs:   %s\n", strings.Join(extraAttrs, ", "))
		}
		if len(missingChildren) > 0 {
			fmt.Printf("  missing elems: %s\n", strings.Join(missingChildren, ", "))
		}
		if len(extraChildren) > 0 {
			fmt.Printf("  extra elems:   %s\n", strings.Join(extraChildren, ", "))
		}
	}

	if validated == 0 {
		fmt.Printf("No matching structs from %s were found in %s\n", filepath.Base(dtdPath), filepath.Base(goPath))
		return false, nil
	}
	fmt.Printf("Validated %d structs against %s\n", validated, filepath.Base(dtdPath))
	return failures == 0, nil
}

func parseStructModels(goPath string) ([]structModel, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, goPath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	out := make([]structModel, 0)
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}

			sm := structModel{
				StructName: ts.Name.Name,
				Attrs:      map[string]bool{},
				Children:   map[string]bool{},
			}
			for _, field := range st.Fields.List {
				if field.Tag == nil {
					continue
				}
				rawTag, err := strconv.Unquote(field.Tag.Value)
				if err != nil {
					continue
				}
				xmlTag := reflectTag(rawTag, "xml")
				if xmlTag == "" {
					continue
				}
				name, opts := splitXMLTag(xmlTag)
				if name == "-" {
					continue
				}

				fieldName := ""
				if len(field.Names) > 0 {
					fieldName = field.Names[0].Name
				}
				if fieldName == "XMLName" && name != "" {
					sm.Element = lastPathSegment(name)
					continue
				}

				if name == "" || strings.HasPrefix(name, ",") {
					continue
				}
				name = lastPathSegment(name)
				if opts["attr"] {
					sm.Attrs[name] = true
					continue
				}
				if opts["chardata"] || opts["innerxml"] || opts["comment"] {
					continue
				}
				sm.Children[name] = true
			}
			if sm.Element == "" {
				sm.Element = sm.StructName
			}
			out = append(out, sm)
		}
	}
	return out, nil
}

func reflectTag(raw, key string) string {
	parts := strings.Split(raw, " ")
	prefix := key + ":\""
	for _, p := range parts {
		if strings.HasPrefix(p, prefix) && strings.HasSuffix(p, "\"") {
			return strings.TrimSuffix(strings.TrimPrefix(p, prefix), "\"")
		}
	}
	return ""
}

func splitXMLTag(tag string) (string, map[string]bool) {
	parts := strings.Split(tag, ",")
	name := parts[0]
	opts := map[string]bool{}
	for _, p := range parts[1:] {
		if p != "" {
			opts[p] = true
		}
	}
	return name, opts
}

func lastPathSegment(name string) string {
	if strings.Contains(name, ">") {
		parts := strings.Split(name, ">")
		return parts[len(parts)-1]
	}
	return name
}

func diffSet(a, b map[string]bool) []string {
	out := make([]string, 0)
	for k := range a {
		if !b[k] {
			out = append(out, k)
		}
	}
	sort.Strings(out)
	return out
}

// listElements prints a summary of all elements in a DTD file to stdout.
func listElements(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", path, err)
		return
	}
	src := reComment.ReplaceAllString(string(data), "")
	ents := parseEntities(src)
	matches := reElement.FindAllStringSubmatch(src, -1)
	fmt.Fprintf(os.Stderr, "// %s: %d elements\n", filepath.Base(path), len(matches))
	for i, m := range matches {
		raw := strings.TrimSpace(expand(m[2], ents))
		kind := classifyRaw(raw)
		fmt.Printf("%4d %-45s %s\n", i+1, m[1], kind)
	}
}

// generateStructs reads a DTD file and writes Go struct definitions to stdout.
func generateStructs(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	src := reComment.ReplaceAllString(string(data), "")
	ents := parseEntities(src)

	// Parse ELEMENT declarations
	elems := map[string]*ElementInfo{}
	var order []string
	for _, m := range reElement.FindAllStringSubmatch(src, -1) {
		name := m[1]
		raw := strings.TrimSpace(expand(m[2], ents))
		info := &ElementInfo{Name: name, RawModel: raw}
		classify(info)
		elems[name] = info
		order = append(order, name)
	}

	// Parse ATTLIST declarations and merge into elements
	for _, m := range reAttlist.FindAllStringSubmatch(src, -1) {
		name := m[1]
		body := expand(strings.TrimSpace(m[2]), ents)
		if info, ok := elems[name]; ok {
			info.Attrs = append(info.Attrs, parseAttrs(body)...)
		}
	}

	// Count stats
	nAttrs := 0
	for _, info := range elems {
		nAttrs += len(info.Attrs)
	}
	fmt.Fprintf(os.Stderr, "// %s: %d elements, %d attributes\n", filepath.Base(path), len(order), nAttrs)

	// Output Go source
	fmt.Printf("// Code generated from %s by dtd-to-go. DO NOT EDIT DIRECTLY.\n", filepath.Base(path))
	fmt.Println("// Review and refine for idiomatic Go before committing.")
	fmt.Println()
	fmt.Println("package model")
	fmt.Println()
	fmt.Println("import \"encoding/xml\"")
	fmt.Println()
	for _, n := range order {
		printStruct(elems[n])
	}

	fmt.Fprintf(os.Stderr, "// Done. %d structs written.\n", len(order))
	return nil
}

// --- Parsing helpers ---

// parseEntities extracts parameter entity definitions: <!ENTITY % name "value">
func parseEntities(src string) map[string]string {
	m := map[string]string{}
	for _, match := range reEntity.FindAllStringSubmatch(src, -1) {
		m[match[1]] = match[2]
	}
	return m
}

// expand replaces %entity; references with their values, with multi-pass for nested refs.
func expand(s string, ents map[string]string) string {
	for i := 0; i < 10; i++ {
		prev := s
		s = reEntRef.ReplaceAllStringFunc(s, func(ref string) string {
			if v, ok := ents[ref[1:len(ref)-1]]; ok {
				return v
			}
			return ref
		})
		if s == prev {
			return s
		}
	}
	return s
}

// classifyRaw returns the kind string for a raw (expanded) content model.
func classifyRaw(raw string) string {
	switch {
	case raw == "EMPTY":
		return "EMPTY"
	case raw == "ANY":
		return "ANY"
	case raw == "(#PCDATA)" || raw == "#PCDATA":
		return "PCDATA"
	case strings.Contains(raw, "#PCDATA"):
		return "MIXED"
	default:
		return "CHILDREN"
	}
}

// classify sets Kind and Children on an ElementInfo based on its RawModel.
func classify(info *ElementInfo) {
	info.Kind = classifyRaw(info.RawModel)
	switch info.Kind {
	case "MIXED":
		info.Children = extractChildren(info.RawModel, true)
		// Mixed content (DTD spec requires outer *): all children are zero-or-more
		for i := range info.Children {
			info.Children[i].Cardinality = "*"
			info.Children[i].InChoice = false
		}
	case "CHILDREN":
		info.Children = extractChildren(info.RawModel, false)
	}
}

// extractChildren finds all child element references in a content model string.
func extractChildren(model string, mixed bool) []ChildRef {
	var refs []ChildRef
	seen := map[string]bool{}
	topChoice := isTopLevelChoice(model)

	for _, m := range reChild.FindAllStringSubmatch(model, -1) {
		name, card := m[1], m[2]
		// Filter out DTD keywords that happen to match [A-Z]...
		if name == "PCDATA" || seen[name] {
			continue
		}
		seen[name] = true
		refs = append(refs, ChildRef{
			Name:        name,
			Cardinality: card,
			InChoice:    topChoice,
		})
	}
	return refs
}

// isTopLevelChoice returns true if the top-level content model is a pure choice (| only, no ,).
func isTopLevelChoice(model string) bool {
	s := strings.TrimSpace(model)
	// Strip outer cardinality marker if present: (A | B)? → (A | B)
	s = strings.TrimRight(s, "?*+")
	if len(s) < 3 || s[0] != '(' || s[len(s)-1] != ')' {
		return false
	}
	inner := s[1 : len(s)-1]
	depth, commas, pipes := 0, 0, 0
	for _, ch := range inner {
		switch ch {
		case '(':
			depth++
		case ')':
			depth--
		case ',':
			if depth == 0 {
				commas++
			}
		case '|':
			if depth == 0 {
				pipes++
			}
		}
	}
	return pipes > 0 && commas == 0
}

// parseAttrs parses the body of an ATTLIST declaration into attribute declarations.
// Expected format after entity expansion: "name1 TYPE1 DEFAULT1 name2 TYPE2 DEFAULT2 ..."
func parseAttrs(body string) []AttrDecl {
	var attrs []AttrDecl
	tokens := strings.Fields(body)
	i := 0
	for i < len(tokens) {
		// 1. Attribute name
		name := tokens[i]
		i++
		if i >= len(tokens) {
			break
		}

		// 2. Attribute type
		var typ string
		if strings.HasPrefix(tokens[i], "(") {
			// Enumeration type — collect tokens until we find one containing ")"
			start := i
			for i < len(tokens) && !strings.Contains(tokens[i], ")") {
				i++
			}
			if i >= len(tokens) {
				break
			}
			typ = strings.Join(tokens[start:i+1], "")
			i++
		} else {
			typ = tokens[i]
			i++
		}
		if i >= len(tokens) {
			break
		}

		// 3. Default value
		def := tokens[i]
		i++
		if def == "#FIXED" && i < len(tokens) {
			def += " " + tokens[i]
			i++
		} else if strings.HasPrefix(def, "\"") && !strings.HasSuffix(def, "\"") {
			// Multi-word quoted default value
			for i < len(tokens) {
				def += " " + tokens[i]
				i++
				if strings.HasSuffix(tokens[i-1], "\"") {
					break
				}
			}
		}

		attrs = append(attrs, AttrDecl{Name: name, Type: typ, Default: def})
	}
	return attrs
}

// --- Output ---

// printStruct writes a Go struct definition to stdout for the given element.
func printStruct(info *ElementInfo) {
	// Header comment
	raw := info.RawModel
	if len(raw) > 120 {
		raw = raw[:117] + "..."
	}
	switch info.Kind {
	case "EMPTY", "ANY", "PCDATA":
		fmt.Printf("// %s — %s\n", info.Name, info.Kind)
	default:
		fmt.Printf("// %s — %s\n// Content: %s\n", info.Name, info.Kind, raw)
	}

	fmt.Printf("type %s struct {\n", goName(info.Name))
	fmt.Printf("\tXMLName xml.Name `xml:\"%s\"`\n", info.Name)

	// Attributes
	for _, a := range info.Attrs {
		printAttrField(a)
	}

	// Content
	switch info.Kind {
	case "PCDATA":
		fmt.Printf("\tValue string `xml:\",chardata\"`\n")
	case "ANY":
		fmt.Printf("\tContent string `xml:\",innerxml\"`\n")
	case "MIXED":
		fmt.Printf("\tValue string `xml:\",chardata\"`\n")
		for _, c := range info.Children {
			printChildField(c)
		}
	case "CHILDREN":
		for _, c := range info.Children {
			printChildField(c)
		}
	}

	fmt.Println("}")
	fmt.Println()
}

// printAttrField writes a single struct field line for an attribute.
func printAttrField(a AttrDecl) {
	field := goName(a.Name)
	tag := a.Name + ",attr"
	if a.Default != "#REQUIRED" {
		tag += ",omitempty"
	}

	// Build comment with enum values and REQUIRED marker
	var parts []string
	if strings.HasPrefix(a.Type, "(") {
		parts = append(parts, a.Type)
	}
	if a.Default == "#REQUIRED" {
		parts = append(parts, "REQUIRED")
	}
	comment := ""
	if len(parts) > 0 {
		comment = " // " + strings.Join(parts, " ")
	}

	fmt.Printf("\t%s string `xml:\"%s\"`%s\n", field, tag, comment)
}

// printChildField writes a single struct field line for a child element reference.
func printChildField(c ChildRef) {
	field := goName(c.Name)
	typ := goName(c.Name)

	switch {
	case c.Cardinality == "*" || c.Cardinality == "+":
		fmt.Printf("\t%s []*%s `xml:\"%s\"`\n", field, typ, c.Name)
	case c.InChoice || c.Cardinality == "?":
		fmt.Printf("\t%s *%s `xml:\"%s,omitempty\"`\n", field, typ, c.Name)
	default:
		fmt.Printf("\t%s *%s `xml:\"%s\"`\n", field, typ, c.Name)
	}
}

// goName converts a DTD name to an exported Go identifier.
func goName(s string) string {
	// Handle xml: namespace prefix
	if strings.HasPrefix(s, "xml:") {
		rest := s[4:]
		if len(rest) > 0 {
			return strings.ToUpper(rest[:1]) + rest[1:]
		}
		return s
	}
	// Capitalize first letter for export
	if len(s) > 0 && s[0] >= 'a' && s[0] <= 'z' {
		return strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}
