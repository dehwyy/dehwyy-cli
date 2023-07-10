package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/dehwyy/dehwyy-cli/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	cmdJp = &cobra.Command{
		Use: "jp [word]",
		Short: "Translate from Japanase to English or from English to Japanase",
		Long: "Translate from Japanase to English or from English to Japanase using Jisho.org",
		Args: cobra.ExactArgs(1),
		Run: runCmdJp,
	}

	FlagMaxLenJp int
)

func init() {
	cmdJp.Flags().IntVarP(&FlagMaxLenJp, "len", "l", -1, "Define max length of matching words. For limitless words use -1 (specified by default)")
	rootCmd.AddCommand(cmdJp)
}

// __________ // __________ //

type JishoResponse struct {
	Data []struct {
		Slug     string
		Jlpt     []string
		Japanese []struct {
			Word    string
			Reading string
		}
		Senses []struct {
			EnglishDefinitions []string `json:"english_definitions"`
		}
	}
}

func runCmdJp(cmd *cobra.Command, args []string) {
	baseUrl := "http://beta.jisho.org/api/v1/search/"
	word := url.PathEscape(args[0])

	// summary url
	url := baseUrl + fmt.Sprintf("words?keyword=%s", string(word))

	var body JishoResponse
	utils.FetchUrl(url, &body)

	// rendering table
	makeTableJp(body, FlagMaxLenJp)
}



func makeTableJp(tableData JishoResponse, maxLen int) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"English", "Kanji", "Reading"})

	for i, w := range tableData.Data {
		if i == maxLen {
			break
		}

		englishWord := strings.Join(w.Senses[0].EnglishDefinitions, ",")
		japKanji := w.Japanese[0].Word
		japReading := w.Japanese[0].Reading

		t.AppendRow(table.Row{englishWord, japKanji, japReading})
		t.AppendSeparator()
	}

	t.SetStyle(table.StyleLight)
	t.Render()
}
