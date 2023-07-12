package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/dehwyy/dehwyy-cli/ternary"
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
		Run: runCmdEn,
	}

	flagMaxLenEn int
)

func init() {
	cmdEng.Flags().IntVarP(&flagMaxLenEn, "len", "l", -1, "Define max length of matching words. For limitless words use -1 (specified by default)")
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

func runCmdEn(cmd *cobra.Command, args []string) {
	// getting API_KEY from .env
	utils.LoadEnv()
	key, _ := os.LookupEnv("YANDEX_TRANSLATE_API_KEY")

	word := args[0]
	// Parsing word to prevent possible collapses that could happen due to non-english request
	wordInUrl := url.PathEscape(word)

	// clarifying whether word is english or russian, err appears when any symbol is other then previous
	isEng, err := utils.IsEnglishWord(word)
	if err != nil {
		log.Fatalln(err)
	}

	// if isEng => en-ru else ru-en
	var translate = ternary.Use(isEng, "en-ru", "ru-en")

	// url to fetch
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

	for i, w := range tableData.Def {
		// if limit reached => quit
		if i == flagMaxLenEn {
			break
		}

		// word that was queried by user
		queried_word := w.Text

		// Iterating through all translations
		for _, tr := range w.Tr {

			// meaning
			var meaning string
			for _, m := range tr.Mean {
				meaning += fmt.Sprintf("%s, ", m.Text)
			}
			// translation
			translation := tr.Text

			// synonyms
			var synonyms string
			for _, syn := range tr.Syn {
				synonyms += fmt.Sprintf("%s, ", syn.Text)
			}

			// adding to the table
			t.AppendRow(table.Row{queried_word, translation, meaning, synonyms})
			t.AppendSeparator()
		}

	}

	t.SetStyle(table.StyleLight)
	t.Render()
}
