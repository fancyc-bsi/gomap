package reporting

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Comprehensive Report</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootswatch/4.5.2/darkly/bootstrap.min.css">
    <style>
        body { padding: 20px; }
        pre { background: #2d2d2d; padding: 10px; border-radius: 5px; color: #f8f9fa; }
        code { background: #2d2d2d; padding: 2px 4px; border-radius: 3px; color: #f8f9fa; }
        .card { margin-bottom: 20px; background: #343a40; color: #f8f9fa; }
        .navbar { margin-bottom: 20px; background: #343a40; color: #f8f9fa; }
        .nav-link { color: #f8f9fa; }
        .nav-link:hover { color: #d1d1d1; }
    </style>
</head>
<body>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <a class="navbar-brand" href="#">Report</a>
        <div class="collapse navbar-collapse">
            <ul class="navbar-nav mr-auto">
                %s
            </ul>
        </div>
    </nav>
    <div class="container">
        %s
    </div>
</body>
</html>`

func RenderReport(outputDir, reportFile string) error {
	mdFiles, err := filepath.Glob(filepath.Join(outputDir, "*.md"))
	if err != nil {
		return fmt.Errorf("failed to list markdown files: %v", err)
	}

	var navItems, reportContent string
	for _, mdFile := range mdFiles {
		filename := filepath.Base(mdFile)
		title := strings.TrimSuffix(filename, filepath.Ext(filename))
		navItems += fmt.Sprintf(`<li class="nav-item"><a class="nav-link" href="#%s">%s</a></li>`, title, title)

		content, err := os.ReadFile(mdFile)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %v", mdFile, err)
		}

		wrappedContent := "```\n" + string(content) + "\n```"
		mdHTML := blackfriday.Run([]byte(wrappedContent), blackfriday.WithExtensions(blackfriday.FencedCode|blackfriday.Autolink|blackfriday.Strikethrough|blackfriday.SpaceHeadings|blackfriday.Tables))
		reportContent += fmt.Sprintf(`<div id="%s" class="card"><div class="card-body"><h5 class="card-title">%s</h5>%s</div></div>`, title, title, mdHTML)
	}

	htmlContent := fmt.Sprintf(htmlTemplate, navItems, reportContent)

	err = os.WriteFile(reportFile, []byte(htmlContent), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write report file: %v", err)
	}

	return nil
}
