package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"text/template"
)

// FieldInfo содержит информацию о поле структуры
type FieldInfo struct {
	Name string // Имя поля
	Type string // Тип поля в виде строки
	Tag  string // Теги поля (например, `yaml:"name"`)
}

// GenerateMarshaller генерирует метод StructToMap для структуры Config
func GenerateMarshaller() error {
	// Путь к файлу с конфигом
	configPath := filepath.Join("internal", "config", "config.go")
	
	// Анализируем исходный файл
	fields, err := analyzeConfig(configPath)
	if err != nil {
		return fmt.Errorf("analysis failed: %w", err)
	}

	// Генерируем код
	if err := generateCode(fields); err != nil {
		return fmt.Errorf("generation failed: %w", err)
	}

	return nil
}

// analyzeConfig анализирует файл и извлекает информацию о полях
func analyzeConfig(filePath string) ([]FieldInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var fields []FieldInfo

	// Ищем структуру Config в AST
	ast.Inspect(node, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok || ts.Name.Name != "Config" {
			return true
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		// Обрабатываем каждое поле структуры
		for _, f := range st.Fields.List {
			tag := ""
			if f.Tag != nil {
				tag = f.Tag.Value
			}

			for _, name := range f.Names {
				fields = append(fields, FieldInfo{
					Name: name.Name,
					Type: exprToString(f.Type),
					Tag:  tag,
				})
			}
		}
		return false
	})

	return fields, nil
}

// generateCode создает файл marshaller_gen.go
func generateCode(fields []FieldInfo) error {
	// Загружаем шаблон
	tmpl, err := template.ParseFiles(filepath.Join("assets", "template", "marshaller.gotmpl"))
	if err != nil {
		return err
	}

	// Создаем выходной файл
	outPath := filepath.Join("internal", "config", "marshaller_gen.go")
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Генерируем код
	data := struct {
		Package string
		Fields  []FieldInfo
	}{
		Package: "config",
		Fields:  fields,
	}

	// Записываем сгенерированный код
	if err := tmpl.Execute(f, data); err != nil {
		return err
	}

	return nil
}

// exprToString преобразует AST-выражение в строку типа
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	case *ast.MapType:
		return "map[" + exprToString(t.Key) + "]" + exprToString(t.Value)
	default:
		return "interface{}"
	}
}
