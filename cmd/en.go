package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/dehwyy/dehwyy-cli/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	cmdEng = &cobra.Command{
		Use: "en [word]",
		Short: "Translate from Russian to English or from English to Russian",
		Long: "Translate from Russian to English or from English to Russian using YandexDictionary",
		Args: cobra.ExactArgs(1),
		Run: runCmdEng,
	}

)

func init() {
	rootCmd.AddCommand(cmdEng)
}

// __________ // __________ //

type YandexResponse struct {
	Def []struct {
		Text string
		Tr  []struct {
			Text string
			Syn  []struct {
				Text string
			}
			Mean []struct {
				Text string
			}
		}
	}
}

func runCmdEng(cmd *cobra.Command, args []string) {
	// loading env and getting API_KEY
	utils.LoadEnv()
	key, _ := os.LookupEnv("YANDEX_TRANSLATE_API_KEY")

	// which word to look in dict
	word := args[0]

	// Parsing word to prevent possible collapses that could happen due to non-english request
	wordInUrl := url.PathEscape(word)

	// clarifying whether word is english or russian, err appears when any symbol is other then previous
	isEng, err := utils.IsEnglishWord(word)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var translate string

	if isEng {
		translate = "en-ru"
	} else {
		translate = "ru-en"
	}

	url := fmt.Sprintf("https://dictionary.yandex.net/api/v1/dicservice.json/lookup?key=%s&lang=%s&text=%s", key, translate, wordInUrl)

	var body YandexResponse
	utils.FetchUrl(url, &body)

	// rendering table
	makeTableEn(body)
}



func makeTableEn(tableData YandexResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Rows' names
	t.AppendHeader(table.Row{"Word", "Translation", "Meaning", "Synonyms"})

	for _, w := range tableData.Def {
		word := w.Text

		for _, tr := range w.Tr {
			// meaning
			var mean []string
			for _, m := range tr.Mean {
				mean = append(mean, m.Text)
			}
			meanString := strings.Join(mean, ", ")

			// translation
			translation := tr.Text

			// synonyms
			var synonyms []string
			for _, syn := range tr.Syn {
				synonyms = append(synonyms, syn.Text)
			}
			synonymsString := strings.Join(synonyms, ", ")

			// adding to the table
			t.AppendRow(table.Row{word, translation, meanString, synonymsString})
			t.AppendSeparator()
		}

	}

	t.SetStyle(table.StyleLight)
	t.Render()
}
